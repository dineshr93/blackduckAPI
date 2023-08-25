package service

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dineshr93/blackduckAPI/helpers"
	"github.com/dineshr93/blackduckAPI/models"
	"github.com/dineshr93/blackduckAPI/outputmodel"
	"github.com/dineshr93/blackduckAPI/scansummariesmodel"
	"github.com/dineshr93/blackduckAPI/sourcebomentriesmodel"
)

// Create a struct to hold the parsed data
type HUB struct {
	config        *models.ConfigFile
	links         *models.APILinks
	accepts       *models.Accepts
	reqEssentials *models.RequestPrepEssentials
	client        *http.Client
}

func (h *HUB) Loadinit() {
	h.config = helpers.ParseYML()
	h.links = helpers.LoadLinks(h.config.Apihost)
	h.accepts = helpers.LoadAccepts()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	tr := &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	h.client = &http.Client{Transport: tr}
	h.LoadBearerToken()
}

func (h *HUB) LoadPVlinks(pvlink string) {
	h.links.SetPVBulkSnippetBomEntries(pvlink)
}
func (h *HUB) LoadSpecificCodeLocation(clLink string) {
	h.links.SetCodelocationScanSummaries(clLink)
}
func (h *HUB) LoadMatchToComponentVersion(matchTOComponent string, matchToComponentVersion string) {
	h.links.MatchTocomponent = matchTOComponent
	h.links.MatchTocomponentversion = matchToComponentVersion
}
func (h *HUB) LoadBearerToken() {
	succRes := &models.AuthResponse{}
	errRes := &models.ErroredResponse{}
	// 1. Req Prep
	h.reqEssentials = &models.RequestPrepEssentials{
		Config:      *h.config,
		BearerToken: "",
		Method:      http.MethodPost,
		APILinks:    h.links.Authlink,
		QueryString: "",
		Accepts:     h.accepts.AppUser4json,
		ContentType: "",
		Body:        nil,
	}
	req := helpers.PrepareRequest(h.reqEssentials)
	//2. Get response
	dataBytes, statusCode, err := helpers.GetResponse(h.client, req)
	if err != nil {
		panic(err)
	}
	if statusCode != http.StatusOK {
		log.Fatalln("Couln't read response body")

		if err != nil {
			log.Fatalln("Error while unmarshalling json", err)
			panic(errRes.ErrorMessage)
		}
		panic(errRes.ErrorMessage)
	}
	err = json.Unmarshal(dataBytes, succRes)
	if err != nil {
		log.Fatalln("Error while unmarshalling json", err)
		panic(err)
	}
	h.reqEssentials.BearerToken = "bearer " + succRes.BearerToken
	// println("stored btoken: ", h.reqEssentials.BearerToken)
}
func (h *HUB) SetScansummariesURL() {
	succRes := &scansummariesmodel.ScanSummaries{}
	errRes := &models.ErroredResponse{}
	// 1. Req Prep
	h.reqEssentials = &models.RequestPrepEssentials{
		Config:      *h.config,
		BearerToken: h.reqEssentials.BearerToken,
		Method:      http.MethodGet,
		APILinks:    h.links.SpecificCodelocationAllScansummarieslink,
		QueryString: "",
		Accepts:     h.accepts.AppScan6json,
		ContentType: "",
		Body:        nil,
	}
	req := helpers.PrepareRequest(h.reqEssentials)
	//2. Get response
	dataBytes, statusCode, err := helpers.GetResponse(h.client, req)
	if err != nil {
		panic(err)
	}
	if statusCode != http.StatusOK {
		log.Fatalln("Couln't read response body")

		if err != nil {
			log.Fatalln("Error while unmarshalling json", err)
			panic(errRes.ErrorMessage)
		}
		panic(errRes.ErrorMessage)
	}
	err = json.Unmarshal(dataBytes, succRes)
	if err != nil {
		log.Fatalln("Error while unmarshalling json", err)
		panic(err)
	}
	for _, i := range succRes.Items {
		if i.ScanType == "SNIPPET" {
			h.links.ScansummariesURI = i.Meta.Href
		}

	}

	println("ScansummariesURI: ", h.links.ScansummariesURI)
}

func (h *HUB) SetSourceBomEntries(entries string) {
	succRes := &sourcebomentriesmodel.SourceBOMEntries{}
	errRes := &models.ErroredResponse{}
	// 1. Req Prep
	h.reqEssentials = &models.RequestPrepEssentials{
		Config:      *h.config,
		BearerToken: h.reqEssentials.BearerToken,
		Method:      http.MethodGet,
		APILinks:    h.links.SourceBOMEntriesLink,
		QueryString: "?limit=" + entries + "&offset=0&filter=bomMatchType%3Asnippet&filter=bomMatchReviewStatus%3Anot_reviewed",
		Accepts:     h.accepts.AppJson,
		ContentType: "",
		Body:        nil,
	}
	req := helpers.PrepareRequest(h.reqEssentials)
	// req.Header.Add("limit", entries)
	// req.Header.Add("offset", "0")
	// req.Header.Add("filter", "bomMatchType:snippet")
	// req.Header.Add("filter", "bomMatchReviewStatus:not_reviewed")
	//2. Get response
	dataBytes, statusCode, err := helpers.GetResponse(h.client, req)
	if err != nil {
		panic(err)
	}
	if statusCode != http.StatusOK {
		log.Fatalln("Couln't read response body")

		if err != nil {
			log.Fatalln("Error while unmarshalling json", err)
			panic(errRes.ErrorMessage)
		}
		panic(errRes.ErrorMessage)
	}
	err = json.Unmarshal(dataBytes, succRes)
	if err != nil {
		log.Fatalln("Error while unmarshalling json", err)
		panic(err)
	}

	h.reqEssentials.SourcebomEntries = succRes

	println("SourceBOMEntries items: ", len(h.reqEssentials.SourcebomEntries.Items))
}

func (h *HUB) PrepareBSBEBody(matchString string) {
	sourcebomEntries := h.reqEssentials.SourcebomEntries
	println(sourcebomEntries.TotalCount)
	h.reqEssentials.BSBEBody = ""

	// var items outputmodel.OutputModel
	var items []outputmodel.Item
	// items.Items

	for _, s_item := range sourcebomEntries.Items {
		var item outputmodel.Item
		if strings.Contains(s_item.CompositePath.Path, matchString) {
			item.ScanSummary = h.links.ScansummariesURI
			item.URI = "file:///" + s_item.CompositePath.Path
			item.CompositePath.Path = s_item.CompositePath.Path
			item.CompositePath.ArchiveContext = s_item.CompositePath.ArchiveContext
			item.CompositePath.CompositePathContext = s_item.CompositePath.CompositePathContext
			item.CompositePath.FileName = s_item.CompositePath.FileName
			var sbcrArray []outputmodel.SnippetBomComponentRequest
			for _, fsbc := range s_item.FileSnippetBomComponents {
				var sbcr outputmodel.SnippetBomComponentRequest
				sbcr.SourceStartLines = fsbc.SourceStartLines
				sbcr.SourceEndLines = fsbc.SourceEndLines
				sbcr.Ignored = fsbc.Ignored
				sbcr.HashID = fsbc.HashID
				sbcr.ReviewStatus = "REVIEWED"
				sbcr.Component = h.links.MatchTocomponent
				sbcr.ComponentVersion = h.links.MatchTocomponentversion
				sbcrArray = append(sbcrArray, sbcr)
			}
			item.SnippetBomComponentRequests = sbcrArray
			items = append(items, item)
		}
	}
	// fmt.Println(items)
	b, err := json.Marshal(&items)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := string(b)
	fmt.Println(data)

	h.reqEssentials.BSBEBody = data
}

func (h *HUB) MatchSnippetToGivenComponent() {

	errRes := &models.ErroredResponse{}
	// 1. Req Prep
	h.reqEssentials = &models.RequestPrepEssentials{
		Config:      *h.config,
		BearerToken: h.reqEssentials.BearerToken,
		Method:      http.MethodPut,
		APILinks:    h.links.SpecificPVBulkSnippetBomEntries,
		QueryString: "",
		Accepts:     h.accepts.AppBom6json,
		ContentType: h.accepts.AppBom6json,
		Body:        strings.NewReader(h.reqEssentials.BSBEBody),
	}
	req := helpers.PrepareRequest(h.reqEssentials)
	// req.Header.Add("limit", entries)
	// req.Header.Add("offset", "0")
	// req.Header.Add("filter", "bomMatchType:snippet")
	// req.Header.Add("filter", "bomMatchReviewStatus:not_reviewed")
	//2. Get response
	_, statusCode, err := helpers.GetResponse(h.client, req)
	if err != nil {
		panic(err)
	}
	if statusCode != http.StatusNoContent {
		log.Fatalln("Couln't read response body")

		if err != nil {
			log.Fatalln("Error while unmarshalling json", err)
			panic(errRes.ErrorMessage)
		}
		panic(errRes.ErrorMessage)
	}

	println("Snippet match status code: ", statusCode)
}

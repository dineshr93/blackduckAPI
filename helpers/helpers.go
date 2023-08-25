package helpers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dineshr93/blackduckAPI/models"
	"gopkg.in/yaml.v3"
)

const (
	Accept        = "Accept"
	Authorization = "Authorization"
	ContentType   = "Content-Type"
)

func ParseYML() *models.ConfigFile {

	// Read the YAML file
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// Create a struct to hold the parsed data
	var config models.ConfigFile

	// Parse the YAML data into the struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}
	// println(config.Apihost)
	// println(config.InitToken)
	return &config
}

func LoadLinks(baselink string) *models.APILinks {
	return &models.APILinks{
		Authlink:          baselink + "tokens/authenticate",
		Codelocationslink: baselink + "codelocations",
		Projectlink:       baselink + "projects",
	}
}

func LoadAccepts() *models.Accepts {
	return &models.Accepts{
		AppOctetStream:               "application/octet-stream",
		AppScanstatus1json:           "application/vnd.blackducksoftare.scanstatus-1+json",
		AppAdmin4json:                "application/vnd.blackducksoftware.admin-4+json",
		AppAdmin5json:                "application/vnd.blackducksoftware.admin-5+json",
		AppBdiozip:                   "application/vnd.blackducksoftware.bdio+zip",
		AppBom4json:                  "application/vnd.blackducksoftware.bill-of-materials-4+json",
		AppBom5json:                  "application/vnd.blackducksoftware.bill-of-materials-5+json",
		AppBom6json:                  "application/vnd.blackducksoftware.bill-of-materials-6+json",
		AppComponentDetail4json:      "application/vnd.blackducksoftware.component-detail-4+json",
		AppComponentDetail5json:      "application/vnd.blackducksoftware.component-detail-5+json",
		Appcopyright4json:            "application/vnd.blackducksoftware.copyright-4+json",
		AppDeveloperscan1ld2json:     "application/vnd.blackducksoftware.developer-scan-1-ld-2+json",
		AppDeveloperscan1ld2yaml1zip: "application/vnd.blackducksoftware.developer-scan-1-ld-2-yaml-1+zip",
		AppJournal4json:              "application/vnd.blackducksoftware.journal-4+json",
		AppNotification4json:         "application/vnd.blackducksoftware.notification-4+json",
		AppPolicy4json:               "application/vnd.blackducksoftware.policy-4+json",
		AppPolicy5json:               "application/vnd.blackducksoftware.policy-5+json",
		AppProjectDetail4json:        "application/vnd.blackducksoftware.project-detail-4+json",
		AppProjectDetail5json:        "application/vnd.blackducksoftware.project-detail-5+json",
		AppProjectDetail6json:        "application/vnd.blackducksoftware.project-detail-6+json",
		AppReport1csv:                "application/vnd.blackducksoftware.report-1+csv",
		AppReport4json:               "application/vnd.blackducksoftware.report-4+json",
		AppScan4json:                 "application/vnd.blackducksoftware.scan-4+json",
		AppScan5json:                 "application/vnd.blackducksoftware.scan-5+json",
		AppScan6json:                 "application/vnd.blackducksoftware.scan-6+json",
		AppScanmonitor1json:          "application/vnd.blackducksoftware.scan-monitor-1+json",
		AppScanreadiness1json:        "application/vnd.blackducksoftware.scan-readiness-1+json",
		AppSourceview1json:           "application/vnd.blackducksoftware.source-view-1+json",
		AppStatus4json:               "application/vnd.blackducksoftware.status-4+json",
		AppSystemAnnouncement1json:   "application/vnd.blackducksoftware.system-announcement-1+json",
		AppToolview1json:             "application/vnd.blackducksoftware.tool-view-1+json",
		AppUser4json:                 "application/vnd.blackducksoftware.user-4+json",
		AppVulnerability4json:        "application/vnd.blackducksoftware.vulnerability-4+json",
		AppZip:                       "application/zip",
		MultipartFormData:            "multipart/form-data",
		TextPlain:                    "text/plain",
		AppJson:                      "application/json",
	}
}

func PrepareRequest(reqPrep *models.RequestPrepEssentials) *http.Request {
	println(reqPrep.APILinks + reqPrep.QueryString)
	req, err := http.NewRequest(reqPrep.Method, reqPrep.APILinks+reqPrep.QueryString, reqPrep.Body)

	if err != nil {
		log.Fatalln(err)
	}
	token := reqPrep.BearerToken

	// initial token to get bearer token
	if strings.Contains(reqPrep.APILinks, "authenticate") && token == "" {
		token = reqPrep.Config.InitToken
	}
	// println(token)
	req.Header.Add(Authorization, token)
	req.Header.Add(Accept, reqPrep.Accepts)
	if reqPrep.ContentType != "" {
		req.Header.Add(ContentType, reqPrep.ContentType)
	}

	return req
}
func GetResponse(client *http.Client, req *http.Request) ([]byte, int, error) {
	// println(req.Header.Get(ContentType))
	// println(req.Header.Get(Accept))
	// println(req.RequestURI)
	// Loop over header names
	fmt.Println("=============================")
	for name, values := range req.Header {
		// Loop over all values for the name.
		for _, value := range values {
			if name == "Authorization" {
				continue
			}
			fmt.Println(name, value)
		}
	}
	fmt.Println("=============================")
	// println()

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("error while client.Do(req)", err)
		return nil, res.StatusCode, err
	}

	// Handling for PUT request, 204 is success with no BODY
	if req.Method == http.MethodPut {
		var b []byte
		if res.StatusCode != http.StatusNoContent {
			log.Fatalln("Status code", res.StatusCode, err)

			return b, res.StatusCode, err
		}
		return b, res.StatusCode, err
	}
	databytes, err := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Fatalln("Status code", res.StatusCode, string(databytes))
		return databytes, res.StatusCode, err
	}
	if err != nil {
		log.Fatalln("Couln't read response body")
		return databytes, res.StatusCode, err
	}

	return databytes, res.StatusCode, nil
}

// SPDX-FileCopyrightText: 2023 Dinesh Ravi
//
// SPDX-License-Identifier: GPL-3.0-only

package models

import (
	"io"
	"strings"

	"github.com/dineshr93/blackduckAPI/outputmodel"
	"github.com/dineshr93/blackduckAPI/sourcebomentriesmodel"
)

type ConfigFile struct {
	Apihost   string `yaml:"apihost"`
	InitToken string `yaml:"token"`
}
type AuthResponse struct {
	BearerToken           string `json:"bearerToken"`
	ExpiresInMilliseconds int    `json:"expiresInMilliseconds"`
}
type ErroredResponse struct {
	LogRef       string `json:"logRef"`
	ErrorMessage string `json:"errorMessage"`
	Arguments    struct {
	} `json:"arguments"`
	ErrorCode string `json:"errorCode"`
	Errors    []any  `json:"errors"`
	Links     []any  `json:"links"`
}

type Accepts struct {
	AppOctetStream               string
	AppScanstatus1json           string
	AppAdmin4json                string
	AppAdmin5json                string
	AppBdiozip                   string
	AppBom4json                  string
	AppBom5json                  string
	AppBom6json                  string
	AppComponentDetail4json      string
	AppComponentDetail5json      string
	Appcopyright4json            string
	AppDeveloperscan1ld2json     string
	AppDeveloperscan1ld2yaml1zip string
	AppJournal4json              string
	AppNotification4json         string
	AppPolicy4json               string
	AppPolicy5json               string
	AppProjectDetail4json        string
	AppProjectDetail5json        string
	AppProjectDetail6json        string
	AppReport1csv                string
	AppReport4json               string
	AppScan4json                 string
	AppScan5json                 string
	AppScan6json                 string
	AppScanmonitor1json          string
	AppScanreadiness1json        string
	AppSourceview1json           string
	AppStatus4json               string
	AppSystemAnnouncement1json   string
	AppToolview1json             string
	AppUser4json                 string
	AppVulnerability4json        string
	AppZip                       string
	MultipartFormData            string
	TextPlain                    string
	AppJson                      string
}

type APILinks struct {
	Authlink string
	//
	Codelocationslink                        string
	SpecificCodelocationslink                string
	SpecificCodelocationAllScansummarieslink string
	SpecificScansummaries                    string
	//
	Projectlink                     string
	SpecificProjectLink             string
	SpecificPVlink                  string
	SpecificPVBulkSnippetBomEntries string
	SpecificPVScansummaries         string
	//
	//
	InternalSourceBOMEntries string

	MatchTocomponent        string
	MatchTocomponentversion string

	SourceBOMEntriesLink string
	ScansummariesURI     string
}

func (l *APILinks) SetPVBulkSnippetBomEntries(pvlink string) {
	l.SpecificPVlink = pvlink
	l.SpecificPVBulkSnippetBomEntries = l.SpecificPVlink + "/bulk-snippet-bom-entries"
	withInternal := strings.Replace(pvlink, "/api/", "/api/internal/", -1)
	l.SourceBOMEntriesLink = withInternal + "/source-bom-entries"

}
func (l *APILinks) SetCodelocationScanSummaries(cllink string) {
	l.SpecificCodelocationslink = cllink
	l.SpecificCodelocationAllScansummarieslink = l.SpecificCodelocationslink + "/scan-summaries"
}

type RequestPrepEssentials struct {
	Config           ConfigFile
	BearerToken      string
	Method           string
	APILinks         string
	QueryString      string
	Accepts          string
	ContentType      string
	Body             io.Reader
	SourcebomEntries *sourcebomentriesmodel.SourceBOMEntries
	BSBEBody         string
	OutputModel      *outputmodel.OutputModel
}

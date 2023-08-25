package scansummariesmodel

import "time"

type ScanSummaries struct {
	TotalCount     int     `json:"totalCount,omitempty"`
	Items          []Items `json:"items,omitempty"`
	AppliedFilters []any   `json:"appliedFilters,omitempty"`
	Meta           Meta    `json:"_meta,omitempty"`
}
type Links struct {
	Rel  string `json:"rel,omitempty"`
	Href string `json:"href,omitempty"`
}
type Meta struct {
	Allow []string `json:"allow,omitempty"`
	Href  string   `json:"href,omitempty"`
	Links []Links  `json:"links,omitempty"`
}
type Items struct {
	ServerVersion        string    `json:"serverVersion,omitempty"`
	ScanState            string    `json:"scanState,omitempty"`
	TransitionReason     string    `json:"transitionReason,omitempty"`
	CreatedAt            time.Time `json:"createdAt,omitempty"`
	CreatedByUserName    string    `json:"createdByUserName,omitempty"`
	ScanSize             int       `json:"scanSize,omitempty"`
	MatchCount           int       `json:"matchCount,omitempty"`
	DirectoryCount       int       `json:"directoryCount,omitempty"`
	FileCount            int       `json:"fileCount,omitempty"`
	BaseDirectory        string    `json:"baseDirectory,omitempty"`
	HostName             string    `json:"hostName,omitempty"`
	ScanType             string    `json:"scanType,omitempty"`
	RetainUnmatchedFiles bool      `json:"retainUnmatchedFiles,omitempty"`
	Meta                 Meta      `json:"_meta,omitempty"`
	UpdatedAt            time.Time `json:"updatedAt,omitempty"`
	StatusMessage        string    `json:"statusMessage,omitempty"`
}

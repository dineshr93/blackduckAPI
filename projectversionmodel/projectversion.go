package projectversionmodel

import "time"

type ProjectVersion struct {
	VersionName          string    `json:"versionName,omitempty"`
	Phase                string    `json:"phase,omitempty"`
	Distribution         string    `json:"distribution,omitempty"`
	License              License   `json:"license,omitempty"`
	CreatedAt            time.Time `json:"createdAt,omitempty"`
	CreatedBy            string    `json:"createdBy,omitempty"`
	CreatedByUser        string    `json:"createdByUser,omitempty"`
	SettingUpdatedAt     time.Time `json:"settingUpdatedAt,omitempty"`
	SettingUpdatedBy     string    `json:"settingUpdatedBy,omitempty"`
	SettingUpdatedByUser string    `json:"settingUpdatedByUser,omitempty"`
	Source               string    `json:"source,omitempty"`
	Meta                 Meta      `json:"_meta,omitempty"`
}
type LicenseFamilySummary struct {
	Name string `json:"name,omitempty"`
	Href string `json:"href,omitempty"`
}
type Licenses struct {
	License              string               `json:"license,omitempty"`
	Licenses             []any                `json:"licenses,omitempty"`
	Name                 string               `json:"name,omitempty"`
	Ownership            string               `json:"ownership,omitempty"`
	LicenseDisplay       string               `json:"licenseDisplay,omitempty"`
	LicenseFamilySummary LicenseFamilySummary `json:"licenseFamilySummary,omitempty"`
}
type License struct {
	Type           string     `json:"type,omitempty"`
	Licenses       []Licenses `json:"licenses,omitempty"`
	LicenseDisplay string     `json:"licenseDisplay,omitempty"`
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

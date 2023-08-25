package projectsmodel

import "time"

type Projects struct {
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
	Name                     string    `json:"name,omitempty"`
	ProjectLevelAdjustments  bool      `json:"projectLevelAdjustments,omitempty"`
	CloneCategories          []string  `json:"cloneCategories,omitempty"`
	CustomSignatureEnabled   bool      `json:"customSignatureEnabled,omitempty"`
	CustomSignatureDepth     string    `json:"customSignatureDepth,omitempty"`
	DeepLicenseDataEnabled   bool      `json:"deepLicenseDataEnabled,omitempty"`
	SnippetAdjustmentApplied bool      `json:"snippetAdjustmentApplied,omitempty"`
	LicenseConflictsEnabled  bool      `json:"licenseConflictsEnabled,omitempty"`
	ProjectGroup             string    `json:"projectGroup,omitempty"`
	CreatedAt                time.Time `json:"createdAt,omitempty"`
	CreatedBy                string    `json:"createdBy,omitempty"`
	CreatedByUser            string    `json:"createdByUser,omitempty"`
	UpdatedAt                time.Time `json:"updatedAt,omitempty"`
	UpdatedBy                string    `json:"updatedBy,omitempty"`
	UpdatedByUser            string    `json:"updatedByUser,omitempty"`
	Source                   string    `json:"source,omitempty"`
	Meta                     Meta      `json:"_meta,omitempty"`
}

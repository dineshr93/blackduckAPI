package sourcebomentriesmodel

import "time"

type SourceBOMEntries struct {
	TotalCount     int              `json:"totalCount,omitempty"`
	Items          []Items          `json:"items,omitempty"`
	AppliedFilters []AppliedFilters `json:"appliedFilters,omitempty"`
	Meta           Meta             `json:"_meta,omitempty"`
}
type CompositePath struct {
	Path                 string `json:"path,omitempty"`
	ArchiveContext       string `json:"archiveContext,omitempty"`
	CompositePathContext string `json:"compositePathContext,omitempty"`
	FileName             string `json:"fileName,omitempty"`
}
type FileSignaturePair struct {
	SignatureType  string `json:"signatureType,omitempty"`
	SignatureValue string `json:"signatureValue,omitempty"`
}
type Project struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Restructured bool   `json:"restructured,omitempty"`
}
type Release struct {
	ID      string `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
}
type ChannelRelease struct {
	ID                            string `json:"id,omitempty"`
	ComponentID                   string `json:"componentId,omitempty"`
	VersionID                     string `json:"versionId,omitempty"`
	Name                          string `json:"name,omitempty"`
	ExternalNamespace             string `json:"externalNamespace,omitempty"`
	ExternalID                    string `json:"externalId,omitempty"`
	ExternalNamespaceDistribution bool   `json:"externalNamespaceDistribution,omitempty"`
}
type License struct {
	LicenseID      string   `json:"licenseId,omitempty"`
	Name           string   `json:"name,omitempty"`
	SpdxID         string   `json:"spdxId,omitempty"`
	Ownership      string   `json:"ownership,omitempty"`
	CodeSharing    int      `json:"codeSharing,omitempty"`
	Licenses       []any    `json:"licenses,omitempty"`
	LicenseDisplay string   `json:"licenseDisplay,omitempty"`
	LicenseIds     []string `json:"licenseIds,omitempty"`
}
type FileSnippetBomComponent struct {
	VersionBomComponentID            int            `json:"versionBomComponentId,omitempty"`
	VersionBomFileID                 int            `json:"versionBomFileId,omitempty"`
	VersionBomEntryID                int            `json:"versionBomEntryId,omitempty"`
	Project                          Project        `json:"project,omitempty"`
	Release                          Release        `json:"release,omitempty"`
	ChannelRelease                   ChannelRelease `json:"channelRelease,omitempty"`
	VersionBomMatchType              string         `json:"versionBomMatchType,omitempty"`
	Usage                            string         `json:"usage,omitempty"`
	License                          License        `json:"license,omitempty"`
	MatchContent                     string         `json:"matchContent,omitempty"`
	Adjusted                         bool           `json:"adjusted,omitempty"`
	MatchFileSha1                    string         `json:"matchFileSha1,omitempty"`
	MatchFileBasename                string         `json:"matchFileBasename,omitempty"`
	MatchFileExtension               string         `json:"matchFileExtension,omitempty"`
	MatchFilePath                    string         `json:"matchFilePath,omitempty"`
	MatchComponentRank               float64        `json:"matchComponentRank,omitempty"`
	MatchScore                       float64        `json:"matchScore,omitempty"`
	MatchCoverage                    int            `json:"matchCoverage,omitempty"`
	MatchComponentVersionReleaseDate time.Time      `json:"matchComponentVersionReleaseDate,omitempty"`
	MatchComponentTotalReleases      int            `json:"matchComponentTotalReleases,omitempty"`
	MatchStartLines                  []int          `json:"matchStartLines,omitempty"`
	MatchEndLines                    []int          `json:"matchEndLines,omitempty"`
	MatchLicenseCodeSharingType      int            `json:"matchLicenseCodeSharingType,omitempty"`
	AlternateSnippetMatchesURL       string         `json:"alternateSnippetMatchesUrl,omitempty"`
	MatchedFileContentURL            string         `json:"matchedFileContentUrl,omitempty"`
	SourceStartLines                 []int          `json:"sourceStartLines,omitempty"`
	SourceEndLines                   []int          `json:"sourceEndLines,omitempty"`
	Ignored                          bool           `json:"ignored,omitempty"`
	ReviewStatus                     string         `json:"reviewStatus,omitempty"`
	HashID                           string         `json:"hashId,omitempty"`
	PackageURL                       string         `json:"packageUrl,omitempty"`
	MatchFileFullPath                string         `json:"matchFileFullPath,omitempty"`
	MatchFileoccurences              int            `json:"matchFileoccurences,omitempty"`
}
type FileStringSearchMatches struct {
	MatchType string `json:"matchType,omitempty"`
	Name      string `json:"name,omitempty"`
	Vsl       string `json:"vsl,omitempty"`
}

type Items struct {
	ScanID                        string                    `json:"scanId,omitempty"`
	Name                          string                    `json:"name,omitempty"`
	CompositeID                   int                       `json:"compositeId,omitempty"`
	Size                          int                       `json:"size,omitempty"`
	Type                          string                    `json:"type,omitempty"`
	CompositePath                 CompositePath             `json:"compositePath,omitempty"`
	FileSignaturePair             FileSignaturePair         `json:"fileSignaturePair,omitempty"`
	FileDependencyBomComponents   []any                     `json:"fileDependencyBomComponents,omitempty"`
	PartialFileMatchBomComponents []any                     `json:"partialFileMatchBomComponents,omitempty"`
	BinaryFileMatchBomComponents  []any                     `json:"binaryFileMatchBomComponents,omitempty"`
	FileSnippetBomComponents      []FileSnippetBomComponent `json:"fileSnippetBomComponents,omitempty"`
	FileStringSearchMatches       []FileStringSearchMatches `json:"fileStringSearchMatches,omitempty"`
	PolicyApprovalStatus          string                    `json:"policyApprovalStatus,omitempty"`
	ActivePolicies                []any                     `json:"activePolicies,omitempty"`
	URI                           string                    `json:"uri,omitempty"`
	PolicyOverriddenBySet         []any                     `json:"policyOverriddenBySet,omitempty"`
	CommentPath                   string                    `json:"commentPath,omitempty"`
	CommentCount                  int                       `json:"commentCount,omitempty"`
	ActivePolicyCount             int                       `json:"activePolicyCount,omitempty"`
	SourceUploaded                bool                      `json:"sourceUploaded,omitempty"`
	FileBomMatchType              string                    `json:"fileBomMatchType,omitempty"`
	Meta                          Meta                      `json:"_meta,omitempty"`
	FileMatchPresent              bool                      `json:"fileMatchPresent,omitempty"`
	AnyMatchPresent               bool                      `json:"anyMatchPresent,omitempty"`
}
type Selected struct {
	Key   string `json:"key,omitempty"`
	Label string `json:"label,omitempty"`
}
type AppliedFilters struct {
	Name     string     `json:"name,omitempty"`
	Label    string     `json:"label,omitempty"`
	Selected []Selected `json:"selected,omitempty"`
}
type Links struct {
	Rel   string `json:"rel,omitempty"`
	Href  string `json:"href,omitempty"`
	Name  string `json:"name,omitempty"`
	Label string `json:"label,omitempty"`
}
type Meta struct {
	Allow []string `json:"allow,omitempty"`
	Href  string   `json:"href,omitempty"`
	Links []Links  `json:"links,omitempty"`
}

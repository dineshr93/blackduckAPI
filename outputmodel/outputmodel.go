// SPDX-FileCopyrightText: 2023 Dinesh Ravi
//
// SPDX-License-Identifier: GPL-3.0-only

package outputmodel

type OutputModel struct {
	Items []Item
}

type Item struct {
	ScanSummary                 string                       `json:"scanSummary,omitempty"`
	URI                         string                       `json:"uri,omitempty"`
	CompositePath               CompositePath                `json:"compositePath,omitempty"`
	SnippetBomComponentRequests []SnippetBomComponentRequest `json:"snippetBomComponentRequests,omitempty"`
}
type CompositePath struct {
	Path                 string `json:"path,omitempty"`
	ArchiveContext       string `json:"archiveContext,omitempty"`
	CompositePathContext string `json:"compositePathContext,omitempty"`
	FileName             string `json:"fileName,omitempty"`
}
type SnippetBomComponentRequest struct {
	SourceStartLines []int  `json:"sourceStartLines,omitempty"`
	SourceEndLines   []int  `json:"sourceEndLines,omitempty"`
	Ignored          bool   `json:"ignored,omitempty"`
	HashID           string `json:"hashId,omitempty"`
	ReviewStatus     string `json:"reviewStatus,omitempty"`
	Component        string `json:"component,omitempty"`
	ComponentVersion string `json:"componentVersion,omitempty"`
}

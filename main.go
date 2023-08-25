package main

import (
	service "github.com/dineshr93/blackduckAPI/service"
)

func main() {
	// To activate
	// /bulk-snippet-bom-entries

	// projectversionlink := "https://blackduck.ebgroup.elektrobit.com/api/projects/210a7fcb-b49d-4cda-9302-c888ad562024/versions/ed598519-d0de-437d-8996-4ca2b932ea8c"
	projectversionlink := "https://blackduck.ebgroup.elektrobit.com/api/projects/210a7fcb-b49d-4cda-9302-c888ad562024/versions/81df0df9-aa6d-4750-a1e7-6a049c6b4967"
	// /scan-summaries
	codelocation := "https://blackduck.ebgroup.elektrobit.com/api/codelocations/3cf0a11a-e453-4d64-ba3b-5ee87aaaeabd"
	// boost
	matchTOComponent := "https://blackduck.ebgroup.elektrobit.com/api/components/4eac8f37-d9e5-4344-83d0-be0e9fd42a6a"
	matchToComponentVersion := "https://blackduck.ebgroup.elektrobit.com/api/components/4eac8f37-d9e5-4344-83d0-be0e9fd42a6a/versions/9b57e540-34c1-4f6c-89cd-a0137e1959b6"
	matchString := "boost"
	// bluetoothcomponent
	// matchTOComponent := "https://blackduck.ebgroup.elektrobit.com/api/components/771d0705-eff8-4865-9d20-81a17c072717"
	// matchToComponentVersion := "https://blackduck.ebgroup.elektrobit.com/api/components/771d0705-eff8-4865-9d20-81a17c072717/versions/22dfc3be-aaa3-44e8-a9f5-9bae149f4e42"
	// matchString := "Bluetooth2"
	entries := "13"

	hub := service.HUB{}
	hub.Loadinit()
	hub.LoadPVlinks(projectversionlink)
	hub.LoadSpecificCodeLocation(codelocation)
	hub.LoadMatchToComponentVersion(matchTOComponent, matchToComponentVersion)
	hub.SetScansummariesURL()

	hub.SetSourceBomEntries(entries)

	hub.PrepareBSBEBody(matchString)
	hub.MatchSnippetToGivenComponent()

}

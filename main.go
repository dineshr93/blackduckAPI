package main

import (
	"log"
	"os"

	"github.com/dineshr93/blackduckAPI/helpers"
	service "github.com/dineshr93/blackduckAPI/service"
)

func main() {
	if len(os.Args) != 8 {
		log.Println(os.Args)
		log.Println(len(os.Args))
		panic("Provide 1. Hub api url 2. hub token 3. projectversionlink 4. codelocation 5. matchToComponentVersion 6.matchString 7. EntriesToResolve ")

	}
	// projectversionlink := "https://blackduck.ebgroup.elektrobit.com/api/projects/210a7fcb-b49d-4cda-9302-c888ad562024/versions/ed598519-d0de-437d-8996-4ca2b932ea8c"
	// projectversionlink := "https://blackduck.ebgroup.elektrobit.com/api/projects/210a7fcb-b49d-4cda-9302-c888ad562024/versions/81df0df9-aa6d-4750-a1e7-6a049c6b4967"
	projectversionlink := os.Args[3]
	// /scan-summaries
	// codelocation := "https://blackduck.ebgroup.elektrobit.com/api/codelocations/3cf0a11a-e453-4d64-ba3b-5ee87aaaeabd"
	codelocation := os.Args[4]
	// boost
	// matchTOComponent := "https://blackduck.ebgroup.elektrobit.com/api/components/4eac8f37-d9e5-4344-83d0-be0e9fd42a6a"
	// matchToComponentVersion := "https://blackduck.ebgroup.elektrobit.com/api/components/4eac8f37-d9e5-4344-83d0-be0e9fd42a6a/versions/9b57e540-34c1-4f6c-89cd-a0137e1959b6"
	matchToComponentVersion := os.Args[5]

	matchToComponent := helpers.GetTextsUpto(matchToComponentVersion, "/versions")

	// matchString := "boost"
	matchString := os.Args[6]
	// bluetoothcomponent
	// matchTOComponent := "https://blackduck.ebgroup.elektrobit.com/api/components/771d0705-eff8-4865-9d20-81a17c072717"
	// matchToComponentVersion := "https://blackduck.ebgroup.elektrobit.com/api/components/771d0705-eff8-4865-9d20-81a17c072717/versions/22dfc3be-aaa3-44e8-a9f5-9bae149f4e42"
	// matchString := "Bluetooth2"
	// entries := "13"
	entries := os.Args[7]

	hub := service.HUB{}
	hub.Loadinit()
	hub.LoadPVlinks(projectversionlink)
	hub.LoadSpecificCodeLocation(codelocation)
	hub.LoadMatchToComponentVersion(matchToComponent, matchToComponentVersion)
	hub.SetScansummariesURL()

	hub.SetSourceBomEntries(entries)

	hub.PrepareBSBEBody(matchString)
	hub.MatchSnippetToGivenComponent()

}

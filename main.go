// SPDX-FileCopyrightText: 2023 Dinesh Ravi
//
// SPDX-License-Identifier: GPL-3.0-only

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
	projectversionlink := os.Args[3]
	codelocation := os.Args[4]
	matchToComponentVersion := os.Args[5]

	matchToComponent := helpers.GetTextsUpto(matchToComponentVersion, "/versions")

	matchString := os.Args[6]

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

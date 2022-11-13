package main

import (
	"fmt"

	"github.com/infamousjoeg/com.infamousjoeg.nhl-game-today.sdPlugin/pkg/nhlgametoday"
)

func main() {
	// Team ID 14 = Tampa Bay Lightning
	scheduleResponse := nhlgametoday.GetSchedule(14)
	awayTeamID := scheduleResponse.Dates[0].Games[0].Teams.Away.Team.ID
	homeTeamID := scheduleResponse.Dates[0].Games[0].Teams.Home.Team.ID
	awayTeam := nhlgametoday.GetTeam(awayTeamID)
	homeTeam := nhlgametoday.GetTeam(homeTeamID)
	awayTeamShortName := awayTeam.Teams[0].Abbreviation
	homeTeamShortName := homeTeam.Teams[0].Abbreviation
	fmt.Println(awayTeamShortName, "at", homeTeamShortName)
}

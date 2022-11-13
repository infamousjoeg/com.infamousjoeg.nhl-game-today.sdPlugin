package main

import (
	"fmt"

	"github.com/infamousjoeg/com.infamousjoeg.nhl-game-today.sdPlugin/pkg/nhlgametoday"
)

func main() {
	// Team ID 14 = Tampa Bay Lightning
	scheduleResponse := nhlgametoday.GetSchedule(14)
	awayTeam := nhlgametoday.GetTeam(scheduleResponse.Dates[0].Games[0].Teams.Away.Team.ID)
	homeTeam := nhlgametoday.GetTeam(scheduleResponse.Dates[0].Games[0].Teams.Home.Team.ID)
	fmt.Println(awayTeam.Teams[0].Abbreviation, "at", homeTeam.Teams[0].Abbreviation)
}

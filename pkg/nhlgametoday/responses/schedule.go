package responses

import "time"

// Schedule is the response from the NHL Stats API when requesting a team's schedule
type Schedule struct {
	Copyright    string `json:"copyright,omitempty"`
	TotalItems   int    `json:"totalItems,omitempty"`
	TotalEvents  int    `json:"totalEvents,omitempty"`
	TotalGames   int    `json:"totalGames,omitempty"`
	TotalMatches int    `json:"totalMatches,omitempty"`
	MetaData     struct {
		TimeStamp string `json:"timeStamp,omitempty"`
	} `json:"metaData,omitempty"`
	Wait  int `json:"wait,omitempty"`
	Dates []struct {
		Date         string `json:"date,omitempty"`
		TotalItems   int    `json:"totalItems,omitempty"`
		TotalEvents  int    `json:"totalEvents,omitempty"`
		TotalGames   int    `json:"totalGames,omitempty"`
		TotalMatches int    `json:"totalMatches,omitempty"`
		Games        []struct {
			GamePk   int       `json:"gamePk,omitempty"`
			Link     string    `json:"link,omitempty"`
			GameType string    `json:"gameType,omitempty"`
			Season   string    `json:"season,omitempty"`
			GameDate time.Time `json:"gameDate,omitempty"`
			Status   struct {
				AbstractGameState string `json:"abstractGameState,omitempty"`
				CodedGameState    string `json:"codedGameState,omitempty"`
				DetailedState     string `json:"detailedState,omitempty"`
				StatusCode        string `json:"statusCode,omitempty"`
				StartTimeTBD      bool   `json:"startTimeTBD,omitempty"`
			} `json:"status,omitempty"`
			Teams struct {
				Away struct {
					LeagueRecord struct {
						Wins   int    `json:"wins,omitempty"`
						Losses int    `json:"losses,omitempty"`
						Ot     int    `json:"ot,omitempty"`
						Type   string `json:"type,omitempty"`
					} `json:"leagueRecord,omitempty"`
					Score int `json:"score,omitempty"`
					Team  struct {
						ID   int    `json:"id,omitempty"`
						Name string `json:"name,omitempty"`
						Link string `json:"link,omitempty"`
					} `json:"team,omitempty"`
				} `json:"away,omitempty"`
				Home struct {
					LeagueRecord struct {
						Wins   int    `json:"wins,omitempty"`
						Losses int    `json:"losses,omitempty"`
						Ot     int    `json:"ot,omitempty"`
						Type   string `json:"type,omitempty"`
					} `json:"leagueRecord,omitempty"`
					Score int `json:"score,omitempty"`
					Team  struct {
						ID   int    `json:"id,omitempty"`
						Name string `json:"name,omitempty"`
						Link string `json:"link,omitempty"`
					} `json:"team,omitempty"`
				} `json:"home,omitempty"`
			} `json:"teams,omitempty"`
			Venue struct {
				ID   int    `json:"id,omitempty"`
				Name string `json:"name,omitempty"`
				Link string `json:"link,omitempty"`
			} `json:"venue,omitempty"`
			Content struct {
				Link string `json:"link,omitempty"`
			} `json:"content,omitempty"`
		} `json:"games,omitempty"`
		Events  []interface{} `json:"events,omitempty"`
		Matches []interface{} `json:"matches,omitempty"`
	} `json:"dates,omitempty"`
}

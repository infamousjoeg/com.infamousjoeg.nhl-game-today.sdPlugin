package responses

// Team is the response from the NHL Stats API when requesting a team's information
type Team struct {
	Copyright string `json:"copyright,omitempty"`
	Teams     []struct {
		ID    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Link  string `json:"link,omitempty"`
		Venue struct {
			Name     string `json:"name,omitempty"`
			Link     string `json:"link,omitempty"`
			City     string `json:"city,omitempty"`
			TimeZone struct {
				ID     string `json:"id,omitempty"`
				Offset int    `json:"offset,omitempty"`
				Tz     string `json:"tz,omitempty"`
			} `json:"timeZone,omitempty"`
		} `json:"venue,omitempty"`
		Abbreviation    string `json:"abbreviation,omitempty"`
		TeamName        string `json:"teamName,omitempty"`
		LocationName    string `json:"locationName,omitempty"`
		FirstYearOfPlay string `json:"firstYearOfPlay,omitempty"`
		Division        struct {
			ID           int    `json:"id,omitempty"`
			Name         string `json:"name,omitempty"`
			NameShort    string `json:"nameShort,omitempty"`
			Link         string `json:"link,omitempty"`
			Abbreviation string `json:"abbreviation,omitempty"`
		} `json:"division,omitempty"`
		Conference struct {
			ID   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
			Link string `json:"link,omitempty"`
		} `json:"conference,omitempty"`
		Franchise struct {
			FranchiseID int    `json:"franchiseId,omitempty"`
			TeamName    string `json:"teamName,omitempty"`
			Link        string `json:"link,omitempty"`
		} `json:"franchise,omitempty"`
		ShortName       string `json:"shortName,omitempty"`
		OfficialSiteURL string `json:"officialSiteUrl,omitempty"`
		FranchiseID     int    `json:"franchiseId,omitempty"`
		Active          bool   `json:"active,omitempty"`
	} `json:"teams,omitempty"`
}

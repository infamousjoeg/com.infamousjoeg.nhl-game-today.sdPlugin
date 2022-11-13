package nhlgametoday

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/infamousjoeg/com.infamousjoeg.nhl-game-today.sdPlugin/pkg/nhlgametoday/responses"
)

const scheduleURL = "https://statsapi.web.nhl.com/api/v1/schedule"
const teamsURL = "https://statsapi.web.nhl.com/api/v1/teams"

// sendRequest will send a request and return the response from the NHL Stats API
func sendRequest(method, url string, body io.Reader) ([]byte, error) {
	// Create http client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Check status code
	if resp.StatusCode >= 300 {
		return nil, err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)

	var reader io.Reader
	reader, err = gzip.NewReader(tee)
	if err != nil {
		reader = &buf
	}
	responseBody, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// GetSchedule will call the NHL Stats API and return the results for the team ID provided
func GetSchedule(teamID int) *responses.Schedule {
	if teamID == 0 {
		return nil
	}

	url, _ := url.Parse(scheduleURL)
	queryParams := url.Query()
	queryParams.Add("teamId", strconv.Itoa(teamID))
	url.RawQuery = queryParams.Encode()

	response, err := sendRequest("GET", url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal response
	ScheduleResponse := responses.Schedule{}
	err = json.Unmarshal(response, &ScheduleResponse)

	return &ScheduleResponse
}

// GetTeam will call the NHL Stats API and return the results for the team ID provided
func GetTeam(teamID int) *responses.Team {
	if teamID == 0 {
		return nil
	}

	url, _ := url.Parse(teamsURL + "/" + strconv.Itoa(teamID))

	response, err := sendRequest("GET", url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal response
	TeamResponse := responses.Team{}
	err = json.Unmarshal(response, &TeamResponse)

	return &TeamResponse
}

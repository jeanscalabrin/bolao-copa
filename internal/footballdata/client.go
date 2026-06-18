package footballdata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	apiKey string
	http   *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) GetWorldCupMatches() ([]MatchResponse, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.football-data.org/v4/competitions/WC/matches", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Auth-Token", c.apiKey)

	resp, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("football-data returned %d",
			resp.StatusCode)
	}

	var result MatchesResponse

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result.Matches, nil

}

package destinyapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL string = "https://www.bungie.net/Platform"

type DestinyAPIClient struct {
	APIKey   string
	Endpoint string
}

func NewDestinyAPISession(token string) *DestinyAPIClient {
	return &DestinyAPIClient{
		APIKey:   token,
		Endpoint: baseURL,
	}
}

type DestinyUser struct {
	membershipid int64 `json:"membershipid"`
}

func (dac *DestinyAPIClient) GetDestinyUser(id string) (*DestinyUser, error) {
	url := fmt.Sprintf("%s//Destiny2/3/Profile/%s/?components=100", dac.Endpoint, id)
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	fmt.Println(req.URL)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-API-Key", dac.APIKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bytes))
	du := &DestinyUser{}
	err = json.Unmarshal(bytes, &du)
	if err != nil {
		return nil, err
	}
	return du, err
}

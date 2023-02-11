package d2client

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

package discordwebapi

import (
	"encoding/json"
	"fmt"
	"github.com/SomethingBot/discordingestor/discord/libinfo"
	"net/http"
	"net/url"
)

type GatewayWebsocketInformationSessionStartLimit struct {
	Total          int `json:"total"`
	Remaining      int `json:"remaining"`
	ResetAfter     int `json:"reset_after"`
	MaxConcurrency int `json:"max_concurrency"`
}

type GatewayWebsocketInformation struct {
	Url               string                                       `json:"url"`
	Shards            int                                          `json:"shards"`
	SessionStartLimit GatewayWebsocketInformationSessionStartLimit `json:"session_start_limit"`
}

//GetGatewayWebsocketInformation for a bot to connect to the Discord websocket
func GetGatewayWebsocketInformation(discordApiGatewayURL string, apiKey string) (GatewayWebsocketInformation, error) {
	if discordApiGatewayURL == "" {
		discordApiGatewayURL = "https://discord.com/api/gateway/bot"
	}

	req, err := http.NewRequest("GET", discordApiGatewayURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", libinfo.BotUserAgent)

	return nil
}

//GetGatewayWebsocketURI returns the current Discord Gateway WSS URL, pass discordApiGatewayURL as "" to use default  //todo: make it so test doesn't have to hit server
func GetGatewayWebsocketURI(discordApiGatewayURL string) (url.URL, error) {
	if discordApiGatewayURL == "" {
		discordApiGatewayURL = "https://discord.com/api/gateway"
	}

	req, err := http.NewRequest("GET", discordApiGatewayURL, nil)
	if err != nil {
		return url.URL{}, err
	}
	req.Header.Set("User-Agent", libinfo.BotUserAgent)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return url.URL{}, err
	}

	urlJson := struct {
		Url string `json:"url"`
	}{}

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&urlJson)
	if err != nil {
		err2 := resp.Body.Close()
		if err2 != nil {
			return url.URL{}, fmt.Errorf("could not close body (%v), after error (%w)", err2, err)
		}
		return url.URL{}, err
	}

	err = resp.Body.Close()
	if err != nil {
		return url.URL{}, err
	}

	uri, err := url.ParseRequestURI(urlJson.Url)
	if err != nil {
		return url.URL{}, err
	}

	return *uri, err
}

package weather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	Host       string
	HttpClient *http.Client
	ApiKey     string
}

func NewClient(host string, apiKey string, timeout time.Duration) *Client {
	client := &http.Client{
		Timeout: timeout,
	}
	return &Client{
		Host:       host,
		HttpClient: client,
		ApiKey:     apiKey,
	}
}

func (c *Client) getResponse(lat string, long string, appid string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.Host, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("lat", lat)
	q.Add("lon", long)
	q.Add("appid", appid)

	req.URL.RawQuery = q.Encode()
	return c.HttpClient.Do(req)
}

func (c *Client) GetCurrentWeatherResponse(lat string, long string) (*WeatherResponse, error) {
	resp, err := c.getResponse(lat, long, c.ApiKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var weatherResponse WeatherResponse
	if err = json.Unmarshal(body, &weatherResponse); err != nil {
		return nil, err
	}
	return &weatherResponse, nil

}

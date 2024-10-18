package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpclient http.Client
}

func NewClient() *Client {
	return &Client{
		httpclient: http.Client{
			Timeout: time.Minute,
		},
	}

}

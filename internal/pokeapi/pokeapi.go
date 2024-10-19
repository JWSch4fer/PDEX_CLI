package pokeapi

import (
	"PDEX_CLI/internal/pokecache"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpclient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpclient: http.Client{
			Timeout: time.Minute,
		},
	}

}

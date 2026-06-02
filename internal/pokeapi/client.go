package pokeapi

import (
	"net/http"
	"time"

	"github.com/cmilliron/bootdev-pokedex-go/internal/pokecache"
)

// Client -
type Client struct {
	httpClient	http.Client
	cache 		*pokecache.Cache
}

// NewClient -
func NewClient(duration time.Duration, cacheInternval time.Duration) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: duration,
		},
		cache: pokecache.NewCache(cacheInternval),
	}
}
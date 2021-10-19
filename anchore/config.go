package anchore

import (
	"github.com/excellaco/anchore-client-go/client"
)

type Config struct {
	ApiURL   string
	Username string
	Password string
}

func newAnchoreClient(c *Config) client.AnchoreClient {
	return client.NewClient(&c.ApiURL, &c.Username, &c.Password)
}

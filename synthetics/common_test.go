package synthetics

import (
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func checkAPIServerConnection(t *testing.T) {
	apiURL, ok := os.LookupEnv("KTAPI_URL")
	if !ok {
		t.Fatal("KTAPI_URL env variable not set")
	}

	_, err := http.Get(apiURL)
	if err != nil {
		t.Fatalf("failed to connect to the API Server on URL %q: %v", apiURL, err)
	}
}

func providerFactories() map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"kentik-synthetics": func() (*schema.Provider, error) {
			return NewProvider(), nil
		},
	}
}

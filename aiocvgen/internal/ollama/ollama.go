package ollama

import (
	"aiocvgen/internal/config"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/ollama/ollama/api"
)

// Module-wide client object for interacting with Ollama
var client *api.Client

// Setups a global client to be used with api
func SetupClient() {
	endpoint := config.Get().OllamaUrl
	if !strings.HasPrefix(endpoint, "http://") && !strings.HasPrefix(endpoint, "https://") {
		endpoint = "http://" + endpoint
	}
	u, _ := url.Parse(endpoint)
	client = api.NewClient(u, http.DefaultClient)
	fmt.Printf("Available models: \n- %s\n\n", strings.Join(GetModels(), "\n- "))
}

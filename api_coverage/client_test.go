package api_coverage

import (
	"net/http"
	"os"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/will7200/go-wakatime/client"
)

var (
	client     *apiclient.Wakatime
	apiKeyAuth = httptransport.APIKeyAuth("api_key", "query", os.Getenv("WAKATIME_API_KEY"))
)

func init() {
	setupClient()
}

func setupClient() *apiclient.Wakatime {
	defaultT := apiclient.DefaultTransportConfig()
	runtime := httptransport.NewWithClient(defaultT.Host, defaultT.BasePath, defaultT.Schemes,
		&http.Client{Timeout: 10 * time.Second, Transport: http.DefaultTransport})
	client = apiclient.New(runtime, strfmt.Default)
	return client
}

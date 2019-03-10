package api_coverage

import (
	"github.com/sony/gobreaker"
	"os"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/will7200/go-wakatime/client"
)

var (
	client            *apiclient.Wakatime
	globalTransporter *Transport

	apiKeyAuth = httptransport.APIKeyAuth("api_key", "query", os.Getenv("WAKATIME_API_KEY"))
)

func init() {
	setupClient()
}

func setupClient() *apiclient.Wakatime {

	// Setup HTTP Client for Testing

	var (
		breaker    = newCircuitBreaker()
		transport  = NewTransport(breaker)
		httpClient = NewClient(WithRoundTripper(transport))
	)
	globalTransporter = transport

	// Setup API Client
	defaultT := apiclient.DefaultTransportConfig()
	runtime := httptransport.NewWithClient(
		defaultT.Host, defaultT.BasePath, defaultT.Schemes,
		httpClient)
	client = apiclient.New(runtime, strfmt.Default)
	return client
}

func newCircuitBreaker() *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "HTTP Client",
		Timeout: time.Second * 30,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			// do smth when circuit breaker trips.
		},
	})
}

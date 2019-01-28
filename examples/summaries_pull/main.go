package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/kr/pretty"
	apiclient "github.com/will7200/go-wakatime/client"
	user2 "github.com/will7200/go-wakatime/client/user"
)

func main() {
	defaultT := apiclient.DefaultTransportConfig()
	runtime := httptransport.NewWithClient(defaultT.Host, defaultT.BasePath, defaultT.Schemes,
		&http.Client{Timeout: 10 * time.Second, Transport: http.DefaultTransport})
	client := apiclient.New(runtime, strfmt.Default)

	apiKeyAuth := httptransport.APIKeyAuth("api_key", "query", os.Getenv("WAKATIME_API_KEY"))

	user, err := client.User.User(nil, apiKeyAuth)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v\n", pretty.Formatter(user))

	params := user2.SummariesParams{
		Start:   strfmt.Date(time.Now().Add(-1 * time.Hour * 24)),
		End:     strfmt.Date(time.Now()),
		User:    "current",
		Context: context.Background(),
	}
	use, err := client.User.Summaries(&params, apiKeyAuth)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v\n", pretty.Formatter(use))
}

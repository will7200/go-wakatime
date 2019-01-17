package main

import (
	"fmt"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/kr/pretty"
	apiclient "github.com/will7200/wakatime/client"
)

func main() {
	client := apiclient.NewHTTPClient(nil)

	apiKeyAuth := httptransport.APIKeyAuth("api_key", "query", os.Getenv("WAKATIME_API_KEY"))

	user, err := client.User.CurrentUser(nil, apiKeyAuth)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v\n", pretty.Formatter(user))

	use, err := client.Stats.CurrentUserStats(nil, apiKeyAuth)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v\n", pretty.Formatter(use))
}

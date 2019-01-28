# go-wakatime api
[![GoDoc](https://godoc.org/github.com/will7200/go-wakatime?status.svg)](http://godoc.org/github.com/will7200/go-wakatime)
[![GolangCI](https://golangci.com/badges/github.com/will7200/go-wakatime.svg)](https://golangci.com)
[![Go Report Card](https://goreportcard.com/badge/github.com/will7200/go-wakatime)](https://goreportcard.com/report/github.com/will7200/go-wakatime)

This package contains a golang implementation of Wakatime's API V1. I didn't like the other implementations of go-wakatime so I made this one.

Generated from swagger spec for the most part except for some exceptions
1. The swagger spec has not default date parse for the likes of 01/31/2019, so the internal package houses a near replica
of the date type the strfmt that open-api has.
2. Some parts have not been full tested so if you notice something or have an error open an issue.
3. go-swagger has not be pleasant with custom types as I have to make the whole struct type, which i tried to avoid.
Luckily, wakatime used a lot of formats the openapi has standardized in the strfmt package.
## Example

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-openapi/strfmt"
	user2 "github.com/will7200/go-wakatime/client/user"
	"net/http"
	"os"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/kr/pretty"
	apiclient "github.com/will7200/go-wakatime/client"
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

```



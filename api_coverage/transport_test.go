package api_coverage

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"net"
	"net/http"
	"net/url"
	"time"
)

type (
	// ClientOption is the request options.
	ClientOption func(*http.Client)

	// Client is the application http request.
	Client struct {
		client *http.Client
	}
)

// WithRoundTripper receives the http.RoundTripper implementation.
func WithRoundTripper(rt http.RoundTripper) ClientOption {
	return func(r *http.Client) {
		r.Transport = rt
	}
}

// NewClient returns a new configured Client.
func NewClient(opts ...ClientOption) *http.Client {
	r := new(http.Client)
	for _, o := range opts {
		o(r)
	}
	return r
}

type (
	// Breaker is the http circuit breaker.
	Breaker interface {
		// Execute runs the given request if the circuit breaker is closed or half-open states.
		// An error is instantly returned when the circuit breaker is tripped.
		Execute(func() (interface{}, error)) (interface{}, error)
	}

	// Transport is the application http transport.
	Transport struct {
		tripper http.RoundTripper
		breaker Breaker

		// contains all the url requests made
		urlRequests *arraylist.List
	}

	RequestMade struct {
		Method        string
		URL           url.URL
		ErrorOccurred bool
	}
)

// NewTransport returns a new configured Transport.
func NewTransport(cb Breaker) *Transport {
	t := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 90 * time.Second,
			DualStack: true,
		}).DialContext,
	}

	return &Transport{
		tripper:     t,
		breaker:     cb,
		urlRequests: arraylist.New(),
	}
}

// RoundTrip decorates tripper.RoundTrip with a circuit breaker.
// An error is returned if the circuit breaker rejects the request.
func (t *Transport) RoundTrip(r *http.Request) (_ *http.Response, err error) {
	rm := RequestMade{Method: r.Method, URL: *r.URL, ErrorOccurred: false}

	defer func() {
		if err != nil {
			rm.ErrorOccurred = true
		}
		t.urlRequests.Add(rm)
	}()

	res, err := t.breaker.Execute(func() (interface{}, error) {
		res, err := t.tripper.RoundTrip(r)
		if err != nil {
			return nil, err
		}

		if res != nil && res.StatusCode >= http.StatusInternalServerError {
			return res, fmt.Errorf("http response error: %v", res.StatusCode)
		}

		return res, err
	})

	if err != nil {
		return nil, err
	}

	return res.(*http.Response), err
}

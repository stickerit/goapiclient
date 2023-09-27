package goapiclient

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/stickerit/goapiclient/basicauth"
)

type ApiClient struct {
	username string
	apiKey   string
	instance string
	endpoint string
	gcl      graphql.Client
}

type ApiClientOpts struct {
	Endpoint *string
	Client   *http.Client
}

func New(username, apiKey, instance string) *ApiClient {
	return NewWithOpts(username, apiKey, instance, ApiClientOpts{})
}

func NewWithOpts(username, apiKey, instance string, opts ApiClientOpts) *ApiClient {
	acl := ApiClient{
		username: username,
		apiKey:   apiKey,
		instance: instance,
		endpoint: "https://cs-api.live.stickerit.co/graphql",
	}

	if opts.Endpoint != nil {
		acl.endpoint = *opts.Endpoint
	}

	cl := &http.Client{
		Timeout: time.Second * 300, // handle slow uploads
	}

	if opts.Client != nil {
		cl = opts.Client
	}

	extraHeaders := url.Values{}
	extraHeaders.Set("x-store-hash", instance)

	basicAuthClient := basicauth.NewBasicAuthClient(cl, username, apiKey, extraHeaders)

	acl.gcl = graphql.NewClient(acl.endpoint, basicAuthClient)

	return &acl
}

func (acl *ApiClient) CreateOrder(ctx context.Context, order CreateOrderRequest) (*CreateOrderResponse, error) {
	return CreateOrder(ctx, acl.gcl, order)
}

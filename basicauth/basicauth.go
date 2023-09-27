package basicauth

import (
	"net/http"
	"net/url"
)

type BasicAuthClient struct {
	cl           *http.Client
	username     string
	password     string
	extraHeaders url.Values
}

func (ba *BasicAuthClient) Do(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(ba.username, ba.password)

	for k, v := range ba.extraHeaders {
		req.Header[k] = v
	}

	return ba.cl.Do(req)
}

func NewBasicAuthClient(cl *http.Client, username, password string, extraHeaders url.Values) *BasicAuthClient {
	return &BasicAuthClient{
		cl:           cl,
		username:     username,
		password:     password,
		extraHeaders: extraHeaders,
	}
}

type BearerClient struct {
	cl           *http.Client
	token        string
	extraHeaders url.Values
}

func (ba *BearerClient) Do(req *http.Request) (*http.Response, error) {
	for k, v := range ba.extraHeaders {
		req.Header[k] = v
	}

	req.Header.Set("authorization", "Bearer "+ba.token)

	return ba.cl.Do(req)
}

func NewBearerClient(cl *http.Client, token string, extraHeaders url.Values) *BearerClient {
	return &BearerClient{
		cl:           cl,
		token:        token,
		extraHeaders: extraHeaders,
	}
}

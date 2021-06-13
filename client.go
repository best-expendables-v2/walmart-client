package walmart_client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/schema"
)

type Client struct {
	httpClient    *http.Client
	beforeRequest []func(ctx context.Context, c *Client) error
	Config        Config
	Token         *Token

	InventoryService
	AuthService
	OrderService
}

func NewClient(client *http.Client, conf Config, callbacks ...Callback) *Client {
	c := &Client{
		httpClient: client,
		Config:     conf,
	}

	c.InventoryService = NewInventoryService(c)
	c.AuthService = NewAuthService(c)
	c.OrderService = NewOrderService(c)

	for _, callback := range callbacks {
		callback(c)
	}
	return c
}

type Token struct {
	AccessToken string
	ExpireAt    time.Time
}

func (c *Client) doRequest(ctx context.Context, method string, uri string, options interface{}, content interface{}) (*Response, error) {
	u := c.getURL(uri)
	var err error
	if options != nil {
		var optionsQuery url.Values
		var schemaEncoder = schema.NewEncoder()
		err = schemaEncoder.Encode(options, optionsQuery)
		if err != nil {
			return nil, err
		}

		for k, values := range u.Query() {
			for _, v := range values {
				optionsQuery.Add(k, v)
			}
		}
		u.RawQuery = optionsQuery.Encode()
	}
	var bodyBytes []byte
	if content != nil {
		bodyBytes, err = json.Marshal(content)
		if err != nil {
			return nil, err
		}
	} else {
		bodyBytes = nil
	}
	for _, beforeRequest := range c.beforeRequest {
		beforeRequest(ctx, c)
	}
	return c.sendRequest(ctx, method, u.String(), bodyBytes, c.GetRequestHeaders())
}

func (c Client) sendRequest(ctx context.Context, method string, api string, content []byte, headers http.Header) (*Response, error) {
	request, err := http.NewRequest(method, api, bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx)
	request.Header = headers
	resp, err := c.httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rsp := &Response{
		StatusCode: resp.StatusCode,
		Body:       responseBody,
	}
	if http.StatusOK > rsp.StatusCode || rsp.StatusCode > http.StatusMultipleChoices {
		return rsp, errors.New(fmt.Sprintf("An error occurred when call request: %s", string(rsp.Body)))
	}
	return rsp, nil
}

func (c Client) GetHost() string {
	if c.Config.SandboxMode {
		return BaseUrlSandbox
	}
	return BaseUrlProd
}

func (c Client) getURL(uri string) *url.URL {
	host := c.GetHost()
	baseURL, err := url.Parse(host)
	if err != nil {
		panic(err)
	}

	rel, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	return baseURL.ResolveReference(rel)
}

func (c *Client) Get(ctx context.Context, uri string, options interface{}) (*Response, error) {
	return c.doRequest(ctx, http.MethodGet, uri, options, nil)
}

func (c *Client) Post(ctx context.Context, uri string, content interface{}) (*Response, error) {
	return c.doRequest(ctx, http.MethodPost, uri, nil, content)
}

func (c *Client) Put(ctx context.Context, uri string, content interface{}) (*Response, error) {
	return c.doRequest(ctx, http.MethodPut, uri, nil, content)
}

func (c *Client) Delete(ctx context.Context, uri string, content interface{}) (*Response, error) {
	return c.doRequest(ctx, http.MethodDelete, uri, nil, nil)
}

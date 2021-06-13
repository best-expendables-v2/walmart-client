package walmart_client

import "context"

type Callback func(c *Client)

func BeforeRequest(funcs ...func(ctx context.Context, c *Client) error) Callback {
	return func(c *Client) {
		c.beforeRequest = append(c.beforeRequest, funcs...)
	}
}

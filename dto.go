package walmart_client

import "encoding/json"

type Response struct {
	StatusCode int
	Body       []byte
}

func (r Response) Decode(output interface{}) error {
	return json.Unmarshal(r.Body, output)
}

package walmart_client

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

func (c Client) GetRequestHeaders() http.Header {
	authKey := c.Config.ClientID + ":" + c.Config.ClientSecret
	correlationID, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	headers := http.Header{
		"WM_SVC.NAME":           []string{c.Config.WMServiceName},
		"WM_QOS.CORRELATION_ID": []string{correlationID.String()},
		"Authorization":         []string{"Basic " + base64.StdEncoding.EncodeToString([]byte(authKey))},
		"Accept":                []string{RequestAccept},
		"Content-Type":          []string{RequestContentType},
	}
	if c.Token != nil {
		headers.Add("WM_SEC.ACCESS_TOKEN", c.Token.AccessToken)
	}
	return headers
}

func (c *Client) SetAccessToken(accessToken string, expireAt *string) {
	c.Token = &Token{
		AccessToken: accessToken,
	}
	if expireAt != nil {
		nanoTime, err := strconv.ParseInt(*expireAt, 0, 64)
		if err != nil {
			panic(err)
		}
		c.Token.ExpireAt = time.Unix(nanoTime / 1000, 0)
	}
}

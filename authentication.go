package walmart_client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

const (
	AuthenticateBasePath   = "/v3/token"
	GetTokenDetailBasePath = "/v3/token/detail"
)

type AuthService interface {
	Authenticate(ctx context.Context) (*AuthResponse, error)
	GetTokenDetail(ctx context.Context) (*GetTokenDetailResponse, error)
}

type authService struct {
	client *Client
}

func NewAuthService(client *Client) AuthService {
	return &authService{client: client}
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (c authService) Authenticate(ctx context.Context) (*AuthResponse, error) {
	content := "grant_type=client_credentials"
	url := c.client.getURL(GetAccessTokenUri)
	headers := c.client.GetRequestHeaders()
	headers.Set("Content-Type", RequestAuthContentType)
	rsp, err := c.client.sendRequest(ctx, http.MethodPost, url.String(), []byte(content), headers)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusOK {
		return nil, errors.New("Fail to get access token")
	}
	var response AuthResponse
	if err = json.Unmarshal(rsp.Body, &response); err != nil {
		return nil, errors.Wrap(err, "error when parse response data")
	}
	return &response, nil
}

type GetTokenDetailResponse struct {
	ExpireAt string `json:"expire_at"`
	IssuedAt string `json:"issued_at"`
	IsValid  bool   `json:"is_valid"`
	Scopes   struct {
		Reports   string `json:"reports"`
		Item      string `json:"item"`
		Price     string `json:"price"`
		Lagtime   string `json:"lagtime"`
		Feeds     string `json:"feeds"`
		Returns   string `json:"returns"`
		Orders    string `json:"orders"`
		Inventory string `json:"inventory"`
		Content   string `json:"content"`
	} `json:"scopes"`
}

func (c authService) GetTokenDetail(ctx context.Context) (*GetTokenDetailResponse, error) {
	url := c.client.getURL(GetTokenDetailBasePath)
	response, err := c.client.sendRequest(ctx, http.MethodGet, url.String(), nil, c.client.GetRequestHeaders())
	if err != nil {
		return nil, err
	}
	var res GetTokenDetailResponse
	err = response.Decode(&res)
	return &res, err
}

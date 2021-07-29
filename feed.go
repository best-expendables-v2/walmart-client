package walmart_client

import (
	"context"
	"github.com/best-expendables-v2/walmart-client/feed_dto"
)

const (
	FeedBasePath = "/v3/feeds"
)

type FeedService interface {
	ConvertItemsForWFS(ctx context.Context, content feed_dto.ConvertItemsPayload) (*feed_dto.ConvertItemsResponse, error)
}

type feedService struct {
	client *Client
}

func NewFeedService(client *Client) FeedService {
	return &feedService{client: client}
}

func (s feedService) ConvertItemsForWFS(ctx context.Context, content feed_dto.ConvertItemsPayload) (*feed_dto.ConvertItemsResponse, error) {
	response, err := s.client.Post(ctx, FeedBasePath+"?feedType=OMNI_WFS", content)
	if err != nil {
		return nil, err
	}
	var res feed_dto.ConvertItemsResponse
	err = response.Decode(&res)
	return &res, err
}

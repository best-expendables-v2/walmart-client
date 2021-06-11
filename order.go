package walmart_client

import (
	"context"
	"fmt"

	"github.com/best-expendables-v2/walmart-client/order_dto"
)

const (
	AcknowledgeOrderBasePath = "/v3/orders/%s/acknowledge"
	ShipOrderBasePath        = "/v3/orders/%s/shipping"
	CancelOrderBasePath      = "/v3/orders/%s/cancel"
)

type OrderService interface {
	Acknowledge(ctx context.Context, purchaseOrderID string) (*order_dto.OrderResponse, error)
	Ship(ctx context.Context, purchaseOrderID string, payload order_dto.OrderShipmentPayload) (*order_dto.OrderResponse, error)
	Cancel(ctx context.Context, purchaseOrderID string, payload order_dto.OrderCancellationPayload) (*order_dto.OrderResponse, error)
}

type orderService struct {
	client *Client
}

func NewOrderService(client *Client) OrderService {
	return &orderService{client: client}
}

func (s orderService) Acknowledge(ctx context.Context, purchaseOrderID string) (*order_dto.OrderResponse, error) {
	response, err := s.client.Post(ctx, fmt.Sprintf(AcknowledgeOrderBasePath, purchaseOrderID), nil)
	if err != nil {
		return nil, err
	}
	var res order_dto.OrderResponse
	err = response.Decode(&res)
	return &res, err
}

func (s orderService) Ship(ctx context.Context, purchaseOrderID string, payload order_dto.OrderShipmentPayload) (*order_dto.OrderResponse, error) {
	response, err := s.client.Post(ctx, fmt.Sprintf(ShipOrderBasePath, purchaseOrderID), payload)
	if err != nil {
		return nil, err
	}
	var res order_dto.OrderResponse
	err = response.Decode(&res)
	return &res, err
}

func (s orderService) Cancel(ctx context.Context, purchaseOrderID string, payload order_dto.OrderCancellationPayload) (*order_dto.OrderResponse, error) {
	response, err := s.client.Post(ctx, fmt.Sprintf(CancelOrderBasePath, purchaseOrderID), payload)
	if err != nil {
		return nil, err
	}
	var res order_dto.OrderResponse
	err = response.Decode(&res)
	return &res, err
}

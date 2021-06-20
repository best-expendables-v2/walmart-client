package walmart_client

import (
	"context"

	"github.com/best-expendables-v2/walmart-client/inventory_dto"
)

const (
	GetInventoryBasePath    = "/v3/inventory"
	UpdateInventoryBasePath = "/v3/inventory"
)

type InventoryService interface {
	GetInventory(ctx context.Context, params inventory_dto.InventoryGetParams) (*inventory_dto.InventoryResponse, error)
	UpdateInventory(ctx context.Context, content inventory_dto.InventoryUpdatePayload) (*inventory_dto.InventoryResponse, error)
}

type inventoryService struct {
	client *Client
}

func NewInventoryService(client *Client) InventoryService {
	return &inventoryService{client: client}
}

func (s inventoryService) GetInventory(ctx context.Context, params inventory_dto.InventoryGetParams) (*inventory_dto.InventoryResponse, error) {
	response, err := s.client.Get(ctx, GetInventoryBasePath, params)
	if err != nil {
		return nil, err
	}
	var res inventory_dto.InventoryResponse
	err = response.Decode(&res)
	return &res, err
}

func (s inventoryService) UpdateInventory(ctx context.Context, content inventory_dto.InventoryUpdatePayload) (*inventory_dto.InventoryResponse, error) {
	response, err := s.client.Post(ctx, UpdateInventoryBasePath, content)
	if err != nil {
		return nil, err
	}
	var res inventory_dto.InventoryResponse
	err = response.Decode(&res)
	return &res, err
}

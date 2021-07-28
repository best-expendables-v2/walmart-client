package walmart_client

import (
	"context"
	"github.com/best-expendables-v2/walmart-client/fulfillment_dto"
)

const (
	FulfillmentBasePath                      = "/v3/fulfillment"
	FulfillmentInboundShipmentBasePath       = FulfillmentBasePath + "/inbound-shipments"
	FulfillmentInboundShipmentErrorsBasePath = FulfillmentBasePath + "/inbound-shipment-errors"
	FulfillmentInboundShipmentItemsBasePath  = FulfillmentBasePath + "/inbound-shipment-items"
	FulfillmentShipmentQuantitiesBasePath    = FulfillmentBasePath + "/shipment-quantities"
	FulfillmentShipmentLabelBasePath         = FulfillmentBasePath + "/label"
	FulfillmentShipmentTrackingBasePath      = FulfillmentBasePath + "/shipment-tracking"
)

type FulfillmentService interface {
	CreateInboundShipment(ctx context.Context, content fulfillment_dto.CreateInboundShipmentPayload) (*fulfillment_dto.CreateInboundShipmentResponse, error)
	GetInboundShipmentErrors(ctx context.Context, params fulfillment_dto.GetInboundShipmentErrorsParams) (*fulfillment_dto.GetInboundShipmentErrorsResponse, error)
	GetShipments(ctx context.Context, params fulfillment_dto.GetShipmentParams) (*fulfillment_dto.GetShipmentResponse, error)
	GetInboundShipmentItems(ctx context.Context, params fulfillment_dto.GetInboundShipmentItemsParams) (*fulfillment_dto.GetInboundShipmentItemsResponse, error)
	UpdateShipmentQuantities(ctx context.Context, params fulfillment_dto.UpdateShipmentQuantitiesPayload) (*fulfillment_dto.BasicResponse, error)
	CreateInboundShipmentLabel(ctx context.Context, shipmentID string) error
	UpdateShipmentTracking(ctx context.Context, content fulfillment_dto.UpdateShipmentTrackingPayload) (*fulfillment_dto.BasicResponse, error)
	CancelInboundShipment(ctx context.Context, inboundOrderId string) (*fulfillment_dto.BasicResponse, error)
}

type fulfillmentService struct {
	client *Client
}

func NewFulfillmentService(client *Client) FulfillmentService {
	return &fulfillmentService{client: client}
}

func (s fulfillmentService) CreateInboundShipment(ctx context.Context, content fulfillment_dto.CreateInboundShipmentPayload) (*fulfillment_dto.CreateInboundShipmentResponse, error) {
	response, err := s.client.Post(ctx, FulfillmentInboundShipmentBasePath, content)
	if err != nil {
		return nil, err
	}
	var res fulfillment_dto.CreateInboundShipmentResponse
	err = response.Decode(&res)
	return &res, err
}

func (s fulfillmentService) GetInboundShipmentErrors(ctx context.Context, params fulfillment_dto.GetInboundShipmentErrorsParams) (*fulfillment_dto.GetInboundShipmentErrorsResponse, error) {
	response, err := s.client.Get(ctx, FulfillmentInboundShipmentErrorsBasePath, params)
	if err != nil {
		return nil, err
	}
	var res fulfillment_dto.GetInboundShipmentErrorsResponse
	err = response.Decode(&res)
	return &res, err
}

func (s fulfillmentService) GetShipments(ctx context.Context, params fulfillment_dto.GetShipmentParams) (*fulfillment_dto.GetShipmentResponse, error) {
	response, err := s.client.Get(ctx, FulfillmentInboundShipmentBasePath, params)
	if err != nil {
		return nil, err
	}
	var res fulfillment_dto.GetShipmentResponse
	err = response.Decode(&res)
	return &res, err
}

func (s fulfillmentService) GetInboundShipmentItems(ctx context.Context, params fulfillment_dto.GetInboundShipmentItemsParams) (*fulfillment_dto.GetInboundShipmentItemsResponse, error) {
	response, err := s.client.Get(ctx, FulfillmentInboundShipmentItemsBasePath, params)
	if err != nil {
		return nil, err
	}
	var res fulfillment_dto.GetInboundShipmentItemsResponse
	err = response.Decode(&res)
	return &res, err
}

func (s fulfillmentService) UpdateShipmentQuantities(ctx context.Context, params fulfillment_dto.UpdateShipmentQuantitiesPayload) (*fulfillment_dto.BasicResponse, error) {
	response, err := s.client.Put(ctx, FulfillmentShipmentQuantitiesBasePath, params)
	if err != nil {
		return nil, err
	}
	var res fulfillment_dto.BasicResponse
	err = response.Decode(&res)
	return &res, err
}

func (s fulfillmentService) CreateInboundShipmentLabel(ctx context.Context, shipmentID string) error {
	_, err := s.client.Get(ctx, FulfillmentShipmentLabelBasePath+"/label/"+shipmentID, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s fulfillmentService) UpdateShipmentTracking(ctx context.Context, content fulfillment_dto.UpdateShipmentTrackingPayload) (*fulfillment_dto.BasicResponse, error) {
	response, err := s.client.Post(ctx, FulfillmentShipmentTrackingBasePath, content)
	if err != nil {
		return nil, err
	}
	var res fulfillment_dto.BasicResponse
	err = response.Decode(&res)
	return &res, err
}

func (s fulfillmentService) CancelInboundShipment(ctx context.Context, inboundOrderId string) (*fulfillment_dto.BasicResponse, error) {
	response, err := s.client.Delete(ctx, FulfillmentInboundShipmentBasePath+"/"+inboundOrderId, nil)
	if err != nil {
		return nil, err
	}
	var res fulfillment_dto.BasicResponse
	err = response.Decode(&res)
	return &res, err
}

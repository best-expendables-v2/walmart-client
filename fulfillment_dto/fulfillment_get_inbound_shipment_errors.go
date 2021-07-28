package fulfillment_dto

import "time"

type GetInboundShipmentErrorsParams struct {
	PaginationParams
	ShipmentID string `json:"shipmentId" schema:"shipmentId"`
}

type GetInboundShipmentErrorsResponse struct {
	Headers struct {
		PaginationParams
		TotalCount int `json:"totalCount"`
	} `json:"headers"`
	Payload []GetInboundShipmentErrorsResponsePayload `json:"payload"`
}
type GetInboundShipmentErrorsResponsePayload struct {
	InboundOrderID string      `json:"inboundOrderId"`
	CreatedDate    time.Time   `json:"createdDate"`
	ReturnAddress  Address     `json:"returnAddress"`
	OrderItems     []OrderItem `json:"orderItems"`
	Errors         []Error     `json:"errors"`
}

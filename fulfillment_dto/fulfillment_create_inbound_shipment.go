package fulfillment_dto

import "time"

type CreateInboundShipmentPayload struct {
	InboundOrderID string      `json:"inboundOrderId"`
	ReturnAddress  Address     `json:"returnAddress"`
	OrderItems     []OrderItem `json:"orderItems"`
}

type CreateInboundShipmentResponse struct {
	Status  string                                 `json:"status"`
	Payload []CreateInboundShipmentResponsePayload `json:"payload"`
}
type CreateInboundShipmentResponsePayload struct {
	ShipmentID           string         `json:"shipmentId"`
	ShipToAddress        Address        `json:"shipToAddress"`
	ShipmentItems        []ShipmentItem `json:"shipmentItems"`
	ExpectedDeliveryDate *time.Time     `json:"expectedDeliveryDate"`
}

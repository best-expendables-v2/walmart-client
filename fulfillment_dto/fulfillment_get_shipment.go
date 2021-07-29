package fulfillment_dto

import "time"

type GetShipmentParams struct {
	PaginationParams
	InboundOrderID string `json:"inboundOrderId" schema:"inboundOrderId"`
	ShipmentID     string `json:"shipmentId" schema:"shipmentId"`
	Status         string `json:"status" schema:"status"`
	FromCreateDate string `json:"fromCreateDate" schema:"fromCreateDate"`
	ToCreateDate   string `json:"toCreateDate" schema:"toCreateDate"`
}

type GetShipmentResponse struct {
	Headers struct {
		PaginationParams
		TotalCount int `json:"totalCount"`
	} `json:"headers"`
	Payload []GetShipmentResponsePayload `json:"payload"`
}
type GetShipmentResponsePayload struct {
	InboundOrderID              string     `json:"inboundOrderId"`
	ShipmentID                  string     `json:"shipmentId"`
	ShipToAddress               Address    `json:"shipToAddress"`
	ReturnAddress               Address    `json:"returnAddress"`
	Status                      string     `json:"status"`
	CreatedDate                 *time.Time `json:"createdDate"`
	ShipmentUnits               int        `json:"shipmentUnits"`
	ReceivedUnits               int        `json:"receivedUnits"`
	ExpectedDeliveryDate        *time.Time `json:"expectedDeliveryDate"`
	UpdatedExpectedDeliveryDate *time.Time `json:"updatedExpectedDeliveryDate"`
	TrackingNo                  []string   `json:"trackingNo"`
	CarrierName                 *string    `json:"carrierName"`
	MovedShipmentID             *string    `json:"movedShipmentId"`
}

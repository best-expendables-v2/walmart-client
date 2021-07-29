package fulfillment_dto

type GetShipmentLabelParams struct {
	ShipmentID string `json:"shipmentId" schema:"shipmentId"`
}

type GetShipmentLabelResponse struct {
	Headers struct {
		PaginationParams
		TotalCount int `json:"totalCount"`
	} `json:"headers"`
	Payload []GetShipmentResponsePayload `json:"payload"`
}

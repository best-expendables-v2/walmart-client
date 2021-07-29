package fulfillment_dto

type GetInboundShipmentItemsParams struct {
	PaginationParams
	ShipmentID string `json:"shipmentId" schema:"shipmentId"`
}

type GetInboundShipmentItemsResponse struct {
	Headers struct {
		PaginationParams
		TotalCount int `json:"totalCount"`
	} `json:"headers"`
	Payload []ShipmentItemResponse `json:"payload"`
}

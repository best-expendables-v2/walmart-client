package fulfillment_dto

type UpdateShipmentQuantitiesPayload struct {
	InboundOrderID string                                     `json:"inboundOrderId"`
	ShipmentID     string                                     `json:"shipmentId"`
	OrderItems     []UpdateShipmentQuantitiesPayloadOrderItem `json:"orderItems"`
}
type UpdateShipmentQuantitiesPayloadOrderItem struct {
	Sku                string `json:"sku"`
	UpdatedShipmentQty int    `json:"updatedShipmentQty"`
}

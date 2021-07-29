package fulfillment_dto

type UpdateShipmentTrackingPayload struct {
	ShipmentID   string   `json:"shipmentId"`
	CarrierName  string   `json:"carrierName"`
	TrackingInfo []string `json:"trackingInfo"`
}

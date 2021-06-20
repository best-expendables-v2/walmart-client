package inventory_dto

type InventoryResponse struct {
	Sku      string `json:"sku"`
	Quantity struct {
		Unit   string
		Amount int64
	} `json:"quantity"`
}

type InventoryGetParams struct {
	Sku      string  `json:"sku"`
	ShipNode *string `json:"shipNode,omitempty"`
}

type InventoryUpdatePayload struct {
	Sku      string  `json:"sku"`
	Quantity int64   `json:"quantity"`
	ShipNode *string `json:"shipNode,omitempty"`
}

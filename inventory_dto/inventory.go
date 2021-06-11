package inventory_dto

type InventoryResponse struct {
	Sku      string `json:"sku"`
	Quantity struct {
		Unit   string
		Amount int
	} `json:"quantity"`
}

type InventoryGetParams struct {
	Sku      string  `json:"sku"`
	ShipNode *string `json:"shipNode,omitempty"`
}

type InventoryUpdatePayload struct {
	Sku      string  `json:"sku"`
	Quantity int     `json:"quantity"`
	ShipNode *string `json:"shipNode,omitempty"`
}

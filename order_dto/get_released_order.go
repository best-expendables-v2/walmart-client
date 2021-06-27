package order_dto

type GetReleasedOrdersParams struct {
	CreatedStartDate     string `json:"createdStartDate,omitempty"`
	CreatedEndDate       string `json:"createdEndDate,omitempty"`
	Limit                int    `json:"limit,omitempty"`
	ProductInfo          string `json:"productInfo,omitempty"`
	ShipNodeType         string `json:"shipNodeType,omitempty"`
	Sku                  string `json:"sku,omitempty"`
	CustomerOrderID      string `json:"customerOrderId,omitempty"`
	PurchaseOrderID      string `json:"purchaseOrderId,omitempty"`
	FromExpectedShipDate string `json:"fromExpectedShipDate,omitempty"`
	ToExpectedShipDate   string `json:"toExpectedShipDate,omitempty"`
	ShippingProgramType  string `json:"shippingProgramType,omitempty"`
	ReplacementInfo      string `json:"replacementInfo,omitempty"`
	OrderType            string `json:"orderType,omitempty"`
}

type ReleasedOrdersResponse struct {
	List struct {
		Meta struct {
			TotalCount int    `json:"totalCount"`
			Limit      int    `json:"limit"`
			NextCursor string `json:"nextCursor"`
		} `json:"meta"`
		Elements struct {
			Order []Order `json:"order"`
		} `json:"elements"`
	} `json:"list"`
}

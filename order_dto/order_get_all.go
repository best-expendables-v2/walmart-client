package order_dto

type GetAllOrdersParams struct {
	Sku                   string `json:"sku,omitempty"`
	CustomerOrderID       string `json:"customerOrderId,omitempty"`
	PurchaseOrderID       string `json:"purchaseOrderId,omitempty"`
	Status                string `json:"status,omitempty"`
	CreatedStartDate      string `json:"createdStartDate,omitempty"`
	CreatedEndDate        string `json:"createdEndDate,omitempty"`
	FromExpectedShipDate  string `json:"fromExpectedShipDate,omitempty"`
	ToExpectedShipDate    string `json:"toExpectedShipDate,omitempty"`
	LastModifiedStartDate string `json:"lastModifiedStartDate,omitempty"`
	LastModifiedEndDate   string `json:"lastModifiedEndDate,omitempty"`
	Limit                 int    `json:"limit,omitempty"`
	ProductInfo           string `json:"productInfo,omitempty"`
	ShipNodeType          string `json:"shipNodeType,omitempty"`
	ShippingProgramType   string `json:"shippingProgramType,omitempty"`
	ReplacementInfo       string `json:"replacementInfo,omitempty"`
	OrderType             string `json:"orderType,omitempty"`
}

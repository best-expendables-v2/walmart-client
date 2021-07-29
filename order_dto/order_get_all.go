package order_dto

type GetAllOrdersParams struct {
	Sku                   string `json:"sku,omitempty" schema:"sku,omitempty"`
	CustomerOrderID       string `json:"customerOrderId,omitempty" schema:"customerOrderId,omitempty"`
	PurchaseOrderID       string `json:"purchaseOrderId,omitempty" schema:"purchaseOrderId,omitempty"`
	Status                string `json:"status,omitempty" schema:"status,omitempty"`
	CreatedStartDate      string `json:"createdStartDate,omitempty" schema:"createdStartDate,omitempty"`
	CreatedEndDate        string `json:"createdEndDate,omitempty" schema:"createdEndDate,omitempty"`
	FromExpectedShipDate  string `json:"fromExpectedShipDate,omitempty" schema:"fromExpectedShipDate,omitempty"`
	ToExpectedShipDate    string `json:"toExpectedShipDate,omitempty" schema:"toExpectedShipDate,omitempty"`
	LastModifiedStartDate string `json:"lastModifiedStartDate,omitempty" schema:"lastModifiedStartDate,omitempty"`
	LastModifiedEndDate   string `json:"lastModifiedEndDate,omitempty" schema:"lastModifiedEndDate,omitempty"`
	Limit                 int    `json:"limit,omitempty" schema:"limit,omitempty"`
	ProductInfo           string `json:"productInfo,omitempty" schema:"productInfo,omitempty"`
	ShipNodeType          string `json:"shipNodeType,omitempty" schema:"shipNodeType,omitempty"`
	ShippingProgramType   string `json:"shippingProgramType,omitempty" schema:"shippingProgramType,omitempty"`
	ReplacementInfo       string `json:"replacementInfo,omitempty" schema:"replacementInfo,omitempty"`
	OrderType             string `json:"orderType,omitempty" schema:"orderType,omitempty"`
}

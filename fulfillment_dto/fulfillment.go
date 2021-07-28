package fulfillment_dto

import "time"

type OrderItem struct {
	ProductID            string     `json:"productId"`
	ProductType          string     `json:"productType"`
	Sku                  string     `json:"sku"`
	ItemDesc             string     `json:"itemDesc"`
	ItemQty              int        `json:"itemQty"`
	VendorPackQty        int        `json:"vendorPackQty"`
	InnerPackQty         int        `json:"innerPackQty"`
	ExpectedDeliveryDate *time.Time `json:"expectedDeliveryDate"`
}

type Address struct {
	FcName       string `json:"fcName"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City         string `json:"city"`
	StateCode    string `json:"stateCode"`
	CountryCode  string `json:"countryCode"`
	PostalCode   string `json:"postalCode"`
}

type ShipmentItem struct {
	VendorSku string `json:"vendorSku"`
	ItemQty   int    `json:"itemQty"`
}
type ShipmentItemResponse struct {
	InboundOrderID              string     `json:"inboundOrderId"`
	ShipmentID                  string     `json:"shipmentId"`
	Gtin                        string     `json:"gtin"`
	Sku                         string     `json:"sku"`
	ItemDesc                    string     `json:"itemDesc"`
	ItemQty                     int        `json:"itemQty"`
	VendorPackQty               int        `json:"vendorPackQty"`
	InnerPackQty                int        `json:"innerPackQty"`
	ReceivedQty                 int        `json:"receivedQty"`
	DamagedQty                  int        `json:"damagedQty"`
	FillRate                    float64    `json:"fillRate"`
	ExpectedDeliveryDate        *time.Time `json:"expectedDeliveryDate"`
	UpdatedExpectedDeliveryDate *time.Time `json:"updatedExpectedDeliveryDate"`
}

type PaginationParams struct {
	Offset *int `json:"offset,omitempty" schema:"offset,omitempty"`
	Limit  *int `json:"limit,omitempty" schema:"offset,omitempty"`
}

type BasicResponse struct {
	Status  string      `json:"status"`
	Headers interface{} `json:"headers"`
	Errors  []Error     `json:"errors"`
	Payload interface{} `json:"payload"`
}

type Error struct {
	Code        string `json:"code"`
	Field       string `json:"field"`
	Description string `json:"description"`
	Info        string `json:"info"`
	Severity    string `json:"severity"`
	Category    string `json:"category"`
}

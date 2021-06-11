package order_dto

type OrderShipmentPayload struct {
	OrderShipment OrderShipment `json:"orderShipment"`
}

type OrderShipment struct {
	ProcessMode *string                 `json:"processMode,omitempty"`
	OrderLines  OrderShipmentOrderLines `json:"orderLines"`
}

type OrderShipmentOrderLines struct {
	OrderLine []OrderShipmentOrderLine `json:"orderLine"`
}

type OrderShipmentOrderLine struct {
	LineNumber        string                         `json:"lineNumber"`
	SellerOrderID     string                         `json:"sellerOrderId"`
	OrderLineStatuses OrderShipmentOrderLineStatuses `json:"orderLineStatuses"`
	SellerOrderNo     *string                        `json:"sellerOrderNo,omitempty"`
}

type OrderShipmentOrderLineStatuses struct {
	OrderLineStatus []OrderShipmentOrderLineStatus `json:"orderLineStatus"`
}

type OrderShipmentOrderLineStatus struct {
	Status              string                            `json:"status"`
	Asn                 *OrderShipmentAsn                 `json:"asn"`
	StatusQuantity      OrderShipmentStatusQuantity       `json:"statusQuantity"`
	TrackingInfo        OrderShipmentTrackingInfo         `json:"trackingInfo"`
	ReturnCenterAddress *OrderShipmentReturnCenterAddress `json:"returnCenterAddress,omitempty"`
}

type OrderShipmentAsn struct {
	PackageASN string  `json:"packageASN"`
	PalletASN  *string `json:"palletASN,omitempty"`
}

type OrderShipmentStatusQuantity struct {
	UnitOfMeasurement string `json:"unitOfMeasurement"`
	Amount            string `json:"amount"`
}

type OrderShipmentTrackingInfo struct {
	ShipDateTime   int64                    `json:"shipDateTime"`
	CarrierName    OrderShipmentCarrierName `json:"carrierName"`
	MethodCode     string                   `json:"methodCode"`
	TrackingNumber string                   `json:"trackingNumber"`
	TrackingURL    *string                  `json:"trackingURL,omitempty"`
}

type OrderShipmentReturnCenterAddress struct {
	Name       *string `json:"name,omitempty"`
	Address1   string  `json:"address1"`
	Address2   *string `json:"address2,omitempty"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	PostalCode string  `json:"postalCode"`
	Country    string  `json:"country"`
	DayPhone   *string `json:"dayPhone,omitempty"`
	EmailID    *string `json:"emailId,omitempty"`
}

type OrderShipmentCarrierName struct {
	OtherCarrier *string `json:"otherCarrier,omitempty"`
	Carrier      *string `json:"carrier,omitempty"`
}

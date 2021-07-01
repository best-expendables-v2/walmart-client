package order_dto

type OrderResponse struct {
	Order Order `json:"order"`
}

type OrdersResponse struct {
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

type Order struct {
	PurchaseOrderID string       `json:"purchaseOrderId"`
	CustomerOrderID string       `json:"customerOrderId"`
	SellerOrderID   string       `json:"sellerOrderId"`
	CustomerEmailID string       `json:"customerEmailId"`
	OrderDate       int64        `json:"orderDate"`
	ShippingInfo    ShippingInfo `json:"shippingInfo"`
	OrderLines      OrderLines   `json:"orderLines"`
	ShipNode        *ShipNode    `json:"shipNode"`
}

type OrderLines struct {
	OrderLine []OrderLine `json:"orderLine"`
}

type OrderLine struct {
	LineNumber string `json:"lineNumber"`
	Item       Item   `json:"item"`
	Charges    struct {
		Charge []struct {
			ChargeType   string `json:"chargeType"`
			ChargeName   string `json:"chargeName"`
			ChargeAmount struct {
				Currency string `json:"currency"`
				Amount   int    `json:"amount"`
			} `json:"chargeAmount"`
			Tax struct {
				TaxName   string `json:"taxName"`
				TaxAmount struct {
					Currency string `json:"currency"`
					Amount   int    `json:"amount"`
				} `json:"taxAmount"`
			} `json:"tax"`
		} `json:"charge"`
	} `json:"charges"`
	OrderLineQuantity OrderLineQuantity `json:"orderLineQuantity"`
	SellerOrderID     *string           `json:"sellerOrderId,omitempty"`
	OrderLineStatuses struct {
		OrderLineStatus []struct {
			Status string `json:"status"`
			Asn    *struct {
				PackageASN string  `json:"packageASN"`
				PalletASN  *string `json:"palletASN,omitempty"`
			} `json:"asn"`
			StatusQuantity struct {
				UnitOfMeasurement string `json:"unitOfMeasurement"`
				Amount            string `json:"amount"`
			} `json:"statusQuantity"`
			TrackingInfo struct {
				ShipDateTime int64 `json:"shipDateTime"`
				CarrierName  struct {
					OtherCarrier *string `json:"otherCarrier,omitempty"`
					Carrier      *string `json:"carrier,omitempty"`
				} `json:"carrierName"`
				MethodCode     string  `json:"methodCode"`
				TrackingNumber string  `json:"trackingNumber"`
				TrackingURL    *string `json:"trackingURL,omitempty"`
			} `json:"trackingInfo"`
			ReturnCenterAddress *struct {
				Name       *string `json:"name,omitempty"`
				Address1   string  `json:"address1"`
				Address2   *string `json:"address2,omitempty"`
				City       string  `json:"city"`
				State      string  `json:"state"`
				PostalCode string  `json:"postalCode"`
				Country    string  `json:"country"`
				DayPhone   *string `json:"dayPhone,omitempty"`
				EmailID    *string `json:"emailId,omitempty"`
			} `json:"returnCenterAddress,omitempty"`
		} `json:"orderLineStatus"`
	} `json:"orderLineStatuses"`
	SellerOrderNo *string `json:"sellerOrderNo,omitempty"`
}

type ShippingInfo struct {
	Phone                 string `json:"phone"`
	EstimatedDeliveryDate int64  `json:"estimatedDeliveryDate"`
	EstimatedShipDate     int64  `json:"estimatedShipDate"`
	MethodCode            string `json:"methodCode"`
	PostalAddress         struct {
		Name        string `json:"name"`
		Address1    string `json:"address1"`
		Address2    string `json:"address2"`
		City        string `json:"city"`
		State       string `json:"state"`
		PostalCode  string `json:"postalCode"`
		Country     string `json:"country"`
		AddressType string `json:"addressType"`
	} `json:"postalAddress"`
}

type Item struct {
	ProductName string `json:"productName"`
	Sku         string `json:"sku"`
	ImageURL    string `json:"imageUrl"`
	Weight      struct {
		Value string `json:"value"`
		Unit  string `json:"unit"`
	} `json:"weight"`
}

type OrderLineQuantity struct {
	UnitOfMeasurement string `json:"unitOfMeasurement"`
	Amount            string `json:"amount"`
}

type ShipNode struct {
	Type string `json:"type"`
	Name string `json:"name"`
	ID   string `json:"id"`
}

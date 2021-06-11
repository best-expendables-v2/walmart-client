package order_dto

type OrderCancellationPayload struct {
	OrderCancellation OrderCancellation `json:"orderCancellation"`
}

type OrderCancellation struct {
	OrderLines OrderCancellationOrderLines `json:"orderLines"`
}

type OrderCancellationOrderLines struct {
	OrderLine []OrderCancellationOrderLine `json:"orderLine"`
}

type OrderCancellationOrderLine struct {
	LineNumber        string                             `json:"lineNumber"`
	OrderLineStatuses OrderCancellationOrderLineStatuses `json:"orderLineStatuses"`
}

type OrderCancellationOrderLineStatuses struct {
	OrderLineStatus []OrderCancellationOrderLineStatus `json:"orderLineStatus"`
}

type OrderCancellationOrderLineStatus struct {
	Status             string                          `json:"status"`
	CancellationReason string                          `json:"cancellationReason"`
	StatusQuantity     OrderCancellationStatusQuantity `json:"statusQuantity"`
}

type OrderCancellationStatusQuantity struct {
	UnitOfMeasurement string `json:"unitOfMeasurement"`
	Amount            string `json:"amount"`
}

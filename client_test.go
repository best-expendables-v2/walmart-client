package walmart_client_test

import (
	"context"
	"net/http"
	"testing"

	walmart_client "github.com/best-expendables-v2/walmart-client"
	"github.com/best-expendables-v2/walmart-client/order_dto"
	"github.com/stretchr/testify/assert"
)

var (
	ctx  = context.Background()
	conf = walmart_client.Config{
		SandboxMode:   true,
		WMServiceName: "Walmart Service",
		ClientID:      "2ceb0300-a439-4f64-877b-a3d003dd5f7b",
		ClientSecret:  "FX5cbPeqE_BmSRMMGYl9ArnrgYQFm_VwtnVUIa77lhha-x2R6ZkRAEB3RiOzr7TVaKCkF9AuJgSOT4i2ztvakA",
	}
	client = walmart_client.NewClient(http.DefaultClient, conf)
)

func TestClient_Authenticate(t *testing.T) {
	_, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
}

func TestClient_AcknowlegeOrder(t *testing.T) {
	res, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
	client.SetAccessToken(res.AccessToken, nil)

	resp, err := client.Acknowledge(ctx, "1796277083022")
	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, "1796277083022", resp.Order.PurchaseOrderID)
}

func TestClient_OrderShipment(t *testing.T) {
	res, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
	client.SetAccessToken(res.AccessToken, nil)

	payload := order_dto.OrderShipmentPayload{
		OrderShipment: order_dto.OrderShipment{
			ProcessMode: nil,
			OrderLines: order_dto.OrderShipmentOrderLines{
				OrderLine: []order_dto.OrderShipmentOrderLine{
					{
						LineNumber:    "1",
						SellerOrderID: "92344",
						OrderLineStatuses: order_dto.OrderShipmentOrderLineStatuses{
							OrderLineStatus: []order_dto.OrderShipmentOrderLineStatus{
								{
									Status: "Shipped",
									Asn:    nil,
									StatusQuantity: order_dto.OrderShipmentStatusQuantity{
										UnitOfMeasurement: "EACH",
										Amount:            "1",
									},
									TrackingInfo: order_dto.OrderShipmentTrackingInfo{
										ShipDateTime: 1580821866000,
										CarrierName: order_dto.OrderShipmentCarrierName{
											OtherCarrier: nil,
											Carrier:      toStringPointer("UPS"),
										},
										MethodCode:     "Standard",
										TrackingNumber: "22344",
										TrackingURL:    toStringPointer("http://walmart/tracking/ups?&type=MP&seller_id=12345&promise_date=03/02/2020&dzip=92840&tracking_numbers=92345"),
									},
									ReturnCenterAddress: &order_dto.OrderShipmentReturnCenterAddress{
										Name:       toStringPointer("walmart"),
										Address1:   "walmart store 2",
										Address2:   nil,
										City:       "Huntsville",
										State:      "AL",
										PostalCode: "35805",
										Country:    "USA",
										DayPhone:   toStringPointer("12344"),
										EmailID:    toStringPointer("walmart@walmart.com"),
									},
								},
							},
						},
						SellerOrderNo: nil,
					},
					{
						LineNumber:    "1",
						SellerOrderID: "92344",
						OrderLineStatuses: order_dto.OrderShipmentOrderLineStatuses{
							OrderLineStatus: []order_dto.OrderShipmentOrderLineStatus{
								{
									Status: "Shipped",
									Asn:    nil,
									StatusQuantity: order_dto.OrderShipmentStatusQuantity{
										UnitOfMeasurement: "EACH",
										Amount:            "1",
									},
									TrackingInfo: order_dto.OrderShipmentTrackingInfo{
										ShipDateTime: 1580821866000,
										CarrierName: order_dto.OrderShipmentCarrierName{
											OtherCarrier: nil,
											Carrier:      toStringPointer("UPS"),
										},
										MethodCode:     "Express",
										TrackingNumber: "22344",
										TrackingURL:    toStringPointer("http://walmart/tracking/fedEx?&type=MP&seller_id=12345&promise_date=03/02/2020&dzip=92840&tracking_numbers=92344"),
									},
									ReturnCenterAddress: &order_dto.OrderShipmentReturnCenterAddress{
										Name:       toStringPointer("walmart"),
										Address1:   "walmart store 2",
										Address2:   nil,
										City:       "Huntsville",
										State:      "AL",
										PostalCode: "35805",
										Country:    "USA",
										DayPhone:   toStringPointer("12344"),
										EmailID:    toStringPointer("walmart@walmart.com"),
									},
								},
							},
						},
						SellerOrderNo: nil,
					},
				},
			},
		},
	}

	resp, err := client.Ship(ctx, "1234567891234", payload)
	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, "1234567891234", resp.Order.PurchaseOrderID)
}

func TestClient_OrderCancellation(t *testing.T) {
	res, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
	client.SetAccessToken(res.AccessToken, nil)

	payload := order_dto.OrderCancellationPayload{
		OrderCancellation: order_dto.OrderCancellation{
			OrderLines: order_dto.OrderCancellationOrderLines{
				OrderLine: []order_dto.OrderCancellationOrderLine{
					{
						LineNumber: "1",
						OrderLineStatuses: order_dto.OrderCancellationOrderLineStatuses{
							OrderLineStatus: []order_dto.OrderCancellationOrderLineStatus{
								{
									Status:             "Cancelled",
									CancellationReason: "CUSTOMER_REQUESTED_SELLER_TO_CANCEL",
									StatusQuantity: order_dto.OrderCancellationStatusQuantity{
										UnitOfMeasurement: "EA",
										Amount:            "1",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	resp, err := client.Cancel(ctx, "1577914061094", payload)
	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, "1577914061094", resp.Order.PurchaseOrderID)
}

func toStringPointer(v string) *string {
	return &v
}

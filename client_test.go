package walmart_client_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/best-expendables-v2/walmart-client/feed_dto"
	"github.com/best-expendables-v2/walmart-client/fulfillment_dto"

	walmart_client "github.com/best-expendables-v2/walmart-client"
	"github.com/best-expendables-v2/walmart-client/order_dto"
	"github.com/stretchr/testify/assert"
)

var (
	ctx  = context.Background()
	now  = time.Now()
	conf = walmart_client.Config{
		SandboxMode:   true,
		WMServiceName: "Walmart Service",
		ClientID:      "2ceb0300-a439-4f64-877b-a3d003dd5f7b",
		ClientSecret:  "FX5cbPeqE_BmSRMMGYl9ArnrgYQFm_VwtnVUIa77lhha-x2R6ZkRAEB3RiOzr7TVaKCkF9AuJgSOT4i2ztvakA",
	}
	client = walmart_client.NewClient(http.DefaultClient, conf)
)

func toStringPointer(v string) *string {
	return &v
}

func TestClient_New(t *testing.T) {
	checkAuth := func(ctx context.Context, c *walmart_client.Client) error {
		buffer := time.Minute
		if c.Token == nil || time.Now().Add(buffer).After(c.Token.ExpireAt) {
			token, err := c.Authenticate(ctx)
			if err != nil {
				return err
			}
			c.SetAccessToken(token.AccessToken, nil)
			tokenDetail, err := c.GetTokenDetail(ctx)
			if err != nil {
				return err
			}
			c.SetAccessToken(token.AccessToken, &tokenDetail.ExpireAt)
		}
		return nil
	}
	newClient := walmart_client.NewClient(http.DefaultClient, conf, walmart_client.BeforeRequest(checkAuth))
	resp, err := newClient.Acknowledge(ctx, "1796277083022")
	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, "1796277083022", resp.Order.PurchaseOrderID)
}

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

func TestClient_FeedConvertItemsForWFS(t *testing.T) {
	res, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
	client.SetAccessToken(res.AccessToken, nil)

	payload := feed_dto.ConvertItemsPayload{
		SupplierItemFeedHeader: feed_dto.SupplierItemFeedHeader{
			SubCategory:    "electronics_accessories",
			SellingChannel: "fbw",
			ProcessMode:    "REPLACE",
			Locale:         "en",
			Version:        "1.3",
			Subset:         "EXTERNAL",
		},
		SupplierItem: []feed_dto.SupplierItem{
			{
				Visible: feed_dto.Visible{
					"Electronics Accessories": feed_dto.CategoryAttributes{
						ManufacturerPartNumber: "234566",
						MainImageUrl:           "http://example.com/lf/71192763.jpg",
						Prop65WarningText:      "some text",
						Manufacturer:           "walmart",
					},
				},
				Orderable: feed_dto.Orderable{
					HasBatteries:          "Yes",
					BatteryTechnologyType: "Lithium Ion",
					LithiumIonBatteries: feed_dto.LithiumIonBatteries{
						BatteryFormFactor:        "Battery",
						NumberOfBatteries:        2,
						BatteryModel:             "L800",
						BatteryWeight:            3,
						BatteryWattHour:          10,
						NumberOfBatteryCells:     3,
						IncludedBatteryPackaging: "Contained In",
					},
					ProductName: "UAT Test 1",
					ProductIdentifiers: feed_dto.ProductIdentifiers{
						ProductID:     "00055546562284",
						ProductIDType: "GTIN",
					},
					ElectronicsIndicator: "Yes",
					SafetyDataSheet: []string{
						"https://webviewer.wercsmart.com/walmart_logistics/client/GetDocument?id=d100128c-5153-479d-947f-b84a19fbbf45",
					},
					NumberOfHazardousComponents: 1,
					Price:                       50,
					BatterySize:                 "AA",
					ChemicalAerosolPesticide:    "NO",
					Sku:                         "wfs_wercs_30",
					StateRestrictions: []feed_dto.StateRestrictions{
						{
							StateRestrictionsText: "Commercial",
							States:                "CA,AR",
						},
						{
							StateRestrictionsText: "Illegal for Sale",
							States:                "TX,PA",
						},
					},
					Brand: "test",
				},
				TradeItem: feed_dto.TradeItem{
					CountryOfOriginAssembly: []string{"US - United States"},
					Sku:                     "wfs_wercs_30",
					Each: feed_dto.Each{
						EachDepth:  2,
						EachWeight: 5,
						EachWidth:  4,
						EachHeight: 3,
						EachGTIN:   "00055546562284",
					},
					OrderableGTIN: "00055546562284",
				},
			},
		},
	}
	resp, err := client.ConvertItemsForWFS(ctx, payload)
	assert.NoError(t, err, "Unexpected error")
	fmt.Printf("Response %v\n", resp)
}

func TestClient_CreateInboundShipment(t *testing.T) {
	res, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
	client.SetAccessToken(res.AccessToken, nil)

	payload := fulfillment_dto.CreateInboundShipmentPayload{
		InboundOrderID: "8778881015027123",
		ReturnAddress: fulfillment_dto.Address{
			AddressLine1: "860 W California Ave",
			AddressLine2: "",
			City:         "Sunnyvale",
			StateCode:    "CA",
			CountryCode:  "USA",
			PostalCode:   "94086",
		},
		OrderItems: []fulfillment_dto.OrderItem{
			{
				ProductID:            "00894147009695111",
				ProductType:          "GTIN",
				Sku:                  "WILL-SL96911",
				ItemDesc:             "Blue jeans",
				ItemQty:              10,
				VendorPackQty:        10,
				InnerPackQty:         1,
				ExpectedDeliveryDate: &now,
			},
		},
	}
	resp, err := client.CreateInboundShipment(ctx, payload)
	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, resp.Status, "OK")
	fmt.Printf("Response %v\n", resp)
}

func TestClient_GetInboundShipmentErrors(t *testing.T) {
	res, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
	client.SetAccessToken(res.AccessToken, nil)

	payload := fulfillment_dto.GetInboundShipmentErrorsParams{
		ShipmentID: "0000966GDM",
	}
	resp, err := client.GetInboundShipmentErrors(ctx, payload)
	assert.NoError(t, err, "Unexpected error")
	fmt.Printf("Response %v\n", resp)
}

func TestClient_GetShipments(t *testing.T) {
	res, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
	client.SetAccessToken(res.AccessToken, nil)

	payload := fulfillment_dto.GetShipmentParams{
		InboundOrderID: "",
		ShipmentID:     "",
		Status:         "",
		FromCreateDate: "",
		ToCreateDate:   "",
	}
	resp, err := client.GetShipments(ctx, payload)
	assert.NoError(t, err, "Unexpected error")
	fmt.Printf("Response %v\n", resp)
}

func TestClient_GetInboundShipmentItems(t *testing.T) {
	res, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
	client.SetAccessToken(res.AccessToken, nil)

	payload := fulfillment_dto.GetInboundShipmentItemsParams{
		ShipmentID: "0000966GDM",
	}
	resp, err := client.GetInboundShipmentItems(ctx, payload)
	assert.NoError(t, err, "Unexpected error")
	fmt.Printf("Response %v\n", resp)
}

func TestClient_UpdateShipmentQuantities(t *testing.T) {
	res, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
	client.SetAccessToken(res.AccessToken, nil)

	payload := fulfillment_dto.UpdateShipmentQuantitiesPayload{
		InboundOrderID: "1235113",
		ShipmentID:     "4846GDM",
		OrderItems: []fulfillment_dto.UpdateShipmentQuantitiesPayloadOrderItem{
			{
				Sku:                "ACDD-WZK73685",
				UpdatedShipmentQty: 11,
			},
		},
	}
	resp, err := client.UpdateShipmentQuantities(ctx, payload)
	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, resp.Status, "OK")
	fmt.Printf("Response %v\n", resp)
}

func TestClient_CreateInboundShipmentLabel(t *testing.T) {
	res, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
	client.SetAccessToken(res.AccessToken, nil)

	err = client.CreateInboundShipmentLabel(ctx, "4846GDM")
	assert.NoError(t, err, "Unexpected error")
}

func TestClient_UpdateShipmentTracking(t *testing.T) {
	res, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
	client.SetAccessToken(res.AccessToken, nil)
	payload := fulfillment_dto.UpdateShipmentTrackingPayload{
		ShipmentID:   "4846GDM",
		CarrierName:  "UPS",
		TrackingInfo: []string{"123", "456-001"},
	}

	resp, err := client.UpdateShipmentTracking(ctx, payload)
	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, resp.Status, "OK")
	fmt.Printf("Response %v\n", resp)
}

func TestClient_CancelInboundShipment(t *testing.T) {
	res, err := client.Authenticate(ctx)
	assert.NoError(t, err, "Unexpected error")
	client.SetAccessToken(res.AccessToken, nil)

	resp, err := client.CancelInboundShipment(ctx, "1235113")
	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, resp.Status, "OK")
	fmt.Printf("Response %v\n", resp)
}

package feed_dto

type ConvertItemsPayload struct {
	SupplierItemFeedHeader `json:"SupplierItemFeedHeader"`
	SupplierItem           []SupplierItem `json:"SupplierItem"`
}
type SupplierItemFeedHeader struct {
	SubCategory    string `json:"subCategory"`
	SellingChannel string `json:"sellingChannel"`
	ProcessMode    string `json:"processMode"`
	Locale         string `json:"locale"`
	Version        string `json:"version"`
	Subset         string `json:"subset"`
}

type SupplierItem struct {
	Visible   `json:"Visible"`
	Orderable `json:"Orderable"`
	TradeItem `json:"TradeItem"`
}
type Visible struct {
	ElectronicsAccessories `json:"Electronics Accessories"`
}
type ElectronicsAccessories struct {
	ManufacturerPartNumber string `json:"manufacturerPartNumber"`
	MainImageUrl           string `json:"mainImageUrl"`
	Prop65WarningText      string `json:"prop65WarningText"`
	Manufacturer           string `json:"manufacturer"`
}
type Orderable struct {
	HasBatteries                string `json:"hasBatteries"`
	BatteryTechnologyType       string `json:"batteryTechnologyType"`
	LithiumIonBatteries         `json:"lithiumIonBatteries"`
	ProductName                 string `json:"productName"`
	ProductIdentifiers          `json:"productIdentifiers"`
	ElectronicsIndicator        string              `json:"electronicsIndicator"`
	SafetyDataSheet             []string            `json:"safetyDataSheet"`
	NumberOfHazardousComponents int                 `json:"numberOfHazardousComponents"`
	Price                       float64             `json:"price"`
	BatterySize                 string              `json:"batterySize"`
	ChemicalAerosolPesticide    string              `json:"chemicalAerosolPesticide"`
	Sku                         string              `json:"sku"`
	StateRestrictions           []StateRestrictions `json:"stateRestrictions"`
	Brand                       string              `json:"brand"`
}
type LithiumIonBatteries struct {
	BatteryFormFactor        string `json:"batteryFormFactor"`
	NumberOfBatteries        int    `json:"numberOfBatteries"`
	BatteryModel             string `json:"batteryModel"`
	BatteryWeight            int    `json:"batteryWeight"`
	BatteryWattHour          int    `json:"batteryWattHour"`
	NumberOfBatteryCells     int    `json:"numberOfBatteryCells"`
	IncludedBatteryPackaging string `json:"includedBatteryPackaging"`
}

type ProductIdentifiers struct {
	ProductID     string `json:"productId"`
	ProductIDType string `json:"productIdType"`
}

type StateRestrictions struct {
	StateRestrictionsText string `json:"stateRestrictionsText"`
	States                string `json:"states"`
}

type TradeItem struct {
	CountryOfOriginAssembly []string `json:"countryOfOriginAssembly"`
	Sku                     string   `json:"sku"`
	Each                    `json:"each"`
	OrderableGTIN           string `json:"orderableGTIN"`
}
type Each struct {
	EachDepth  int    `json:"eachDepth"`
	EachWeight int    `json:"eachWeight"`
	EachWidth  int    `json:"eachWidth"`
	EachHeight int    `json:"eachHeight"`
	EachGTIN   string `json:"eachGTIN"`
}

type ConvertItemsResponse struct {
	FeedID string `json:"feedId"`
}

package feed_dto

const (
	CategoryCasesAndBags               = "cases_and_bags"
	CategoryBuildingSupply             = "building_supply"
	CategoryComputerComponents         = "computer_components"
	CategoryHealthAndBeautyElectronics = "health_and_beauty_electronics"
	CategoryFurnitureOther             = "furniture_other"
	CategoryElectronicsAccessories     = "electronics_accessories"
)

type CategoryAttributes struct {
	ManufacturerPartNumber   string `json:"manufacturerPartNumber,omitempty"`
	MainImageUrl             string `json:"mainImageUrl,omitempty"`
	Prop65WarningText        string `json:"prop65WarningText,omitempty"`
	Manufacturer             string `json:"manufacturer,omitempty"`
	ProductSecondaryImageURL string `json:"productSecondaryImageURL,omitempty"`
}

package data

// OptionData holds data about an instrument that is trade-able as an Option.
type OptionData struct {
}

// AssetClass specifies the properties and class of the asset.
type AssetClass struct {
}

// PriceData is a response type when
type PriceData struct {
}

type Option struct {
	StrikePrice float64
	Premium     float64
}

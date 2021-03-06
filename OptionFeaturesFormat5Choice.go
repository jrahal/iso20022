package iso20022

// Choice between a standard code or a proprietary code to specify the features that may apply to a corporate action option.
type OptionFeaturesFormat5Choice struct {

	// Standard code to specify the features that may apply to a corporate action option.
	Code *OptionFeatures3Code `xml:"Cd"`

	// Proprietary identification of the features that may apply to a corporate action option.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (o *OptionFeaturesFormat5Choice) SetCode(value string) {
	o.Code = (*OptionFeatures3Code)(&value)
}

func (o *OptionFeaturesFormat5Choice) AddProprietary() *GenericIdentification20 {
	o.Proprietary = new(GenericIdentification20)
	return o.Proprietary
}

package iso20022

// Unique and unambiguous way to identify an organisation.
type PartyIdentification22 struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICOrBEI *AnyBICIdentifier `xml:"BICOrBEI"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	AlternativeIdentifier []*Max35Text `xml:"AltrntvIdr,omitempty"`

}


func (p *PartyIdentification22) SetBICOrBEI(value string) {
	p.BICOrBEI = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification22) AddAlternativeIdentifier(value string) {
	p.AlternativeIdentifier = append(p.AlternativeIdentifier, (*Max35Text)(&value))
}

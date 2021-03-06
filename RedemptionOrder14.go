package iso20022

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionOrder14 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Category of the investment fund order.
	OrderType []*FundOrderType4Choice `xml:"OrdrTp,omitempty"`

	// Investment fund class related to the order.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls"`

	// Subdivision of the account used to segregate specific holdings.
	SubAccountForHolding *SubAccount6 `xml:"SubAcctForHldg,omitempty"`

	// Amount of money or the number of units or percentage to be redeemed for the redemption order.
	AmountOrUnitsOrPercentage *FinancialInstrumentQuantity28Choice `xml:"AmtOrUnitsOrPctg"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Information needed to process a currency exchange or conversion.
	// How the exchange rate is expressed determines which currency is the Unit Currency and Quoted Currency. If the amounts concerned are EUR 1000 and USD 1300, the exchange rate may be expressed as per either of the following examples:
	// EXAMPLE 1
	// UnitCurrency  EUR
	// QuotedCurrency  USD
	// ExchangeRate  1.300
	// EXAMPLE 2
	// UnitCurrency  USD
	// QuotedCurrency  EUR
	// ExchangeRate  0.769
	ForeignExchangeDetails *ForeignExchangeTerms32 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnit1Code `xml:"Grp1Or2Units,omitempty"`

	// Fees (charges/commission) and tax to be applied to the gross amount.
	TransactionOverhead *FeeAndTax1 `xml:"TxOvrhd,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters12 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction72 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown2 `xml:"StffClntBrkdwn,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Party related to the transaction.
	RelatedPartyDetails []*Intermediary40 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's retained subscription amount that is returned by the fund in order to reimburse preliminary incentive/performance fees.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`

	// Assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Type of signature.
	SignatureType *SignatureType1Choice `xml:"SgntrTp,omitempty"`

	// Information about a non-standard order.
	OrderWaiverDetails *OrderWaiver1 `xml:"OrdrWvrDtls,omitempty"`
}

func (r *RedemptionOrder14) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder14) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder14) AddOrderType() *FundOrderType4Choice {
	newValue := new(FundOrderType4Choice)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionOrder14) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	r.FinancialInstrumentDetails = new(FinancialInstrument57)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionOrder14) AddSubAccountForHolding() *SubAccount6 {
	r.SubAccountForHolding = new(SubAccount6)
	return r.SubAccountForHolding
}

func (r *RedemptionOrder14) AddAmountOrUnitsOrPercentage() *FinancialInstrumentQuantity28Choice {
	r.AmountOrUnitsOrPercentage = new(FinancialInstrumentQuantity28Choice)
	return r.AmountOrUnitsOrPercentage
}

func (r *RedemptionOrder14) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionOrder14) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder14) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionOrder14) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionOrder14) AddForeignExchangeDetails() *ForeignExchangeTerms32 {
	r.ForeignExchangeDetails = new(ForeignExchangeTerms32)
	return r.ForeignExchangeDetails
}

func (r *RedemptionOrder14) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionOrder14) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnit1Code)(&value)
}

func (r *RedemptionOrder14) AddTransactionOverhead() *FeeAndTax1 {
	r.TransactionOverhead = new(FeeAndTax1)
	return r.TransactionOverhead
}

func (r *RedemptionOrder14) AddSettlementAndCustodyDetails() *FundSettlementParameters12 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters12)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionOrder14) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionOrder14) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionOrder14) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *RedemptionOrder14) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (r *RedemptionOrder14) AddCashSettlementDetails() *PaymentTransaction72 {
	r.CashSettlementDetails = new(PaymentTransaction72)
	return r.CashSettlementDetails
}

func (r *RedemptionOrder14) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionOrder14) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown2 {
	newValue := new(InvestmentFundsOrderBreakdown2)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionOrder14) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionOrder14) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionOrder14) AddRelatedPartyDetails() *Intermediary40 {
	newValue := new(Intermediary40)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionOrder14) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}

func (r *RedemptionOrder14) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	r.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return r.CustomerConductClassification
}

func (r *RedemptionOrder14) AddTransactionChannelType() *TransactionChannelType1Choice {
	r.TransactionChannelType = new(TransactionChannelType1Choice)
	return r.TransactionChannelType
}

func (r *RedemptionOrder14) AddSignatureType() *SignatureType1Choice {
	r.SignatureType = new(SignatureType1Choice)
	return r.SignatureType
}

func (r *RedemptionOrder14) AddOrderWaiverDetails() *OrderWaiver1 {
	r.OrderWaiverDetails = new(OrderWaiver1)
	return r.OrderWaiverDetails
}

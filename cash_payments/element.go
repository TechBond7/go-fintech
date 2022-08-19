// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package cash_payments

import (
	"encoding/xml"

	"github.com/moov-io/fincen"
)

type EFilingBatchXML struct {
	XMLName       xml.Name `xml:"EFilingBatchXML"`
	FormTypeCode  string   `xml:"FormTypeCode"`
	Activity      []string `xml:"Activity"`
	TotalAmount   float64  `xml:"TotalAmount,attr"`
	PartyCount    int64    `xml:"PartyCount,attr"`
	ActivityCount int64    `xml:"ActivityCount,attr"`
}

func (r EFilingBatchXML) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Activity struct {
	EFilingPriorDocumentNumber                  int64                             `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	FilingDateText                              fincen.DateYYYYMMDDType           `xml:"FilingDateText"`
	MultipleSubjectsIndicator                   *fincen.ValidateIndicatorType     `xml:"MultipleSubjectsIndicator,omitempty" json:",omitempty"`
	SuspiciousTransactionIndicator              *fincen.ValidateIndicatorType     `xml:"SuspiciousTransactionIndicator,omitempty" json:",omitempty"`
	TransactionOnBehalfMultiplePersonsIndicator *fincen.ValidateIndicatorType     `xml:"TransactionOnBehalfMultiplePersonsIndicator,omitempty" json:",omitempty"`
	ActivityAssociation                         *ActivityAssociationType          `xml:"ActivityAssociation"`
	Party                                       []string                          `xml:"Party"`
	CurrencyTransactionActivity                 string                            `xml:"CurrencyTransactionActivity"`
	ActivityNarrativeInformation                *ActivityNarrativeInformationType `xml:"ActivityNarrativeInformation,omitempty" json:",omitempty"`
	SeqNum                                      int64                             `xml:"SeqNum,attr"`
}

func (r Activity) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityAssociationType struct {
	CorrectsAmendsPriorReportIndicator *fincen.ValidateIndicatorType `xml:"CorrectsAmendsPriorReportIndicator,omitempty" json:",omitempty"`
	InitialReportIndicator             *fincen.ValidateIndicatorType `xml:"InitialReportIndicator,omitempty" json:",omitempty"`
	SeqNum                             int64                         `xml:"SeqNum,attr"`
}

func (r ActivityAssociationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityNarrativeInformationType struct {
	ActivityNarrativeSequenceNumber ValidateActivityNarrativeSequenceNumber `xml:"ActivityNarrativeSequenceNumber"`
	ActivityNarrativeText           fincen.RestrictString750                `xml:"ActivityNarrativeText"`
	SeqNum                          int64                                   `xml:"SeqNum,attr"`
}

func (r ActivityNarrativeInformationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type ActivityType struct {
	EFilingPriorDocumentNumber                  int64                         `xml:"EFilingPriorDocumentNumber,omitempty" json:",omitempty"`
	FilingDateText                              fincen.DateYYYYMMDDType       `xml:"FilingDateText"`
	MultipleSubjectsIndicator                   *fincen.ValidateIndicatorType `xml:"MultipleSubjectsIndicator,omitempty" json:",omitempty"`
	SuspiciousTransactionIndicator              *fincen.ValidateIndicatorType `xml:"SuspiciousTransactionIndicator,omitempty" json:",omitempty"`
	TransactionOnBehalfMultiplePersonsIndicator *fincen.ValidateIndicatorType `xml:"TransactionOnBehalfMultiplePersonsIndicator,omitempty" json:",omitempty"`
	SeqNum                                      int64                         `xml:"SeqNum,attr"`
}

func (r ActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type AddressType struct {
	RawCityText           *fincen.RestrictString50  `xml:"RawCityText,omitempty" json:",omitempty"`
	RawCountryCodeText    *fincen.RestrictString2   `xml:"RawCountryCodeText,omitempty" json:",omitempty"`
	RawStateCodeText      *fincen.RestrictString3   `xml:"RawStateCodeText,omitempty" json:",omitempty"`
	RawStreetAddress1Text *fincen.RestrictString100 `xml:"RawStreetAddress1Text,omitempty" json:",omitempty"`
	RawZIPCode            *fincen.RestrictString9   `xml:"RawZIPCode,omitempty" json:",omitempty"`
	SeqNum                int64                     `xml:"SeqNum,attr"`
}

func (r AddressType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CurrencyTransactionActivity struct {
	InstallmentPaymentOtherIndicator  *fincen.ValidateIndicatorType           `xml:"InstallmentPaymentOtherIndicator,omitempty" json:",omitempty"`
	Total100DollarBillInAmountText    string                                  `xml:"Total100DollarBillInAmountText,omitempty" json:",omitempty"`
	TotalCashInReceiveAmountText      string                                  `xml:"TotalCashInReceiveAmountText"`
	TotalPriceAmountText              string                                  `xml:"TotalPriceAmountText,omitempty" json:",omitempty"`
	TransactionDateText               fincen.DateYYYYMMDDType                 `xml:"TransactionDateText"`
	CurrencyTransactionActivityDetail []CurrencyTransactionActivityDetailType `xml:"CurrencyTransactionActivityDetail"`
	SeqNum                            int64                                   `xml:"SeqNum,attr"`
}

func (r CurrencyTransactionActivity) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CurrencyTransactionActivityDetailType struct {
	CurrencyTransactionActivityDetailTypeCode *fincen.ValidateCurrencyTransactionActivityDetailCodeType `xml:"CurrencyTransactionActivityDetailTypeCode,omitempty" json:",omitempty"`
	DetailTransactionAmountText               *fincen.RestrictString15                                  `xml:"DetailTransactionAmountText,omitempty" json:",omitempty"`
	DetailTransactionDescription              string                                                    `xml:"DetailTransactionDescription,omitempty" json:",omitempty"`
	InstrumentProductServiceTypeCode          *fincen.ValidateInstrumentProductServiceTypeCode          `xml:"InstrumentProductServiceTypeCode,omitempty" json:",omitempty"`
	IssuerNameText                            string                                                    `xml:"IssuerNameText,omitempty" json:",omitempty"`
	OtherForeignCurrencyCountryText           *fincen.RestrictString2                                   `xml:"OtherForeignCurrencyCountryText,omitempty" json:",omitempty"`
	SeqNum                                    int64                                                     `xml:"SeqNum,attr"`
}

func (r CurrencyTransactionActivityDetailType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type CurrencyTransactionActivityType struct {
	InstallmentPaymentOtherIndicator *fincen.ValidateIndicatorType `xml:"InstallmentPaymentOtherIndicator,omitempty" json:",omitempty"`
	Total100DollarBillInAmountText   string                        `xml:"Total100DollarBillInAmountText,omitempty" json:",omitempty"`
	TotalCashInReceiveAmountText     string                        `xml:"TotalCashInReceiveAmountText"`
	TotalPriceAmountText             string                        `xml:"TotalPriceAmountText,omitempty" json:",omitempty"`
	TransactionDateText              fincen.DateYYYYMMDDType       `xml:"TransactionDateText"`
	SeqNum                           int64                         `xml:"SeqNum,attr"`
}

func (r CurrencyTransactionActivityType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type Party struct {
	ActivityPartyTypeCode   ValidateActivityPartyCodeType   `xml:"ActivityPartyTypeCode"`
	IndividualBirthDateText *fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	PartyTypeCode           *fincen.ValidatePartyTypeCode   `xml:"PartyTypeCode,omitempty" json:",omitempty"`
	PartyName               []PartyNameType                 `xml:"PartyName,omitempty" json:",omitempty"`
	Address                 *AddressType                    `xml:"Address,omitempty" json:",omitempty"`
	PhoneNumber             *PhoneNumberType                `xml:"PhoneNumber,omitempty" json:",omitempty"`
	PartyIdentification     []PartyIdentificationType       `xml:"PartyIdentification,omitempty" json:",omitempty"`
	PartyOccupationBusiness *PartyOccupationBusinessType    `xml:"PartyOccupationBusiness,omitempty" json:",omitempty"`
	SeqNum                  int64                           `xml:"SeqNum,attr"`
}

func (r Party) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyIdentificationType struct {
	OtherIssuerStateText          *fincen.RestrictString3             `xml:"OtherIssuerStateText,omitempty" json:",omitempty"`
	PartyIdentificationNumberText *fincen.RestrictString25            `xml:"PartyIdentificationNumberText,omitempty" json:",omitempty"`
	PartyIdentificationTypeCode   ValidatePartyIdentificationCodeType `xml:"PartyIdentificationTypeCode"`
	SeqNum                        int64                               `xml:"SeqNum,attr"`
}

func (r PartyIdentificationType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyNameType struct {
	PartyNameTypeCode           *ValidatePartyNameCodeType `xml:"PartyNameTypeCode,omitempty" json:",omitempty"`
	RawEntityIndividualLastName *fincen.RestrictString150  `xml:"RawEntityIndividualLastName,omitempty" json:",omitempty"`
	RawIndividualFirstName      *fincen.RestrictString35   `xml:"RawIndividualFirstName,omitempty" json:",omitempty"`
	RawIndividualMiddleName     *fincen.RestrictString35   `xml:"RawIndividualMiddleName,omitempty" json:",omitempty"`
	RawIndividualTitleText      *fincen.RestrictString35   `xml:"RawIndividualTitleText,omitempty" json:",omitempty"`
	RawPartyFullName            *fincen.RestrictString150  `xml:"RawPartyFullName,omitempty" json:",omitempty"`
	SeqNum                      int64                      `xml:"SeqNum,attr"`
}

func (r PartyNameType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyOccupationBusinessType struct {
	OccupationBusinessText fincen.RestrictString30 `xml:"OccupationBusinessText"`
	SeqNum                 int64                   `xml:"SeqNum,attr"`
}

func (r PartyOccupationBusinessType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PartyType struct {
	ActivityPartyTypeCode   ValidateActivityPartyCodeType   `xml:"ActivityPartyTypeCode"`
	IndividualBirthDateText *fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty" json:",omitempty"`
	PartyTypeCode           *fincen.ValidatePartyTypeCode   `xml:"PartyTypeCode,omitempty" json:",omitempty"`
	SeqNum                  int64                           `xml:"SeqNum,attr"`
}

func (r PartyType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

type PhoneNumberType struct {
	PhoneNumberText fincen.RestrictString16 `xml:"PhoneNumberText"`
	SeqNum          int64                   `xml:"SeqNum,attr"`
}

func (r PhoneNumberType) Validate(args ...string) error {
	return fincen.Validate(&r, args...)
}

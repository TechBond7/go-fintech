// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package form8300

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

type Activity struct {
	EFilingPriorDocumentNumber                  int64                            `xml:"EFilingPriorDocumentNumber,omitempty"`
	FilingDateText                              fincen.DateYYYYMMDDType          `xml:"FilingDateText"`
	MultipleSubjectsIndicator                   fincen.ValidateIndicatorType     `xml:"MultipleSubjectsIndicator,omitempty"`
	SuspiciousTransactionIndicator              fincen.ValidateIndicatorType     `xml:"SuspiciousTransactionIndicator,omitempty"`
	TransactionOnBehalfMultiplePersonsIndicator fincen.ValidateIndicatorType     `xml:"TransactionOnBehalfMultiplePersonsIndicator,omitempty"`
	ActivityAssociation                         ActivityAssociationType          `xml:"ActivityAssociation"`
	Party                                       []string                         `xml:"Party"`
	CurrencyTransactionActivity                 string                           `xml:"CurrencyTransactionActivity"`
	ActivityNarrativeInformation                ActivityNarrativeInformationType `xml:"ActivityNarrativeInformation,omitempty"`
	SeqNum                                      int64                            `xml:"SeqNum,attr"`
}

type ActivityAssociationType struct {
	CorrectsAmendsPriorReportIndicator fincen.ValidateIndicatorType `xml:"CorrectsAmendsPriorReportIndicator,omitempty"`
	InitialReportIndicator             fincen.ValidateIndicatorType `xml:"InitialReportIndicator,omitempty"`
	SeqNum                             int64                        `xml:"SeqNum,attr"`
}

type ActivityNarrativeInformationType struct {
	ActivityNarrativeSequenceNumber ValidateActivityNarrativeSequenceNumber `xml:"ActivityNarrativeSequenceNumber"`
	ActivityNarrativeText           fincen.RestrictString750                `xml:"ActivityNarrativeText"`
	SeqNum                          int64                                   `xml:"SeqNum,attr"`
}

type ActivityType struct {
	EFilingPriorDocumentNumber                  int64                        `xml:"EFilingPriorDocumentNumber,omitempty"`
	FilingDateText                              fincen.DateYYYYMMDDType      `xml:"FilingDateText"`
	MultipleSubjectsIndicator                   fincen.ValidateIndicatorType `xml:"MultipleSubjectsIndicator,omitempty"`
	SuspiciousTransactionIndicator              fincen.ValidateIndicatorType `xml:"SuspiciousTransactionIndicator,omitempty"`
	TransactionOnBehalfMultiplePersonsIndicator fincen.ValidateIndicatorType `xml:"TransactionOnBehalfMultiplePersonsIndicator,omitempty"`
	SeqNum                                      int64                        `xml:"SeqNum,attr"`
}

type AddressType struct {
	RawCityText           fincen.RestrictString50  `xml:"RawCityText,omitempty"`
	RawCountryCodeText    fincen.RestrictString2   `xml:"RawCountryCodeText,omitempty"`
	RawStateCodeText      fincen.RestrictString3   `xml:"RawStateCodeText,omitempty"`
	RawStreetAddress1Text fincen.RestrictString100 `xml:"RawStreetAddress1Text,omitempty"`
	RawZIPCode            fincen.RestrictString9   `xml:"RawZIPCode,omitempty"`
	SeqNum                int64                    `xml:"SeqNum,attr"`
}

type CurrencyTransactionActivity struct {
	InstallmentPaymentOtherIndicator  fincen.ValidateIndicatorType            `xml:"InstallmentPaymentOtherIndicator,omitempty"`
	Total100DollarBillInAmountText    string                                  `xml:"Total100DollarBillInAmountText,omitempty"`
	TotalCashInReceiveAmountText      string                                  `xml:"TotalCashInReceiveAmountText"`
	TotalPriceAmountText              string                                  `xml:"TotalPriceAmountText,omitempty"`
	TransactionDateText               fincen.DateYYYYMMDDType                 `xml:"TransactionDateText"`
	CurrencyTransactionActivityDetail []CurrencyTransactionActivityDetailType `xml:"CurrencyTransactionActivityDetail"`
	SeqNum                            int64                                   `xml:"SeqNum,attr"`
}

type CurrencyTransactionActivityDetailType struct {
	CurrencyTransactionActivityDetailTypeCode fincen.ValidateCurrencyTransactionActivityDetailCodeType `xml:"CurrencyTransactionActivityDetailTypeCode,omitempty"`
	DetailTransactionAmountText               fincen.RestrictString15                                  `xml:"DetailTransactionAmountText,omitempty"`
	DetailTransactionDescription              string                                                   `xml:"DetailTransactionDescription,omitempty"`
	InstrumentProductServiceTypeCode          fincen.ValidateInstrumentProductServiceTypeCode          `xml:"InstrumentProductServiceTypeCode,omitempty"`
	IssuerNameText                            string                                                   `xml:"IssuerNameText,omitempty"`
	OtherForeignCurrencyCountryText           fincen.RestrictString2                                   `xml:"OtherForeignCurrencyCountryText,omitempty"`
	SeqNum                                    int64                                                    `xml:"SeqNum,attr"`
}

type CurrencyTransactionActivityType struct {
	InstallmentPaymentOtherIndicator fincen.ValidateIndicatorType `xml:"InstallmentPaymentOtherIndicator,omitempty"`
	Total100DollarBillInAmountText   string                       `xml:"Total100DollarBillInAmountText,omitempty"`
	TotalCashInReceiveAmountText     string                       `xml:"TotalCashInReceiveAmountText"`
	TotalPriceAmountText             string                       `xml:"TotalPriceAmountText,omitempty"`
	TransactionDateText              fincen.DateYYYYMMDDType      `xml:"TransactionDateText"`
	SeqNum                           int64                        `xml:"SeqNum,attr"`
}

type Party struct {
	ActivityPartyTypeCode   ValidateActivityPartyCodeType  `xml:"ActivityPartyTypeCode"`
	IndividualBirthDateText fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty"`
	PartyTypeCode           fincen.ValidatePartyTypeCode   `xml:"PartyTypeCode,omitempty"`
	PartyName               []PartyNameType                `xml:"PartyName,omitempty"`
	Address                 AddressType                    `xml:"Address,omitempty"`
	PhoneNumber             PhoneNumberType                `xml:"PhoneNumber,omitempty"`
	PartyIdentification     []PartyIdentificationType      `xml:"PartyIdentification,omitempty"`
	PartyOccupationBusiness PartyOccupationBusinessType    `xml:"PartyOccupationBusiness,omitempty"`
	SeqNum                  int64                          `xml:"SeqNum,attr"`
}

type PartyIdentificationType struct {
	OtherIssuerStateText          fincen.RestrictString3              `xml:"OtherIssuerStateText,omitempty"`
	PartyIdentificationNumberText fincen.RestrictString25             `xml:"PartyIdentificationNumberText,omitempty"`
	PartyIdentificationTypeCode   ValidatePartyIdentificationCodeType `xml:"PartyIdentificationTypeCode"`
	SeqNum                        int64                               `xml:"SeqNum,attr"`
}

type PartyNameType struct {
	PartyNameTypeCode           ValidatePartyNameCodeType `xml:"PartyNameTypeCode,omitempty"`
	RawEntityIndividualLastName fincen.RestrictString150  `xml:"RawEntityIndividualLastName,omitempty"`
	RawIndividualFirstName      fincen.RestrictString35   `xml:"RawIndividualFirstName,omitempty"`
	RawIndividualMiddleName     fincen.RestrictString35   `xml:"RawIndividualMiddleName,omitempty"`
	RawIndividualTitleText      fincen.RestrictString35   `xml:"RawIndividualTitleText,omitempty"`
	RawPartyFullName            fincen.RestrictString150  `xml:"RawPartyFullName,omitempty"`
	SeqNum                      int64                     `xml:"SeqNum,attr"`
}

type PartyOccupationBusinessType struct {
	OccupationBusinessText fincen.RestrictString30 `xml:"OccupationBusinessText"`
	SeqNum                 int64                   `xml:"SeqNum,attr"`
}

type PartyType struct {
	ActivityPartyTypeCode   ValidateActivityPartyCodeType  `xml:"ActivityPartyTypeCode"`
	IndividualBirthDateText fincen.DateYYYYMMDDOrBlankType `xml:"IndividualBirthDateText,omitempty"`
	PartyTypeCode           fincen.ValidatePartyTypeCode   `xml:"PartyTypeCode,omitempty"`
	SeqNum                  int64                          `xml:"SeqNum,attr"`
}

type PhoneNumberType struct {
	PhoneNumberText fincen.RestrictString16 `xml:"PhoneNumberText"`
	SeqNum          int64                   `xml:"SeqNum,attr"`
}

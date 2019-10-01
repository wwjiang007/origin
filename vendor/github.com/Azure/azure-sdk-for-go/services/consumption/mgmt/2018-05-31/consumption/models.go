package consumption

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2018-05-31/consumption"

// Bound enumerates the values for bound.
type Bound string

const (
	// Lower ...
	Lower Bound = "Lower"
	// Upper ...
	Upper Bound = "Upper"
)

// PossibleBoundValues returns an array of possible values for the Bound const type.
func PossibleBoundValues() []Bound {
	return []Bound{Lower, Upper}
}

// ChargeType enumerates the values for charge type.
type ChargeType string

const (
	// ChargeTypeActual ...
	ChargeTypeActual ChargeType = "Actual"
	// ChargeTypeForecast ...
	ChargeTypeForecast ChargeType = "Forecast"
)

// PossibleChargeTypeValues returns an array of possible values for the ChargeType const type.
func PossibleChargeTypeValues() []ChargeType {
	return []ChargeType{ChargeTypeActual, ChargeTypeForecast}
}

// Grain enumerates the values for grain.
type Grain string

const (
	// Daily ...
	Daily Grain = "Daily"
	// Monthly ...
	Monthly Grain = "Monthly"
	// Yearly ...
	Yearly Grain = "Yearly"
)

// PossibleGrainValues returns an array of possible values for the Grain const type.
func PossibleGrainValues() []Grain {
	return []Grain{Daily, Monthly, Yearly}
}

// ErrorDetails the details of the error.
type ErrorDetails struct {
	// Code - READ-ONLY; Error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// ErrorResponse error response indicates that the service is not able to process the incoming request. The
// reason is provided in the error message.
type ErrorResponse struct {
	// Error - The details of the error.
	Error *ErrorDetails `json:"error,omitempty"`
}

// Forecast a forecast resource.
type Forecast struct {
	*ForecastProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Forecast.
func (f Forecast) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if f.ForecastProperties != nil {
		objectMap["properties"] = f.ForecastProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Forecast struct.
func (f *Forecast) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var forecastProperties ForecastProperties
				err = json.Unmarshal(*v, &forecastProperties)
				if err != nil {
					return err
				}
				f.ForecastProperties = &forecastProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				f.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				f.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				f.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				f.Tags = tags
			}
		}
	}

	return nil
}

// ForecastProperties the properties of the forecast charge.
type ForecastProperties struct {
	// UsageDate - READ-ONLY; The usage date of the forecast.
	UsageDate *string `json:"usageDate,omitempty"`
	// Grain - The granularity of forecast. Possible values include: 'Daily', 'Monthly', 'Yearly'
	Grain Grain `json:"grain,omitempty"`
	// Charge - READ-ONLY; The amount of charge
	Charge *decimal.Decimal `json:"charge,omitempty"`
	// Currency - READ-ONLY; The ISO currency in which the meter is charged, for example, USD.
	Currency *string `json:"currency,omitempty"`
	// ChargeType - The type of the charge. Could be actual or forecast. Possible values include: 'ChargeTypeActual', 'ChargeTypeForecast'
	ChargeType ChargeType `json:"chargeType,omitempty"`
	// ConfidenceLevels - READ-ONLY; The details about the forecast confidence levels. This is populated only when chargeType is Forecast.
	ConfidenceLevels *[]ForecastPropertiesConfidenceLevelsItem `json:"confidenceLevels,omitempty"`
}

// ForecastPropertiesConfidenceLevelsItem ...
type ForecastPropertiesConfidenceLevelsItem struct {
	// Percentage - READ-ONLY; The percentage level of the confidence
	Percentage *decimal.Decimal `json:"percentage,omitempty"`
	// Bound - The boundary of the percentage, values could be 'Upper' or 'Lower'. Possible values include: 'Upper', 'Lower'
	Bound Bound `json:"bound,omitempty"`
	// Value - READ-ONLY; The amount of forecast within the percentage level
	Value *decimal.Decimal `json:"value,omitempty"`
}

// ForecastsListResult result of listing forecasts. It contains a list of available forecasts.
type ForecastsListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The list of forecasts.
	Value *[]Forecast `json:"value,omitempty"`
}

// MeterDetails the properties of the meter detail.
type MeterDetails struct {
	// MeterName - READ-ONLY; The name of the meter, within the given meter category
	MeterName *string `json:"meterName,omitempty"`
	// MeterCategory - READ-ONLY; The category of the meter, for example, 'Cloud services', 'Networking', etc..
	MeterCategory *string `json:"meterCategory,omitempty"`
	// MeterSubCategory - READ-ONLY; The subcategory of the meter, for example, 'A6 Cloud services', 'ExpressRoute (IXP)', etc..
	MeterSubCategory *string `json:"meterSubCategory,omitempty"`
	// Unit - READ-ONLY; The unit in which the meter consumption is charged, for example, 'Hours', 'GB', etc.
	Unit *string `json:"unit,omitempty"`
	// MeterLocation - READ-ONLY; The location in which the Azure service is available.
	MeterLocation *string `json:"meterLocation,omitempty"`
	// TotalIncludedQuantity - READ-ONLY; The total included quantity associated with the offer.
	TotalIncludedQuantity *decimal.Decimal `json:"totalIncludedQuantity,omitempty"`
	// PretaxStandardRate - READ-ONLY; The pretax listing price.
	PretaxStandardRate *decimal.Decimal `json:"pretaxStandardRate,omitempty"`
}

// Operation a Consumption REST API operation.
type Operation struct {
	// Name - READ-ONLY; Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - READ-ONLY; Service provider: Microsoft.Consumption.
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; Resource on which the operation is performed: UsageDetail, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of listing consumption operations. It contains a list of operations and a URL
// link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of consumption operations supported by the Microsoft.Consumption resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{fn: getNextPage}
}

// PriceSheetModel price sheet result. It contains the pricesheet associated with billing period
type PriceSheetModel struct {
	// Pricesheets - READ-ONLY; Price sheet
	Pricesheets *[]PriceSheetProperties `json:"pricesheets,omitempty"`
	// NextLink - READ-ONLY; The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// PriceSheetProperties the properties of the price sheet.
type PriceSheetProperties struct {
	// BillingPeriodID - READ-ONLY; The id of the billing period resource that the usage belongs to.
	BillingPeriodID *string `json:"billingPeriodId,omitempty"`
	// MeterID - READ-ONLY; The meter id (GUID)
	MeterID *uuid.UUID `json:"meterId,omitempty"`
	// MeterDetails - READ-ONLY; The details about the meter. By default this is not populated, unless it's specified in $expand.
	MeterDetails *MeterDetails `json:"meterDetails,omitempty"`
	// UnitOfMeasure - READ-ONLY; Unit of measure
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
	// IncludedQuantity - READ-ONLY; Included quality for an offer
	IncludedQuantity *decimal.Decimal `json:"includedQuantity,omitempty"`
	// PartNumber - READ-ONLY; Part Number
	PartNumber *string `json:"partNumber,omitempty"`
	// UnitPrice - READ-ONLY; Unit Price
	UnitPrice *decimal.Decimal `json:"unitPrice,omitempty"`
	// CurrencyCode - READ-ONLY; Currency Code
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// OfferID - READ-ONLY; Offer Id
	OfferID *string `json:"offerId,omitempty"`
}

// PriceSheetResult an pricesheet resource.
type PriceSheetResult struct {
	autorest.Response `json:"-"`
	*PriceSheetModel  `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for PriceSheetResult.
func (psr PriceSheetResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if psr.PriceSheetModel != nil {
		objectMap["properties"] = psr.PriceSheetModel
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PriceSheetResult struct.
func (psr *PriceSheetResult) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var priceSheetModel PriceSheetModel
				err = json.Unmarshal(*v, &priceSheetModel)
				if err != nil {
					return err
				}
				psr.PriceSheetModel = &priceSheetModel
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				psr.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				psr.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				psr.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				psr.Tags = tags
			}
		}
	}

	return nil
}

// Resource the Resource model definition.
type Resource struct {
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// UsageDetail an usage detail resource.
type UsageDetail struct {
	*UsageDetailProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for UsageDetail.
func (ud UsageDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ud.UsageDetailProperties != nil {
		objectMap["properties"] = ud.UsageDetailProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for UsageDetail struct.
func (ud *UsageDetail) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var usageDetailProperties UsageDetailProperties
				err = json.Unmarshal(*v, &usageDetailProperties)
				if err != nil {
					return err
				}
				ud.UsageDetailProperties = &usageDetailProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ud.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ud.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ud.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ud.Tags = tags
			}
		}
	}

	return nil
}

// UsageDetailProperties the properties of the usage detail.
type UsageDetailProperties struct {
	// BillingPeriodID - READ-ONLY; The id of the billing period resource that the usage belongs to.
	BillingPeriodID *string `json:"billingPeriodId,omitempty"`
	// InvoiceID - READ-ONLY; The id of the invoice resource that the usage belongs to.
	InvoiceID *string `json:"invoiceId,omitempty"`
	// UsageStart - READ-ONLY; The start of the date time range covered by the usage detail.
	UsageStart *date.Time `json:"usageStart,omitempty"`
	// UsageEnd - READ-ONLY; The end of the date time range covered by the usage detail.
	UsageEnd *date.Time `json:"usageEnd,omitempty"`
	// InstanceName - READ-ONLY; The name of the resource instance that the usage is about.
	InstanceName *string `json:"instanceName,omitempty"`
	// InstanceID - READ-ONLY; The uri of the resource instance that the usage is about.
	InstanceID *string `json:"instanceId,omitempty"`
	// InstanceLocation - READ-ONLY; The location of the resource instance that the usage is about.
	InstanceLocation *string `json:"instanceLocation,omitempty"`
	// Currency - READ-ONLY; The ISO currency in which the meter is charged, for example, USD.
	Currency *string `json:"currency,omitempty"`
	// UsageQuantity - READ-ONLY; The quantity of usage.
	UsageQuantity *decimal.Decimal `json:"usageQuantity,omitempty"`
	// BillableQuantity - READ-ONLY; The billable usage quantity.
	BillableQuantity *decimal.Decimal `json:"billableQuantity,omitempty"`
	// PretaxCost - READ-ONLY; The amount of cost before tax.
	PretaxCost *decimal.Decimal `json:"pretaxCost,omitempty"`
	// IsEstimated - READ-ONLY; The estimated usage is subject to change.
	IsEstimated *bool `json:"isEstimated,omitempty"`
	// MeterID - READ-ONLY; The meter id (GUID).
	MeterID *uuid.UUID `json:"meterId,omitempty"`
	// MeterDetails - READ-ONLY; The details about the meter. By default this is not populated, unless it's specified in $expand.
	MeterDetails *MeterDetails `json:"meterDetails,omitempty"`
	// SubscriptionGUID - READ-ONLY; Subscription guid.
	SubscriptionGUID *uuid.UUID `json:"subscriptionGuid,omitempty"`
	// SubscriptionName - READ-ONLY; Subscription name.
	SubscriptionName *string `json:"subscriptionName,omitempty"`
	// AccountName - READ-ONLY; Account name.
	AccountName *string `json:"accountName,omitempty"`
	// DepartmentName - READ-ONLY; Department name.
	DepartmentName *string `json:"departmentName,omitempty"`
	// Product - READ-ONLY; Product name.
	Product *string `json:"product,omitempty"`
	// ConsumedService - READ-ONLY; Consumed service name.
	ConsumedService *string `json:"consumedService,omitempty"`
	// CostCenter - READ-ONLY; The cost center of this department if it is a department and a costcenter exists
	CostCenter *string `json:"costCenter,omitempty"`
	// PartNumber - READ-ONLY; Part Number
	PartNumber *string `json:"partNumber,omitempty"`
	// ResourceGUID - READ-ONLY; Resource Guid
	ResourceGUID *string `json:"resourceGuid,omitempty"`
	// OfferID - READ-ONLY; Offer Id
	OfferID *string `json:"offerId,omitempty"`
	// ChargesBilledSeparately - READ-ONLY; Charges billed separately
	ChargesBilledSeparately *bool `json:"chargesBilledSeparately,omitempty"`
	// AdditionalProperties - READ-ONLY; Additional details of this usage item. By default this is not populated, unless it's specified in $expand.
	AdditionalProperties *string `json:"additionalProperties,omitempty"`
}

// UsageDetailsListResult result of listing usage details. It contains a list of available usage details in
// reverse chronological order by billing period.
type UsageDetailsListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The list of usage details.
	Value *[]UsageDetail `json:"value,omitempty"`
	// NextLink - READ-ONLY; The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// UsageDetailsListResultIterator provides access to a complete listing of UsageDetail values.
type UsageDetailsListResultIterator struct {
	i    int
	page UsageDetailsListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *UsageDetailsListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsageDetailsListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *UsageDetailsListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter UsageDetailsListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter UsageDetailsListResultIterator) Response() UsageDetailsListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter UsageDetailsListResultIterator) Value() UsageDetail {
	if !iter.page.NotDone() {
		return UsageDetail{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the UsageDetailsListResultIterator type.
func NewUsageDetailsListResultIterator(page UsageDetailsListResultPage) UsageDetailsListResultIterator {
	return UsageDetailsListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (udlr UsageDetailsListResult) IsEmpty() bool {
	return udlr.Value == nil || len(*udlr.Value) == 0
}

// usageDetailsListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (udlr UsageDetailsListResult) usageDetailsListResultPreparer(ctx context.Context) (*http.Request, error) {
	if udlr.NextLink == nil || len(to.String(udlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(udlr.NextLink)))
}

// UsageDetailsListResultPage contains a page of UsageDetail values.
type UsageDetailsListResultPage struct {
	fn   func(context.Context, UsageDetailsListResult) (UsageDetailsListResult, error)
	udlr UsageDetailsListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *UsageDetailsListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsageDetailsListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.udlr)
	if err != nil {
		return err
	}
	page.udlr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *UsageDetailsListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page UsageDetailsListResultPage) NotDone() bool {
	return !page.udlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page UsageDetailsListResultPage) Response() UsageDetailsListResult {
	return page.udlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page UsageDetailsListResultPage) Values() []UsageDetail {
	if page.udlr.IsEmpty() {
		return nil
	}
	return *page.udlr.Value
}

// Creates a new instance of the UsageDetailsListResultPage type.
func NewUsageDetailsListResultPage(getNextPage func(context.Context, UsageDetailsListResult) (UsageDetailsListResult, error)) UsageDetailsListResultPage {
	return UsageDetailsListResultPage{fn: getNextPage}
}

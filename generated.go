// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"

	"github.com/Khan/genqlient/graphql"
)

type CreateFileRequest struct {
	Purpose FilePurpose `json:"purpose"`
	Url     string      `json:"url"`
}

// GetPurpose returns CreateFileRequest.Purpose, and is useful for accessing the field via an interface.
func (v *CreateFileRequest) GetPurpose() FilePurpose { return v.Purpose }

// GetUrl returns CreateFileRequest.Url, and is useful for accessing the field via an interface.
func (v *CreateFileRequest) GetUrl() string { return v.Url }

type CreateOrderBillingRequest struct {
	Account      string `json:"account"`
	CurrencyCode string `json:"currencyCode"`
}

// GetAccount returns CreateOrderBillingRequest.Account, and is useful for accessing the field via an interface.
func (v *CreateOrderBillingRequest) GetAccount() string { return v.Account }

// GetCurrencyCode returns CreateOrderBillingRequest.CurrencyCode, and is useful for accessing the field via an interface.
func (v *CreateOrderBillingRequest) GetCurrencyCode() string { return v.CurrencyCode }

// CreateOrderCreateOrderGetOrderResponse includes the requested fields of the GraphQL type GetOrderResponse.
type CreateOrderCreateOrderGetOrderResponse struct {
	Fragment_Order `json:"-"`
}

// GetId returns CreateOrderCreateOrderGetOrderResponse.Id, and is useful for accessing the field via an interface.
func (v *CreateOrderCreateOrderGetOrderResponse) GetId() string { return v.Fragment_Order.Id }

// GetOrderNumber returns CreateOrderCreateOrderGetOrderResponse.OrderNumber, and is useful for accessing the field via an interface.
func (v *CreateOrderCreateOrderGetOrderResponse) GetOrderNumber() string {
	return v.Fragment_Order.OrderNumber
}

// GetExternalId returns CreateOrderCreateOrderGetOrderResponse.ExternalId, and is useful for accessing the field via an interface.
func (v *CreateOrderCreateOrderGetOrderResponse) GetExternalId() string {
	return v.Fragment_Order.ExternalId
}

// GetBilling returns CreateOrderCreateOrderGetOrderResponse.Billing, and is useful for accessing the field via an interface.
func (v *CreateOrderCreateOrderGetOrderResponse) GetBilling() Fragment_OrderBillingGetOrderBillingResponse {
	return v.Fragment_Order.Billing
}

// GetShipping returns CreateOrderCreateOrderGetOrderResponse.Shipping, and is useful for accessing the field via an interface.
func (v *CreateOrderCreateOrderGetOrderResponse) GetShipping() Fragment_OrderShippingGetOrderShippingResponse {
	return v.Fragment_Order.Shipping
}

// GetItems returns CreateOrderCreateOrderGetOrderResponse.Items, and is useful for accessing the field via an interface.
func (v *CreateOrderCreateOrderGetOrderResponse) GetItems() []Fragment_OrderItemsGetOrderItemResponse {
	return v.Fragment_Order.Items
}

// GetTracking returns CreateOrderCreateOrderGetOrderResponse.Tracking, and is useful for accessing the field via an interface.
func (v *CreateOrderCreateOrderGetOrderResponse) GetTracking() Fragment_OrderTrackingGetOrderShippingTrackingResponse {
	return v.Fragment_Order.Tracking
}

// GetMeta returns CreateOrderCreateOrderGetOrderResponse.Meta, and is useful for accessing the field via an interface.
func (v *CreateOrderCreateOrderGetOrderResponse) GetMeta() map[string]interface{} {
	return v.Fragment_Order.Meta
}

func (v *CreateOrderCreateOrderGetOrderResponse) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CreateOrderCreateOrderGetOrderResponse
		graphql.NoUnmarshalJSON
	}
	firstPass.CreateOrderCreateOrderGetOrderResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.Fragment_Order)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalCreateOrderCreateOrderGetOrderResponse struct {
	Id string `json:"id"`

	OrderNumber string `json:"orderNumber"`

	ExternalId string `json:"externalId"`

	Billing Fragment_OrderBillingGetOrderBillingResponse `json:"billing"`

	Shipping Fragment_OrderShippingGetOrderShippingResponse `json:"shipping"`

	Items []Fragment_OrderItemsGetOrderItemResponse `json:"items"`

	Tracking Fragment_OrderTrackingGetOrderShippingTrackingResponse `json:"tracking"`

	Meta map[string]interface{} `json:"meta"`
}

func (v *CreateOrderCreateOrderGetOrderResponse) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CreateOrderCreateOrderGetOrderResponse) __premarshalJSON() (*__premarshalCreateOrderCreateOrderGetOrderResponse, error) {
	var retval __premarshalCreateOrderCreateOrderGetOrderResponse

	retval.Id = v.Fragment_Order.Id
	retval.OrderNumber = v.Fragment_Order.OrderNumber
	retval.ExternalId = v.Fragment_Order.ExternalId
	retval.Billing = v.Fragment_Order.Billing
	retval.Shipping = v.Fragment_Order.Shipping
	retval.Items = v.Fragment_Order.Items
	retval.Tracking = v.Fragment_Order.Tracking
	retval.Meta = v.Fragment_Order.Meta
	return &retval, nil
}

type CreateOrderItemRequest struct {
	ExternalId  string                 `json:"externalId"`
	Files       []CreateFileRequest    `json:"files"`
	Height      float64                `json:"height"`
	MaterialSku string                 `json:"materialSku"`
	Meta        map[string]interface{} `json:"meta"`
	ProductSku  string                 `json:"productSku"`
	Quantity    int                    `json:"quantity"`
	Shape       Shape                  `json:"shape"`
	Unit        Unit                   `json:"unit"`
	Width       float64                `json:"width"`
}

// GetExternalId returns CreateOrderItemRequest.ExternalId, and is useful for accessing the field via an interface.
func (v *CreateOrderItemRequest) GetExternalId() string { return v.ExternalId }

// GetFiles returns CreateOrderItemRequest.Files, and is useful for accessing the field via an interface.
func (v *CreateOrderItemRequest) GetFiles() []CreateFileRequest { return v.Files }

// GetHeight returns CreateOrderItemRequest.Height, and is useful for accessing the field via an interface.
func (v *CreateOrderItemRequest) GetHeight() float64 { return v.Height }

// GetMaterialSku returns CreateOrderItemRequest.MaterialSku, and is useful for accessing the field via an interface.
func (v *CreateOrderItemRequest) GetMaterialSku() string { return v.MaterialSku }

// GetMeta returns CreateOrderItemRequest.Meta, and is useful for accessing the field via an interface.
func (v *CreateOrderItemRequest) GetMeta() map[string]interface{} { return v.Meta }

// GetProductSku returns CreateOrderItemRequest.ProductSku, and is useful for accessing the field via an interface.
func (v *CreateOrderItemRequest) GetProductSku() string { return v.ProductSku }

// GetQuantity returns CreateOrderItemRequest.Quantity, and is useful for accessing the field via an interface.
func (v *CreateOrderItemRequest) GetQuantity() int { return v.Quantity }

// GetShape returns CreateOrderItemRequest.Shape, and is useful for accessing the field via an interface.
func (v *CreateOrderItemRequest) GetShape() Shape { return v.Shape }

// GetUnit returns CreateOrderItemRequest.Unit, and is useful for accessing the field via an interface.
func (v *CreateOrderItemRequest) GetUnit() Unit { return v.Unit }

// GetWidth returns CreateOrderItemRequest.Width, and is useful for accessing the field via an interface.
func (v *CreateOrderItemRequest) GetWidth() float64 { return v.Width }

type CreateOrderRequest struct {
	Billing    CreateOrderBillingRequest  `json:"billing"`
	ExternalId string                     `json:"externalId"`
	Items      []CreateOrderItemRequest   `json:"items"`
	Meta       map[string]interface{}     `json:"meta"`
	Shipping   CreateOrderShippingRequest `json:"shipping"`
}

// GetBilling returns CreateOrderRequest.Billing, and is useful for accessing the field via an interface.
func (v *CreateOrderRequest) GetBilling() CreateOrderBillingRequest { return v.Billing }

// GetExternalId returns CreateOrderRequest.ExternalId, and is useful for accessing the field via an interface.
func (v *CreateOrderRequest) GetExternalId() string { return v.ExternalId }

// GetItems returns CreateOrderRequest.Items, and is useful for accessing the field via an interface.
func (v *CreateOrderRequest) GetItems() []CreateOrderItemRequest { return v.Items }

// GetMeta returns CreateOrderRequest.Meta, and is useful for accessing the field via an interface.
func (v *CreateOrderRequest) GetMeta() map[string]interface{} { return v.Meta }

// GetShipping returns CreateOrderRequest.Shipping, and is useful for accessing the field via an interface.
func (v *CreateOrderRequest) GetShipping() CreateOrderShippingRequest { return v.Shipping }

// CreateOrderResponse is returned by CreateOrder on success.
type CreateOrderResponse struct {
	CreateOrder CreateOrderCreateOrderGetOrderResponse `json:"createOrder"`
}

// GetCreateOrder returns CreateOrderResponse.CreateOrder, and is useful for accessing the field via an interface.
func (v *CreateOrderResponse) GetCreateOrder() CreateOrderCreateOrderGetOrderResponse {
	return v.CreateOrder
}

type CreateOrderShippingRequest struct {
	City        string `json:"city"`
	CompanyName string `json:"companyName"`
	CountryCode string `json:"countryCode"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
	Phone       string `json:"phone"`
	PostalCode  string `json:"postalCode"`
	ServiceSku  string `json:"serviceSku"`
	State       string `json:"state"`
}

// GetCity returns CreateOrderShippingRequest.City, and is useful for accessing the field via an interface.
func (v *CreateOrderShippingRequest) GetCity() string { return v.City }

// GetCompanyName returns CreateOrderShippingRequest.CompanyName, and is useful for accessing the field via an interface.
func (v *CreateOrderShippingRequest) GetCompanyName() string { return v.CompanyName }

// GetCountryCode returns CreateOrderShippingRequest.CountryCode, and is useful for accessing the field via an interface.
func (v *CreateOrderShippingRequest) GetCountryCode() string { return v.CountryCode }

// GetEmail returns CreateOrderShippingRequest.Email, and is useful for accessing the field via an interface.
func (v *CreateOrderShippingRequest) GetEmail() string { return v.Email }

// GetFirstName returns CreateOrderShippingRequest.FirstName, and is useful for accessing the field via an interface.
func (v *CreateOrderShippingRequest) GetFirstName() string { return v.FirstName }

// GetLastName returns CreateOrderShippingRequest.LastName, and is useful for accessing the field via an interface.
func (v *CreateOrderShippingRequest) GetLastName() string { return v.LastName }

// GetLine1 returns CreateOrderShippingRequest.Line1, and is useful for accessing the field via an interface.
func (v *CreateOrderShippingRequest) GetLine1() string { return v.Line1 }

// GetLine2 returns CreateOrderShippingRequest.Line2, and is useful for accessing the field via an interface.
func (v *CreateOrderShippingRequest) GetLine2() string { return v.Line2 }

// GetPhone returns CreateOrderShippingRequest.Phone, and is useful for accessing the field via an interface.
func (v *CreateOrderShippingRequest) GetPhone() string { return v.Phone }

// GetPostalCode returns CreateOrderShippingRequest.PostalCode, and is useful for accessing the field via an interface.
func (v *CreateOrderShippingRequest) GetPostalCode() string { return v.PostalCode }

// GetServiceSku returns CreateOrderShippingRequest.ServiceSku, and is useful for accessing the field via an interface.
func (v *CreateOrderShippingRequest) GetServiceSku() string { return v.ServiceSku }

// GetState returns CreateOrderShippingRequest.State, and is useful for accessing the field via an interface.
func (v *CreateOrderShippingRequest) GetState() string { return v.State }

type FilePurpose string

const (
	FilePurposeCustomerartwork      FilePurpose = "customerArtwork"
	FilePurposeCutfile              FilePurpose = "cutFile"
	FilePurposeDesignworkingfile    FilePurpose = "designWorkingFile"
	FilePurposeGeneric              FilePurpose = "generic"
	FilePurposeGraphicdesign        FilePurpose = "graphicDesign"
	FilePurposeGraphicproof         FilePurpose = "graphicProof"
	FilePurposePrintreadyartwork    FilePurpose = "printReadyArtwork"
	FilePurposeProductspecification FilePurpose = "productSpecification"
	FilePurposeProof                FilePurpose = "proof"
	FilePurposeTiliareadyartwork    FilePurpose = "tiliaReadyArtwork"
)

// Fragment_Order includes the GraphQL fields of GetOrderResponse requested by the fragment Fragment_Order.
type Fragment_Order struct {
	Id          string                                                 `json:"id"`
	OrderNumber string                                                 `json:"orderNumber"`
	ExternalId  string                                                 `json:"externalId"`
	Billing     Fragment_OrderBillingGetOrderBillingResponse           `json:"billing"`
	Shipping    Fragment_OrderShippingGetOrderShippingResponse         `json:"shipping"`
	Items       []Fragment_OrderItemsGetOrderItemResponse              `json:"items"`
	Tracking    Fragment_OrderTrackingGetOrderShippingTrackingResponse `json:"tracking"`
	Meta        map[string]interface{}                                 `json:"meta"`
}

// GetId returns Fragment_Order.Id, and is useful for accessing the field via an interface.
func (v *Fragment_Order) GetId() string { return v.Id }

// GetOrderNumber returns Fragment_Order.OrderNumber, and is useful for accessing the field via an interface.
func (v *Fragment_Order) GetOrderNumber() string { return v.OrderNumber }

// GetExternalId returns Fragment_Order.ExternalId, and is useful for accessing the field via an interface.
func (v *Fragment_Order) GetExternalId() string { return v.ExternalId }

// GetBilling returns Fragment_Order.Billing, and is useful for accessing the field via an interface.
func (v *Fragment_Order) GetBilling() Fragment_OrderBillingGetOrderBillingResponse { return v.Billing }

// GetShipping returns Fragment_Order.Shipping, and is useful for accessing the field via an interface.
func (v *Fragment_Order) GetShipping() Fragment_OrderShippingGetOrderShippingResponse {
	return v.Shipping
}

// GetItems returns Fragment_Order.Items, and is useful for accessing the field via an interface.
func (v *Fragment_Order) GetItems() []Fragment_OrderItemsGetOrderItemResponse { return v.Items }

// GetTracking returns Fragment_Order.Tracking, and is useful for accessing the field via an interface.
func (v *Fragment_Order) GetTracking() Fragment_OrderTrackingGetOrderShippingTrackingResponse {
	return v.Tracking
}

// GetMeta returns Fragment_Order.Meta, and is useful for accessing the field via an interface.
func (v *Fragment_Order) GetMeta() map[string]interface{} { return v.Meta }

// Fragment_OrderBillingGetOrderBillingResponse includes the requested fields of the GraphQL type GetOrderBillingResponse.
type Fragment_OrderBillingGetOrderBillingResponse struct {
	Account string `json:"account"`
}

// GetAccount returns Fragment_OrderBillingGetOrderBillingResponse.Account, and is useful for accessing the field via an interface.
func (v *Fragment_OrderBillingGetOrderBillingResponse) GetAccount() string { return v.Account }

// Fragment_OrderItemsGetOrderItemResponse includes the requested fields of the GraphQL type GetOrderItemResponse.
type Fragment_OrderItemsGetOrderItemResponse struct {
	Id          string                                                `json:"id"`
	DesignCode  string                                                `json:"designCode"`
	ExternalId  string                                                `json:"externalId"`
	ProductSku  string                                                `json:"productSku"`
	MaterialSku string                                                `json:"materialSku"`
	Quantity    int                                                   `json:"quantity"`
	Width       float64                                               `json:"width"`
	Height      float64                                               `json:"height"`
	Unit        Unit                                                  `json:"unit"`
	Files       []Fragment_OrderItemsGetOrderItemResponseFilesFile    `json:"files"`
	Meta        []Fragment_OrderItemsGetOrderItemResponseMetaMetaItem `json:"meta"`
}

// GetId returns Fragment_OrderItemsGetOrderItemResponse.Id, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponse) GetId() string { return v.Id }

// GetDesignCode returns Fragment_OrderItemsGetOrderItemResponse.DesignCode, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponse) GetDesignCode() string { return v.DesignCode }

// GetExternalId returns Fragment_OrderItemsGetOrderItemResponse.ExternalId, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponse) GetExternalId() string { return v.ExternalId }

// GetProductSku returns Fragment_OrderItemsGetOrderItemResponse.ProductSku, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponse) GetProductSku() string { return v.ProductSku }

// GetMaterialSku returns Fragment_OrderItemsGetOrderItemResponse.MaterialSku, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponse) GetMaterialSku() string { return v.MaterialSku }

// GetQuantity returns Fragment_OrderItemsGetOrderItemResponse.Quantity, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponse) GetQuantity() int { return v.Quantity }

// GetWidth returns Fragment_OrderItemsGetOrderItemResponse.Width, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponse) GetWidth() float64 { return v.Width }

// GetHeight returns Fragment_OrderItemsGetOrderItemResponse.Height, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponse) GetHeight() float64 { return v.Height }

// GetUnit returns Fragment_OrderItemsGetOrderItemResponse.Unit, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponse) GetUnit() Unit { return v.Unit }

// GetFiles returns Fragment_OrderItemsGetOrderItemResponse.Files, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponse) GetFiles() []Fragment_OrderItemsGetOrderItemResponseFilesFile {
	return v.Files
}

// GetMeta returns Fragment_OrderItemsGetOrderItemResponse.Meta, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponse) GetMeta() []Fragment_OrderItemsGetOrderItemResponseMetaMetaItem {
	return v.Meta
}

// Fragment_OrderItemsGetOrderItemResponseFilesFile includes the requested fields of the GraphQL type File.
type Fragment_OrderItemsGetOrderItemResponseFilesFile struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

// GetId returns Fragment_OrderItemsGetOrderItemResponseFilesFile.Id, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponseFilesFile) GetId() string { return v.Id }

// GetUrl returns Fragment_OrderItemsGetOrderItemResponseFilesFile.Url, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponseFilesFile) GetUrl() string { return v.Url }

// Fragment_OrderItemsGetOrderItemResponseMetaMetaItem includes the requested fields of the GraphQL type MetaItem.
type Fragment_OrderItemsGetOrderItemResponseMetaMetaItem struct {
	Id    int    `json:"id"`
	Ns    string `json:"ns"`
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// GetId returns Fragment_OrderItemsGetOrderItemResponseMetaMetaItem.Id, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponseMetaMetaItem) GetId() int { return v.Id }

// GetNs returns Fragment_OrderItemsGetOrderItemResponseMetaMetaItem.Ns, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponseMetaMetaItem) GetNs() string { return v.Ns }

// GetKey returns Fragment_OrderItemsGetOrderItemResponseMetaMetaItem.Key, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponseMetaMetaItem) GetKey() string { return v.Key }

// GetType returns Fragment_OrderItemsGetOrderItemResponseMetaMetaItem.Type, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponseMetaMetaItem) GetType() string { return v.Type }

// GetValue returns Fragment_OrderItemsGetOrderItemResponseMetaMetaItem.Value, and is useful for accessing the field via an interface.
func (v *Fragment_OrderItemsGetOrderItemResponseMetaMetaItem) GetValue() string { return v.Value }

// Fragment_OrderShippingGetOrderShippingResponse includes the requested fields of the GraphQL type GetOrderShippingResponse.
type Fragment_OrderShippingGetOrderShippingResponse struct {
	City        string `json:"city"`
	CompanyName string `json:"companyName"`
	CountryCode string `json:"countryCode"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
	Phone       string `json:"phone"`
	PostalCode  string `json:"postalCode"`
	ServiceSku  string `json:"serviceSku"`
	State       string `json:"state"`
}

// GetCity returns Fragment_OrderShippingGetOrderShippingResponse.City, and is useful for accessing the field via an interface.
func (v *Fragment_OrderShippingGetOrderShippingResponse) GetCity() string { return v.City }

// GetCompanyName returns Fragment_OrderShippingGetOrderShippingResponse.CompanyName, and is useful for accessing the field via an interface.
func (v *Fragment_OrderShippingGetOrderShippingResponse) GetCompanyName() string {
	return v.CompanyName
}

// GetCountryCode returns Fragment_OrderShippingGetOrderShippingResponse.CountryCode, and is useful for accessing the field via an interface.
func (v *Fragment_OrderShippingGetOrderShippingResponse) GetCountryCode() string {
	return v.CountryCode
}

// GetEmail returns Fragment_OrderShippingGetOrderShippingResponse.Email, and is useful for accessing the field via an interface.
func (v *Fragment_OrderShippingGetOrderShippingResponse) GetEmail() string { return v.Email }

// GetFirstName returns Fragment_OrderShippingGetOrderShippingResponse.FirstName, and is useful for accessing the field via an interface.
func (v *Fragment_OrderShippingGetOrderShippingResponse) GetFirstName() string { return v.FirstName }

// GetLastName returns Fragment_OrderShippingGetOrderShippingResponse.LastName, and is useful for accessing the field via an interface.
func (v *Fragment_OrderShippingGetOrderShippingResponse) GetLastName() string { return v.LastName }

// GetLine1 returns Fragment_OrderShippingGetOrderShippingResponse.Line1, and is useful for accessing the field via an interface.
func (v *Fragment_OrderShippingGetOrderShippingResponse) GetLine1() string { return v.Line1 }

// GetLine2 returns Fragment_OrderShippingGetOrderShippingResponse.Line2, and is useful for accessing the field via an interface.
func (v *Fragment_OrderShippingGetOrderShippingResponse) GetLine2() string { return v.Line2 }

// GetPhone returns Fragment_OrderShippingGetOrderShippingResponse.Phone, and is useful for accessing the field via an interface.
func (v *Fragment_OrderShippingGetOrderShippingResponse) GetPhone() string { return v.Phone }

// GetPostalCode returns Fragment_OrderShippingGetOrderShippingResponse.PostalCode, and is useful for accessing the field via an interface.
func (v *Fragment_OrderShippingGetOrderShippingResponse) GetPostalCode() string { return v.PostalCode }

// GetServiceSku returns Fragment_OrderShippingGetOrderShippingResponse.ServiceSku, and is useful for accessing the field via an interface.
func (v *Fragment_OrderShippingGetOrderShippingResponse) GetServiceSku() string { return v.ServiceSku }

// GetState returns Fragment_OrderShippingGetOrderShippingResponse.State, and is useful for accessing the field via an interface.
func (v *Fragment_OrderShippingGetOrderShippingResponse) GetState() string { return v.State }

// Fragment_OrderTrackingGetOrderShippingTrackingResponse includes the requested fields of the GraphQL type GetOrderShippingTrackingResponse.
type Fragment_OrderTrackingGetOrderShippingTrackingResponse struct {
	Courier string `json:"courier"`
	Code    string `json:"code"`
	Url     string `json:"url"`
}

// GetCourier returns Fragment_OrderTrackingGetOrderShippingTrackingResponse.Courier, and is useful for accessing the field via an interface.
func (v *Fragment_OrderTrackingGetOrderShippingTrackingResponse) GetCourier() string {
	return v.Courier
}

// GetCode returns Fragment_OrderTrackingGetOrderShippingTrackingResponse.Code, and is useful for accessing the field via an interface.
func (v *Fragment_OrderTrackingGetOrderShippingTrackingResponse) GetCode() string { return v.Code }

// GetUrl returns Fragment_OrderTrackingGetOrderShippingTrackingResponse.Url, and is useful for accessing the field via an interface.
func (v *Fragment_OrderTrackingGetOrderShippingTrackingResponse) GetUrl() string { return v.Url }

// GetOrderOrderByIdGetOrderResponse includes the requested fields of the GraphQL type GetOrderResponse.
type GetOrderOrderByIdGetOrderResponse struct {
	Fragment_Order `json:"-"`
}

// GetId returns GetOrderOrderByIdGetOrderResponse.Id, and is useful for accessing the field via an interface.
func (v *GetOrderOrderByIdGetOrderResponse) GetId() string { return v.Fragment_Order.Id }

// GetOrderNumber returns GetOrderOrderByIdGetOrderResponse.OrderNumber, and is useful for accessing the field via an interface.
func (v *GetOrderOrderByIdGetOrderResponse) GetOrderNumber() string {
	return v.Fragment_Order.OrderNumber
}

// GetExternalId returns GetOrderOrderByIdGetOrderResponse.ExternalId, and is useful for accessing the field via an interface.
func (v *GetOrderOrderByIdGetOrderResponse) GetExternalId() string {
	return v.Fragment_Order.ExternalId
}

// GetBilling returns GetOrderOrderByIdGetOrderResponse.Billing, and is useful for accessing the field via an interface.
func (v *GetOrderOrderByIdGetOrderResponse) GetBilling() Fragment_OrderBillingGetOrderBillingResponse {
	return v.Fragment_Order.Billing
}

// GetShipping returns GetOrderOrderByIdGetOrderResponse.Shipping, and is useful for accessing the field via an interface.
func (v *GetOrderOrderByIdGetOrderResponse) GetShipping() Fragment_OrderShippingGetOrderShippingResponse {
	return v.Fragment_Order.Shipping
}

// GetItems returns GetOrderOrderByIdGetOrderResponse.Items, and is useful for accessing the field via an interface.
func (v *GetOrderOrderByIdGetOrderResponse) GetItems() []Fragment_OrderItemsGetOrderItemResponse {
	return v.Fragment_Order.Items
}

// GetTracking returns GetOrderOrderByIdGetOrderResponse.Tracking, and is useful for accessing the field via an interface.
func (v *GetOrderOrderByIdGetOrderResponse) GetTracking() Fragment_OrderTrackingGetOrderShippingTrackingResponse {
	return v.Fragment_Order.Tracking
}

// GetMeta returns GetOrderOrderByIdGetOrderResponse.Meta, and is useful for accessing the field via an interface.
func (v *GetOrderOrderByIdGetOrderResponse) GetMeta() map[string]interface{} {
	return v.Fragment_Order.Meta
}

func (v *GetOrderOrderByIdGetOrderResponse) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*GetOrderOrderByIdGetOrderResponse
		graphql.NoUnmarshalJSON
	}
	firstPass.GetOrderOrderByIdGetOrderResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.Fragment_Order)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalGetOrderOrderByIdGetOrderResponse struct {
	Id string `json:"id"`

	OrderNumber string `json:"orderNumber"`

	ExternalId string `json:"externalId"`

	Billing Fragment_OrderBillingGetOrderBillingResponse `json:"billing"`

	Shipping Fragment_OrderShippingGetOrderShippingResponse `json:"shipping"`

	Items []Fragment_OrderItemsGetOrderItemResponse `json:"items"`

	Tracking Fragment_OrderTrackingGetOrderShippingTrackingResponse `json:"tracking"`

	Meta map[string]interface{} `json:"meta"`
}

func (v *GetOrderOrderByIdGetOrderResponse) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *GetOrderOrderByIdGetOrderResponse) __premarshalJSON() (*__premarshalGetOrderOrderByIdGetOrderResponse, error) {
	var retval __premarshalGetOrderOrderByIdGetOrderResponse

	retval.Id = v.Fragment_Order.Id
	retval.OrderNumber = v.Fragment_Order.OrderNumber
	retval.ExternalId = v.Fragment_Order.ExternalId
	retval.Billing = v.Fragment_Order.Billing
	retval.Shipping = v.Fragment_Order.Shipping
	retval.Items = v.Fragment_Order.Items
	retval.Tracking = v.Fragment_Order.Tracking
	retval.Meta = v.Fragment_Order.Meta
	return &retval, nil
}

// GetOrderResponse is returned by GetOrder on success.
type GetOrderResponse struct {
	OrderById GetOrderOrderByIdGetOrderResponse `json:"orderById"`
}

// GetOrderById returns GetOrderResponse.OrderById, and is useful for accessing the field via an interface.
func (v *GetOrderResponse) GetOrderById() GetOrderOrderByIdGetOrderResponse { return v.OrderById }

type Shape string

const (
	ShapeCircle       Shape = "circle"
	ShapeCustom       Shape = "custom"
	ShapeOval         Shape = "oval"
	ShapeRectangle    Shape = "rectangle"
	ShapeRoundedconer Shape = "roundedconer"
	ShapeSquare       Shape = "square"
)

type Unit string

const (
	UnitCm Unit = "cm"
	UnitIn Unit = "in"
	UnitMm Unit = "mm"
)

// __CreateOrderInput is used internally by genqlient
type __CreateOrderInput struct {
	Input CreateOrderRequest `json:"input"`
}

// GetInput returns __CreateOrderInput.Input, and is useful for accessing the field via an interface.
func (v *__CreateOrderInput) GetInput() CreateOrderRequest { return v.Input }

// __GetOrderInput is used internally by genqlient
type __GetOrderInput struct {
	OrderId string `json:"orderId"`
}

// GetOrderId returns __GetOrderInput.OrderId, and is useful for accessing the field via an interface.
func (v *__GetOrderInput) GetOrderId() string { return v.OrderId }

// The query or mutation executed by CreateOrder.
const CreateOrder_Operation = `
mutation CreateOrder ($input: CreateOrderRequest!) {
	createOrder(input: $input) {
		... Fragment_Order
	}
}
fragment Fragment_Order on GetOrderResponse {
	id
	orderNumber
	externalId
	billing {
		account
	}
	shipping {
		city
		companyName
		countryCode
		email
		firstName
		lastName
		line1
		line2
		phone
		postalCode
		serviceSku
		state
	}
	items {
		id
		designCode
		externalId
		productSku
		materialSku
		quantity
		width
		height
		unit
		files {
			id
			url
		}
		meta {
			id
			ns
			key
			type
			value
		}
	}
	tracking {
		courier
		code
		url
	}
	meta
}
`

func CreateOrder(
	ctx context.Context,
	client graphql.Client,
	input CreateOrderRequest,
) (*CreateOrderResponse, error) {
	req := &graphql.Request{
		OpName: "CreateOrder",
		Query:  CreateOrder_Operation,
		Variables: &__CreateOrderInput{
			Input: input,
		},
	}
	var err error

	var data CreateOrderResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by GetOrder.
const GetOrder_Operation = `
query GetOrder ($orderId: ID!) {
	orderById(id: $orderId) {
		... Fragment_Order
	}
}
fragment Fragment_Order on GetOrderResponse {
	id
	orderNumber
	externalId
	billing {
		account
	}
	shipping {
		city
		companyName
		countryCode
		email
		firstName
		lastName
		line1
		line2
		phone
		postalCode
		serviceSku
		state
	}
	items {
		id
		designCode
		externalId
		productSku
		materialSku
		quantity
		width
		height
		unit
		files {
			id
			url
		}
		meta {
			id
			ns
			key
			type
			value
		}
	}
	tracking {
		courier
		code
		url
	}
	meta
}
`

func GetOrder(
	ctx context.Context,
	client graphql.Client,
	orderId string,
) (*GetOrderResponse, error) {
	req := &graphql.Request{
		OpName: "GetOrder",
		Query:  GetOrder_Operation,
		Variables: &__GetOrderInput{
			OrderId: orderId,
		},
	}
	var err error

	var data GetOrderResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
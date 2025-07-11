package models

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	GetStatus            *GetShippingStatusRequest     `xml:"GetShippingStatus,omitempty"`
	UpdateAddress        *UpdateDeliveryAddressRequest `xml:"UpdateDeliveryAddress,omitempty"`
	CancelShipping       *CancelShippingOrderRequest   `xml:"CancelShippingOrder,omitempty"`
}

type GetShippingStatusRequest struct {
	TrackingNumber string `xml:"TrackingNumber"`
}

type UpdateDeliveryAddressRequest struct {
	TrackingNumber string `xml:"TrackingNumber"`
	NewAddress     string `xml:"NewAddress"`
}

type CancelShippingOrderRequest struct {
	TrackingNumber string `xml:"TrackingNumber"`
}

package models

type ShippingOrder struct {
	ID             string
	RecipientName  string
	DeliveryAddress string
	ItemDescription string
	Status         string // e.g. "Pending", "In Transit", "Delivered"
}

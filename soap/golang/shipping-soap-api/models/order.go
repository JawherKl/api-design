package models

type ShippingOrder struct {
	ID             string
	RecipientName  string
	DeliveryAddress string
	ItemDescription string
	Status         string // "Pending", "In Transit", "Delivered"
}

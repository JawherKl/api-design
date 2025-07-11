package storage

import (
	"sync"
	"github.com/JawherKl/shipping-soap-api/models"
)

type Store struct {
	sync.Mutex
	orders map[string]models.ShippingOrder
}

func NewStore() *Store {
	return &Store{
		orders: make(map[string]models.ShippingOrder),
	}
}

func (s *Store) CreateOrder(order models.ShippingOrder) {
	s.Lock()
	defer s.Unlock()
	s.orders[order.ID] = order
}

func (s *Store) GetOrder(id string) (models.ShippingOrder, bool) {
	s.Lock()
	defer s.Unlock()
	order, exists := s.orders[id]
	return order, exists
}

func (s *Store) UpdateAddress(id, newAddress string) bool {
	s.Lock()
	defer s.Unlock()

	order, exists := s.orders[id]
	if !exists || order.Status == "Delivered" {
		return false
	}
	order.DeliveryAddress = newAddress
	s.orders[id] = order
	return true
}

func (s *Store) CancelOrder(id string) bool {
	s.Lock()
	defer s.Unlock()

	order, exists := s.orders[id]
	if !exists || order.Status == "Delivered" {
		return false
	}
	order.Status = "Canceled"
	s.orders[id] = order
	return true
}

package readdomain

import "github.com/thisiserico/golabox/domain"

type Repository interface {
	GetAllOrders() []*Order
	CreateOrder(*domain.Order)
	AddItemToOrder(*domain.Order, *domain.Item, *domain.Product)
}

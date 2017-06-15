package memoryds

import (
	"time"

	"github.com/thisiserico/golabox/domain"
)

type WriteClient struct {
	products []*domain.Product
	orders   []*domain.Order
	items    []*domain.Item
	payments []*domain.Payment
}

func NewWriteClient() *WriteClient {
	return &WriteClient{
		products: initializeProducts(),
		orders:   initializeOrders(),
		items:    []*domain.Item{},
		payments: []*domain.Payment{},
	}
}

func (cl *WriteClient) UpsertProduct(product *domain.Product) error {
	for i, p := range cl.products {
		if p.ID.Equals(product.ID) {
			cl.products[i] = product
			return nil
		}
	}

	cl.products = append(cl.products, product)

	return nil
}

func (cl *WriteClient) GetProduct(id domain.ProductId) *domain.Product {
	for _, p := range cl.products {
		if p.ID.Equals(id) {
			return p
		}
	}

	return nil
}

func (cl *WriteClient) GetAllProducts() []*domain.Product {
	return cl.products
}

func (cl *WriteClient) UpsertOrder(order *domain.Order) error {
	for i, o := range cl.orders {
		if o.ID.Equals(order.ID) {
			cl.orders[i] = order
			return nil
		}
	}

	cl.orders = append(cl.orders, order)

	return nil
}

func (cl *WriteClient) GetOrder(id domain.OrderId) *domain.Order {
	for _, o := range cl.orders {
		if o.ID.Equals(id) {
			return o
		}
	}

	return nil
}

func (cl *WriteClient) UpsertOrderItem(item *domain.Item) error {
	for i, oi := range cl.items {
		if oi.ID.Equals(item.ID) {
			cl.items[i] = item
			return nil
		}
	}

	cl.items = append(cl.items, item)

	return nil
}

func (cl *WriteClient) GetOrderItem(id domain.ItemId) *domain.Item {
	for _, oi := range cl.items {
		if oi.ID.Equals(id) {
			return oi
		}
	}

	return nil
}

func (cl *WriteClient) DeleteOrderItem(id domain.ItemId) error {
	for i, oi := range cl.items {
		if oi.ID.Equals(id) {
			cl.items = append(cl.items[:i], cl.items[i+1:]...)
			return nil
		}
	}

	return nil
}

func (cl *WriteClient) UpsertPayment(payment *domain.Payment) error {
	for i, p := range cl.payments {
		if p.ID.Equals(payment.ID) {
			cl.payments[i] = payment
			return nil
		}
	}

	cl.payments = append(cl.payments, payment)

	return nil
}

func initializeProducts() []*domain.Product {
	return []*domain.Product{
		&domain.Product{
			ID:    domain.ProductId("d21de357"),
			Name:  "olives",
			Stock: 100,
		},
		&domain.Product{
			ID:    domain.ProductId("aa53eeab"),
			Name:  "cookies",
			Stock: 35,
		},
		&domain.Product{
			ID:    domain.ProductId("324d6cc9"),
			Name:  "pasta",
			Stock: 76,
		},
		&domain.Product{
			ID:    domain.ProductId("b01fd3f4"),
			Name:  "chocolate",
			Stock: 12,
		},
		&domain.Product{
			ID:    domain.ProductId("f8948e3c"),
			Name:  "oil",
			Stock: 215,
		},
		&domain.Product{
			ID:    domain.ProductId("a088c4bc"),
			Name:  "tomato sauce",
			Stock: 73,
		},
	}
}

func initializeOrders() []*domain.Order {
	return []*domain.Order{
		&domain.Order{
			ID:        domain.OrderId("97ae944e"),
			CreatedAt: time.Now().UTC(),
			Items:     []domain.ItemId{},
			Payment:   "",
		},
	}
}

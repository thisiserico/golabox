package readdomain

import (
	"time"

	"github.com/thisiserico/golabox/domain"
)

type OrderID string

func (id OrderID) Matches(oid domain.OrderId) bool {
	return string(id) == string(oid)
}

type Order struct {
	ID        OrderID   `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Items     []*Item   `json:"items"`
	Payment   *Payment  `json:"payment"`
}

func NewOrder(o *domain.Order) *Order {
	return &Order{
		ID:        OrderID(o.ID),
		CreatedAt: o.CreatedAt,
		Items:     []*Item{},
		Payment:   nil,
	}
}

type ItemID string

type Item struct {
	ID       ItemID   `json:"id"`
	Product  *Product `json:"product"`
	Quantity int      `json:"quantity"`
}

func NewItem(i *domain.Item, p *domain.Product) *Item {
	return &Item{
		ID:       ItemID(i.ID),
		Product:  NewProduct(p),
		Quantity: i.Quantity,
	}
}

type ProductID string

type Product struct {
	ID    ProductID `json:"id"`
	Name  string    `json:"name"`
	Stock int       `json:"stock"`
}

func NewProduct(p *domain.Product) *Product {
	return &Product{
		ID:    ProductID(p.ID),
		Name:  p.Name,
		Stock: p.Stock,
	}
}

type PaymentID string

type Payment struct {
	ID     PaymentID `json:"id"`
	PaidAt time.Time `json:"paid_at"`
}

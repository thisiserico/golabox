package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ProductId string

func NewProductId() ProductId {
	return ProductId(randomId())
}

func (id ProductId) Equals(oid ProductId) bool {
	return string(id) == string(oid)
}

type Product struct {
	ID    ProductId `json:"id"`
	Name  string    `json:"name"`
	Stock int       `json:"stock"`
}

type OrderId string

func NewOrderId() OrderId {
	return OrderId(randomId())
}

func (id OrderId) Equals(oid OrderId) bool {
	return string(id) == string(oid)
}

type Order struct {
	ID        OrderId   `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Items     []ItemId  `json:"items"`
	Payment   PaymentId `json:"payment"`
}

func NewOrder() *Order {
	return &Order{
		ID:        NewOrderId(),
		CreatedAt: time.Now().UTC(),
		Items:     []ItemId{},
		Payment:   "",
	}
}

type ItemId string

func NewItemId() ItemId {
	return ItemId(randomId())
}

func (id ItemId) Equals(iid ItemId) bool {
	return string(id) == string(iid)
}

type Item struct {
	ID       ItemId    `json:"id"`
	Product  ProductId `json:"product"`
	Quantity int       `json:"quantity"`
}

func NewItem(pId ProductId, quantity int) *Item {
	return &Item{
		ID:       NewItemId(),
		Product:  pId,
		Quantity: quantity,
	}
}

func (i *Item) UpdateQuantity(q int) {
	i.Quantity = q
}

type PaymentId string

func NewPaymentId() PaymentId {
	return PaymentId(randomId())
}

func (id PaymentId) Equals(pID PaymentId) bool {
	return string(id) == string(pID)
}

type Payment struct {
	ID     PaymentId `json:"id"`
	PaidAt time.Time `json:"paid_at"`
}

func NewPayment() *Payment {
	return &Payment{
		ID:     NewPaymentId(),
		PaidAt: time.Now().UTC(),
	}
}

func randomId() string {
	return uuid.NewV4().String()[:8]
}

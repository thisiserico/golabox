package memoryds

import (
	"time"

	"github.com/thisiserico/golabox/domain"
	"github.com/thisiserico/golabox/readdomain"
)

type ReadClient struct {
	orders []*readdomain.Order
}

func NewReadClient() *ReadClient {
	return &ReadClient{
		orders: initializeReadOrders(),
	}
}

func (cl *ReadClient) GetAllOrders() []*readdomain.Order {
	return cl.orders
}

func (cl *ReadClient) CreateOrder(o *domain.Order) {
	cl.orders = append(cl.orders, readdomain.NewOrder(o))
}

func (cl *ReadClient) AddItemToOrder(o *domain.Order, oi *domain.Item, p *domain.Product) {
	for i, order := range cl.orders {
		if order.ID.Matches(o.ID) {
			cl.orders[i].Items = append(cl.orders[i].Items, readdomain.NewItem(oi, p))
			return
		}
	}
}

func initializeReadOrders() []*readdomain.Order {
	return []*readdomain.Order{
		&readdomain.Order{
			ID:        readdomain.OrderID("97ae944e"),
			CreatedAt: time.Now().UTC(),
			Items:     []*readdomain.Item{},
		},
	}
}

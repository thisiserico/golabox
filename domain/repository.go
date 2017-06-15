package domain

type CommandRepository interface {
	UpsertProduct(*Product) error
	UpsertOrder(*Order) error
	UpsertOrderItem(*Item) error
	DeleteOrderItem(ItemId) error
	UpsertPayment(*Payment) error
}

type QueryRepository interface {
	GetProduct(ProductId) *Product
	GetAllProducts() []*Product
	GetOrder(OrderId) *Order
	GetOrderItem(ItemId) *Item
}

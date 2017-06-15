package apiservice

import (
	"github.com/justinas/alice"
	"goji.io/pat"
)

func (cl *Client) defineMonitoringRoutes() {
	cl.mux.Handle(
		pat.Get("/status"),
		alice.New().ThenFunc(status),
	)
}

func (cl *Client) defineProductRoutes() {
	cl.mux.Handle(
		pat.Get("/products"),
		alice.New().ThenFunc(cl.getProductCollection),
	)
}

func (cl *Client) defineOrderRoutes() {
	cl.mux.Handle(
		pat.Post("/orders"),
		alice.New().ThenFunc(cl.postOrderResource),
	)

	cl.mux.Handle(
		pat.Post("/orders/:order_id/items"),
		alice.New().ThenFunc(cl.postOrderItemResource),
	)

	cl.mux.Handle(
		pat.Get("/orders/:order_id/items/:item_id"),
		alice.New().ThenFunc(cl.getOrderItemResource),
	)

	cl.mux.Handle(
		pat.Patch("/orders/:order_id/items/:item_id"),
		alice.New().ThenFunc(cl.patchOrderItemResource),
	)

	cl.mux.Handle(
		pat.Delete("/orders/:order_id/items/:item_id"),
		alice.New().ThenFunc(cl.deleteOrderItemResource),
	)
}

func (cl *Client) defineReadModelRoutes() {
	cl.mux.Handle(
		pat.Get("/all-orders"),
		alice.New().ThenFunc(cl.getAllReadModelOrders),
	)
}

func (cl *Client) definePaymentRoutes() {
	cl.mux.Handle(
		pat.Post("/orders/:order_id/payments"),
		alice.New().ThenFunc(cl.postPaymentResource),
	)
}

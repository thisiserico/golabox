package apiservice

import (
	"encoding/json"
	"net/http"

	"goji.io/pat"

	"github.com/thisiserico/golabox/domain"
)

func (cl *Client) getProductCollection(w http.ResponseWriter, r *http.Request) {
	products := cl.domainService.GetAllProducts()

	js, _ := json.Marshal(products)
	w.Write(js)
}

func (cl *Client) postOrderResource(w http.ResponseWriter, r *http.Request) {
	order, err := cl.domainService.CreateNewOrder()
	if err != nil {
		ReturnInternalServerError(w, "ORDER_PERSISTING", err)
		return
	}

	js, _ := json.Marshal(order)
	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}

func (cl *Client) postOrderItemResource(w http.ResponseWriter, r *http.Request) {
	orderID := domain.OrderId(pat.Param(r, "order_id"))
	order, err := cl.domainService.GetOrder(orderID)
	if err != nil {
		ReturnNotFoundError(w, "NON_EXISTING_ORDER", err)
		return
	}

	var rawItem struct {
		Product  domain.ProductId `json:"product"`
		Quantity int              `json:"quantity"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawItem); err != nil {
		ReturnBadRequestError(w, "INVALID_ITEM_PAYLOAD", err)
		return
	}
	defer r.Body.Close()

	if p, err := cl.domainService.GetProduct(rawItem.Product); p == nil {
		ReturnNotFoundError(w, "NON_EXISTING_PRODUCT", err)
		return
	}

	item, err := cl.domainService.CreateNewOrderItem(rawItem.Product, rawItem.Quantity)
	if err != nil {
		ReturnInternalServerError(w, "ITEM_PERSISTING", err)
		return
	}

	if err := cl.domainService.AddItemToOrder(order, item); err != nil {
		ReturnInternalServerError(w, "ORDER_PERSISTING", err)
		return
	}

	js, _ := json.Marshal(item)
	w.Write(js)
}

func (cl *Client) getOrderItemResource(w http.ResponseWriter, r *http.Request) {
	orderID := domain.OrderId(pat.Param(r, "order_id"))
	if _, err := cl.domainService.GetOrder(orderID); err != nil {
		ReturnNotFoundError(w, "NON_EXISTING_ORDER", err)
		return
	}

	itemID := domain.ItemId(pat.Param(r, "item_id"))
	item, err := cl.domainService.GetItem(itemID)
	if err != nil {
		ReturnNotFoundError(w, "NON_EXISTING_ITEM", err)
		return
	}

	js, _ := json.Marshal(item)
	w.Write(js)
}

func (cl *Client) patchOrderItemResource(w http.ResponseWriter, r *http.Request) {
	orderID := domain.OrderId(pat.Param(r, "order_id"))
	if _, err := cl.domainService.GetOrder(orderID); err != nil {
		ReturnNotFoundError(w, "NON_EXISTING_ORDER", err)
		return
	}

	itemID := domain.ItemId(pat.Param(r, "item_id"))
	item, err := cl.domainService.GetItem(itemID)
	if err != nil {
		ReturnNotFoundError(w, "NON_EXISTING_ITEM", err)
		return
	}

	var rawItem struct {
		Quantity int `json:"quantity"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawItem); err != nil {
		ReturnBadRequestError(w, "INVALID_ITEM_PAYLOAD", err)
		return
	}
	defer r.Body.Close()

	if err := cl.domainService.UpdateItemQuantity(item, rawItem.Quantity); err != nil {
		ReturnInternalServerError(w, "ITEM_PERSISTING", err)
		return
	}

	js, _ := json.Marshal(item)
	w.Write(js)
}

func (cl *Client) deleteOrderItemResource(w http.ResponseWriter, r *http.Request) {
	orderID := domain.OrderId(pat.Param(r, "order_id"))
	if _, err := cl.domainService.GetOrder(orderID); err != nil {
		ReturnNotFoundError(w, "NON_EXISTING_ORDER", err)
		return
	}

	itemID := domain.ItemId(pat.Param(r, "item_id"))
	if _, err := cl.domainService.GetItem(itemID); err != nil {
		ReturnNotFoundError(w, "NON_EXISTING_ITEM", err)
		return
	}

	if err := cl.domainService.DeleteItem(itemID); err != nil {
		ReturnInternalServerError(w, "ITEM_REMOVAL", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cl *Client) postPaymentResource(w http.ResponseWriter, r *http.Request) {
	orderID := domain.OrderId(pat.Param(r, "order_id"))
	order, err := cl.domainService.GetOrder(orderID)
	if err != nil {
		ReturnNotFoundError(w, "NON_EXISTING_ORDER", err)
		return
	}

	payment := domain.NewPayment()
	if err := cl.domainService.PayOrder(order, payment); err != nil {
		ReturnBadRequestError(w, "PAYMENT_PROCESSING", err)
		return
	}

	js, _ := json.Marshal(payment)
	w.Write(js)
}

func (cl *Client) getAllReadModelOrders(w http.ResponseWriter, r *http.Request) {
	orders := cl.readClient.GetAllOrders()

	js, _ := json.Marshal(orders)
	w.Write(js)
}

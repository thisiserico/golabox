# Golabox
> Ilustrating how `golang` works for beginners

## Introduction
Having syncronous code that handles updates, we'll make it asynchronous using native go (no queues involved).

The `golabox` service will handle `products`, `orders`, `order items` and `payments`.
An API will expose endpoints to work with those.

The idea is to end up having asynchronous event handling and a (vague) read model for orders. That will save us lots of API "joins".

## How to run
Avoiding the build with `go run main.go` or building the binary with `make all`.

## RESTful API
### `GET /status`
Will return `200` when the API service is ready.

### `GET /products`
Will get a list of products and their respective stock. There's no way to create new products, a default list will be returned.

### `POST /orders`
Will create a new empty order. Not that useful for this example, but totally worth it in a real market place implementation.

### `POST /orders/:order_id/items`
Will create a new line item in the specified order. The desired product and quantity of such product need to be sent.

### `GET /orders/:order_id/items/:item_id`
Will return the requested line item for the specified order.Useful to know how many units of a particular product was requested.

### `PATCH /orders/:order_id/items/:item_id`
Will accept a new quantity for the requested line item of the specified order.

### `DELETE /orders/:order_id/items/:item_id`
Will remove the line item from the order.

### `POST /orders/:order_id/payments`
Will create a new payment for the specified order.

## Read model API
### `GET /all-orders`
Will get all the orders ready to be sent to a client.

## Example flow
Notice how this implementation was intended for ilustrating purposes only, not for real usage!

Start by creating an order:
```
> POST /orders

< 201 Created
< {
<     "id": "b8c316f1",
<     "created_at": "2017-06-15T11:24:57.791741426Z",
<     "items": [],
<     "payment": ""
< }
```

Add some items to the order:
```
> POST /orders/b8c316f1/items
> {
>     "product": "d21de357",
>     "quantity": 2
> }

< 201 Created
< {
<     "id": "e7aa96ae",
<     "product": "d21de357",
<     "quantity": 2
< }
```

Add another item to the order:
```
> POST /orders/b8c316f1/items
> {
>     "product": "a088c4bc",
>     "quantity": 1
> }

< 201 Created
< {
<     "id": "71d92a01",
<     "product": "a088c4bc",
<     "quantity": 1
< }
```

Change the quantity of that particular product:
```
> PATCH /orders/b8c316f1/items/71d92a01
> {
>     "quantity": 3
> }

< 200 Ok
{
<     "id": "71d92a01",
<     "product": "a088c4bc",
<     "quantity": 3
< }
```

Pay and process the order:
```
> POST /orders/b8c316f1/payments

< 200 Ok
< {
<     "id": "1d0ec4bd",
<     "paid_at": "2017-06-15T11:28:27.19354839Z"
< }
```

## Missing bits and pieces
The program is incomplete on purpose so that people can play with it:

- When deleting an `item`, the `order` should lose its reference to such `item`
- An `order` should not be allowed to be modified once a `payment` has been processed
    - Creating another `payment` should be forbidden
    - Adding `items` should be forbidden
    - Modifying `items` should be forbidden
    - Removing `items` should be forbidden
- Creating two line items for the same product should end up being the same line item
- When handling a successful payment the read model should be updated
- Updates in the read model should be handled through events

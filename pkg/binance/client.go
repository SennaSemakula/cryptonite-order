package binance

import (
	"context"
	"cryptonite/pkg/api"
	"cryptonite/pkg/order"
	"errors"
	"fmt"
	"net/http"
)

type Client struct {
	order.Orderer
	sess   api.Client
	orders []order.Order // TODO: change this to be an interafce
}

var (
	InvalidOrderErr = errors.New("order does not exist for account")
	EmptyOrdersErr  = errors.New("no orders exist")
)

func (c *Client) Do(req *http.Request) (http.Response, error) {
	return http.Response{}, nil
}

// Buy issues a buy market order given order parameters
// Will return an error if order limit is reached
func (c *Client) Buy(ctx context.Context, order order.Order) (int64, error) {
	// 1.
	// a) Check if order limit has been reached
	// b) Check if client has sufficient funds to make purchase
	// c) Handle context deadlines if request takes longer than 4 secs

	// 2. Make a API call using order as properties
	// a) Make API call with order parameters
	// b) Check if API call was successful
	// c) Handle error if any

	// 3. Save order to cache and database
	// a) check if order is not in cache
	// b) save order to cache
	// c) save order to database or update

	// 4. Return order number
	return 0, nil
}

// Sell issues a sell market order given order parameters
// Uses orderId to determine what order needs to be sold
// If orderId cannot be found in database, it will return an error
func (c *Client) Sell(ctx context.Context, order order.Order) (int64, error) {
	// 1.
	// a) Check if orderId user wants to sell still exists in cache/DB
	// b) Handle context deadlines if request takes longer than 4 secs

	// 2. Make a API call using order as properties
	// a) Make API call with order parameters
	// b) Check if API call was successful
	// c) Handle error if any

	// 3.
	// a) Delete buy order in cache
	// c) save order to database

	// 4. Return order number
	return 0, nil
}

// Cancel will cancel an order to prevent execution. Returns true if order cancellation is successful.
// Will return an error if orderId does not exist.
func (c *Client) Cancel(ctx *context.Context, orderId int64) (bool, error) {
	// 1. Check if order exists
	// 2. invoke rest API cancellation
	// a) return API result
	// b)
	orderIds, err := c.getOrderIds()
	if err != nil {
		return false, err
	}
	orderPos, ok := orderExists(orderId, orderIds)
	if !ok {
		return false, fmt.Errorf("%v: id %d", InvalidOrderErr, orderId)
	}
	// Delete order from client
	c.orders = append(c.orders[:orderPos], c.orders[orderPos+1:]...)

	return true, nil
}

// GetOrderIds returns a slice of int64 which represents orderIds
func (c *Client) getOrderIds() ([]int64, error) {
	if c.orders == nil {
		return nil, EmptyOrdersErr
	}
	ids := make([]int64, 0, len(c.orders))
	for _, v := range c.orders {
		ids = append(ids, v.ID)
	}
	return ids, nil
}

// orderExists returns true if order exists with local client
// Additionaly returns index of order
func orderExists(id int64, orderIds []int64) (int, bool) {
	for i, v := range orderIds {
		if v == id {
			return i, true
		}
	}
	return 0, false
}

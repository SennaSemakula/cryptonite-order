// package order is responsible for providing data structures to define cryptocurrency orders
package order

import (
	"context"
	"time"
)

// Order either represents a buy or sell order
// with required fields to make orders
type Order struct {
	ID        int64
	Type      OrderType
	Symbol    string        `json:"symbol"`
	Amount    float64       `json:"amount"`
	Timestamp time.Duration `json:"time"`
}

// OrderType stores an enum for order types as ints
type OrderType int

const (
	Buy OrderType = iota
	Sell
	Cancel
)

type OrderList []*Order

type Orderer interface {
	Buy(context.Context, *Order) (int64, error)
	Sell(context.Context, *Order) (int64, error)
	Cancel(ctx context.Context, orderId int64) (bool, error)
}

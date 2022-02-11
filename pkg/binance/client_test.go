package binance

import (
	"context"
	"testing"

	// do I need to import this into the testing package as well????
	"cryptonite/pkg/order"
)

func TestCancel(t *testing.T) {
	c := Client{
		orders: []order.Order{
			{ID: 1234},
		},
	}
	ctx := context.TODO()
	input := map[string]struct {
		ctx    context.Context
		order  order.Order
		result bool
	}{
		"valid cancel order": {
			ctx,
			order.Order{
				ID:     1234,
				Type:   2,
				Symbol: "DOGE",
				Amount: 10,
			},
			true,
		},
		"order that doesn't exist": {
			ctx,
			order.Order{
				ID:     0000,
				Type:   2,
				Symbol: "BTC",
				Amount: 1.5,
			},
			false,
		},
	}

	for name, tc := range input {
		t.Run(name, func(t *testing.T) {
			actual, err := c.Cancel(&ctx, tc.order.ID)
			if actual && err != nil {
				t.Fatalf("error: %v", err)
			}
			if actual != tc.result {
				t.Fatalf("error: expected %t but got %v", tc.result, actual)
			}
		})
	}

	t.Run("Delete order from local client orders", func(t *testing.T) {
		orderIds, err := c.getOrderIds()
		if err != nil {
			t.Fatal(err)
		}
		var id int64 = 1234
		if _, ok := orderExists(id, orderIds); ok {
			t.Fatalf("did not expect order_id %d to be in local client", id)
		}
	})

}

func TestOrderExists(t *testing.T) {
	input := map[string]struct {
		id     int64
		orders []int64
		result bool
	}{
		"valid_order": {
			id:     123,
			orders: []int64{123, 0, 12},
			result: true,
		},
		"order_non_existent": {
			id:     158,
			orders: []int64{1, 2, 12},
			result: false,
		},
		"out_of_range_order": {
			id:     -2,
			orders: []int64{1, 2, 12},
			result: false,
		},
	}
	for name, tc := range input {
		t.Run(name, func(t *testing.T) {
			actual, ok := orderExists(tc.id, tc.orders)
			if tc.result != ok {
				t.Fatalf("got %t but expected %t", ok, tc.result)
			}
			if actual < 0 || actual > len(tc.orders) {
				t.Fatalf("expected id %d to have index within order length", tc.id)
			}
		})

	}
}

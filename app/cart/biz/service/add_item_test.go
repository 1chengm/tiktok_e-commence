package service

import (
	"context"
	"testing"
	cart "gomall/rpc_gen/kitex_gen/cart"
)

func TestAddItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	req := &cart.AddItemReq{
		UserId: 0,
		Item: &cart.GetItem{
			ProductId: 0,	
			Quantity: 0,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
	
}

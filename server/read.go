package server

import (
	"github.com/huyhvq/minuet_customer"
	"golang.org/x/net/context"
)

func (s MinuetCustomerServer) Read(ctx context.Context, r *minuet_customer.Request) (*minuet_customer.Response, error) {
	//return nil, errors.New("not yet implemented")
	return &minuet_customer.Response{Id: "huyhuy",}, nil
}

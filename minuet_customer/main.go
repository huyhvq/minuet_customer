package main

import (
	"github.com/huyhvq/minuet_customer"
	"github.com/huyhvq/minuet_customer/minuet_customer/cmd"
	"github.com/huyhvq/minuet_customer/server"
	"github.com/lileio/lile"
	"google.golang.org/grpc"
)

func main() {
	s := &server.MinuetCustomerServer{}

	lile.Name("minuet_customer")
	lile.Server(func(g *grpc.Server) {
		minuet_customer.RegisterMinuetCustomerServer(g, s)
	})

	cmd.Execute()
}

package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/huyhvq/minuet_customer"
	"github.com/lileio/lile"
)

var s = MinuetCustomerServer{}
var cli minuet_customer.MinuetCustomerClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		minuet_customer.RegisterMinuetCustomerServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = minuet_customer.NewMinuetCustomerClient(lile.TestConn(addr))

	os.Exit(m.Run())
}

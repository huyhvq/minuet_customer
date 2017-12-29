package server

import (
	"testing"

	"github.com/huyhvq/minuet_customer"
	"github.com/stretchr/testify/assert"
	context "golang.org/x/net/context"
)

func TestRead(t *testing.T) {
	ctx := context.Background()
	req := &minuet_customer.Request{}

	res, err := cli.Read(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

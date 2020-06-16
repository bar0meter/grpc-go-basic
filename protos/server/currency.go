package server

import (
	"context"
	protos "github.com/go-grpc-basics/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency{
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context, request *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", request.GetBase(), "destination", request.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}

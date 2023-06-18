package infra

import (
	"context"

	"github.com/julioc98/citi/internal/domain"
)

type BacenGateway struct {
}

func NewBacenGateway() *BacenGateway {
	return &BacenGateway{}
}

func (g *BacenGateway) CobPUT(ctx context.Context, txID string, shippingDetail domain.ShippingDetail) error {
	return nil
}

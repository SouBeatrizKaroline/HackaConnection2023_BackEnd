package infra

import (
	"context"

	"github.com/julioc98/citi/internal/domain"
)

type ShippingRepository struct {
}

func NewShippingRepository() *ShippingRepository {
	return &ShippingRepository{}
}

func (r *ShippingRepository) Save(ctx context.Context, txID string, shipping *domain.Shipping) error {
	return nil
}

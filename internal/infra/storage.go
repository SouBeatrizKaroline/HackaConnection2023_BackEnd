package infra

import (
	"context"

	"github.com/julioc98/citi/internal/domain"
)

type ReturnStorage struct {
}

func NewReturnStorage() *ReturnStorage {
	return &ReturnStorage{}
}

func (r *ReturnStorage) Save(ctx context.Context, filename string, ret domain.Return) error {
	return nil
}

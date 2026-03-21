package domain

import (
	"context"

	"github.com/shopspring/decimal"
)

type PriceFromExchangeGetter interface {
	GetPriceFromExchange(ctx context.Context, pairName string, exchangeName string) (decimal.Decimal, error)
}

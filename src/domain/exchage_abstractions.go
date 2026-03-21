package domain

import (
	"context"

	"github.com/shopspring/decimal"
)

type ExchangeClient interface {
	GetName() string
	GetExchangePrice(ctx context.Context, pairName string) (decimal.Decimal, error)
}

type BinanceGetPriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

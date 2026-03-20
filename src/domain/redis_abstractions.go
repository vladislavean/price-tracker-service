package domain

import "github.com/shopspring/decimal"

type RedisGetPriceExchangeClient interface {
	GetPriceExchange(pairName string) (decimal.Decimal, error)
	SetPriceExchange(pairName string, price decimal.Decimal) error
}

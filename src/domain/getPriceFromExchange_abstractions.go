package domain

import "github.com/shopspring/decimal"

type GetPriceFromExchange interface {
	GetPriceFromExchange(painName string, exchangeName string) (decimal.Decimal, error)
}

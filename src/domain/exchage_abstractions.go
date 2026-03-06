package domain

import "github.com/shopspring/decimal"

type ExchangeClient interface {
	GetName() string
	GetExchangePrice(pairName string) (decimal.Decimal, error)
}

type BinanceGetPriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

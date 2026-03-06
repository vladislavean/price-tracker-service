package exchange_clients

import "github.com/shopspring/decimal"

type ByBitGetPriceRequest struct {
}

func NewByBitGetPriceRequest() *ByBitGetPriceRequest {
	return &ByBitGetPriceRequest{}
}

func (r *ByBitGetPriceRequest) GetExchangePrice(pairName string) (decimal.Decimal, error) {
	panic("implement me")
}

func (r *ByBitGetPriceRequest) GetName() string {
	return "bybit"
}

package exchange_clients

import "github.com/shopspring/decimal"

type ByBitGetPriceClientImpl struct {
}

func NewByBitGetPriceClientImpl() *ByBitGetPriceClientImpl {
	return &ByBitGetPriceClientImpl{}
}

func (r *ByBitGetPriceClientImpl) GetExchangePrice(pairName string) (decimal.Decimal, error) {
	panic("implement me")
}

func (r *ByBitGetPriceClientImpl) GetName() string {
	return "bybit"
}

package exchange_clients

import "github.com/shopspring/decimal"

type OkxGetPriceClientImpl struct{}

func NewOkxGetPriceClientImpl() *OkxGetPriceClientImpl {
	return &OkxGetPriceClientImpl{}
}

func (r *OkxGetPriceClientImpl) GetExchangePrice(pairName string) (decimal.Decimal, error) {
	panic("implement me")
}

func (r *OkxGetPriceClientImpl) GetName() string {
	return "okx"
}

package exchange_clients

import "github.com/shopspring/decimal"

type OkxGetPriceRequest struct{}

func NewOkxGetPriceRequest() *OkxGetPriceRequest {
	return &OkxGetPriceRequest{}
}

func (r *OkxGetPriceRequest) GetExchangePrice(pairName string) (decimal.Decimal, error) {
	panic("implement me")
}

func (r *OkxGetPriceRequest) GetName() string {
	return "okx"
}

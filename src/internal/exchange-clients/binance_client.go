package exchange_clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"price-tracker-service/src/domain"

	"github.com/shopspring/decimal"
)

type BinanceGetPriceClientImpl struct {
	config *domain.ExchangeClientConfig
}

func NewBinanceGetPriceClientImpl(config *domain.ExchangeClientConfig) *BinanceGetPriceClientImpl {
	return &BinanceGetPriceClientImpl{config: config}
}

func (r *BinanceGetPriceClientImpl) GetExchangePrice(pairName string) (decimal.Decimal, error) {
	baseUrl := r.config.BinanceBaseUrl + fmt.Sprintf("ticker/price?symbol=%s", pairName)

	resp, err := http.Get(baseUrl)
	if err != nil {
		return decimal.Zero, err
	}
	defer resp.Body.Close()

	var data domain.BinanceGetPriceResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return decimal.Zero, err
	}

	val, err := decimal.NewFromString(data.Price)
	if err != nil {
		return decimal.Zero, err
	}

	return val, nil
}

func (r *BinanceGetPriceClientImpl) GetName() string {
	return "binance"
}

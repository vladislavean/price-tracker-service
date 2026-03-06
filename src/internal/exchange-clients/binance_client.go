package exchange_clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"price-tracker-service/src/domain"

	"github.com/shopspring/decimal"
)

type BinanceGetPriceRequest struct {
}

func NewBinanceGetPriceRequest() *BinanceGetPriceRequest {
	return &BinanceGetPriceRequest{}
}

func (r *BinanceGetPriceRequest) GetExchangePrice(pairName string) (decimal.Decimal, error) {
	baseUrl := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", pairName)

	resp, err := http.Get(baseUrl)
	if err != nil {
		return decimal.NewFromInt(0), err
	}
	defer resp.Body.Close()

	var data domain.BinanceGetPriceResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	val, err := decimal.NewFromString(data.Price)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	return val, nil
}

func (r *BinanceGetPriceRequest) GetName() string {
	return "binance"
}

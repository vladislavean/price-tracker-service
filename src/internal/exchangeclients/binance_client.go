package exchangeclients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"price-tracker-service/src/domain"
	"time"

	"github.com/shopspring/decimal"
)

type BinanceGetPriceClientImpl struct {
	config *domain.ExchangeClientConfig
}

func NewBinanceGetPriceClientImpl(config *domain.ExchangeClientConfig) *BinanceGetPriceClientImpl {
	return &BinanceGetPriceClientImpl{config: config}
}

func (r *BinanceGetPriceClientImpl) GetExchangePrice(ctx context.Context, pairName string) (decimal.Decimal, error) {
	baseUrl := r.config.BinanceBaseUrl + fmt.Sprintf("ticker/price?symbol=%s", pairName)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, baseUrl, nil)
	resp, err := httpClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		return decimal.Zero, fmt.Errorf("binance: unexpected status %d for pair %s", resp.StatusCode, pairName)
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

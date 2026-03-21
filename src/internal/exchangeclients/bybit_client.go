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

type ByBitGetPriceResultList struct {
	List []ByBitGetPriceResult `json:"list"`
}

type ByBitGetPriceResult struct {
	LastPrice string `json:"lastPrice"`
}

type ByBitGetPriceResponse struct {
	RetCode int                     `json:"retCode"`
	RetMsg  string                  `json:"retMsg"`
	Result  ByBitGetPriceResultList `json:"result"`
}

type ByBitGetPriceClientImpl struct {
	config *domain.ExchangeClientConfig
}

func NewByBitGetPriceClientImpl(config *domain.ExchangeClientConfig) *ByBitGetPriceClientImpl {
	return &ByBitGetPriceClientImpl{config: config}
}

func (r *ByBitGetPriceClientImpl) GetExchangePrice(ctx context.Context, pairName string) (decimal.Decimal, error) {
	baseUrl := r.config.ByBitBaseUrl + fmt.Sprintf("market/tickers?category=spot&symbol=%s", pairName)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, baseUrl, nil)
	resp, err := httpClient.Do(req)

	if err != nil {
		return decimal.Zero, err
	}
	defer resp.Body.Close()

	var result ByBitGetPriceResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return decimal.Zero, err
	}
	if len(result.Result.List) == 0 {
		return decimal.Zero, fmt.Errorf("bybit: empty response for pair %s", pairName)
	}

	return decimal.NewFromString(result.Result.List[0].LastPrice)
}

func (r *ByBitGetPriceClientImpl) GetName() string {
	return "bybit"
}

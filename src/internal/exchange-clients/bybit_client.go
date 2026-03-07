package exchange_clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"price-tracker-service/src/domain"

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

func (r *ByBitGetPriceClientImpl) GetExchangePrice(pairName string) (decimal.Decimal, error) {
	baseUrl := r.config.ByBitBaseUrl + fmt.Sprintf("market/tickers?category=spot&symbol=%s", pairName)

	resp, err := http.Get(baseUrl)
	if err != nil {
		return decimal.Zero, err
	}
	defer resp.Body.Close()

	var result ByBitGetPriceResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return decimal.Zero, err
	}

	return decimal.NewFromString(result.Result.List[0].LastPrice)
}

func (r *ByBitGetPriceClientImpl) GetName() string {
	return "bybit"
}

package exchange_clients

import (
	"encoding/json"
	"fmt"
	"net/http"

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
}

func NewByBitGetPriceClientImpl() *ByBitGetPriceClientImpl {
	return &ByBitGetPriceClientImpl{}
}

func (r *ByBitGetPriceClientImpl) GetExchangePrice(pairName string) (decimal.Decimal, error) {
	baseUrl := fmt.Sprintf("https://api-testnet.bybit.com/v5/market/tickers?category=spot&symbol=%s", pairName)

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

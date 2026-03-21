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

type OkxGetPriceClientPriceResponse struct {
	Last string `json:"last"`
}

type OkxGetPriceClientResponse struct {
	Code string                           `json:"code"`
	Msg  string                           `json:"msg"`
	Data []OkxGetPriceClientPriceResponse `json:"data"`
}

type OkxGetPriceClientImpl struct {
	config *domain.ExchangeClientConfig
}

func NewOkxGetPriceClientImpl(config *domain.ExchangeClientConfig) *OkxGetPriceClientImpl {
	return &OkxGetPriceClientImpl{config: config}
}

func (r *OkxGetPriceClientImpl) GetExchangePrice(ctx context.Context, pairName string) (decimal.Decimal, error) {
	baseUrl := r.config.OkxBaseUrl + fmt.Sprintf("market/ticker?instId=%s", pairName)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, baseUrl, nil)
	resp, err := httpClient.Do(req)

	if err != nil {
		return decimal.Zero, err
	}
	defer resp.Body.Close()

	var data OkxGetPriceClientResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return decimal.Zero, err
	}
	if len(data.Data) == 0 {
		return decimal.Zero, fmt.Errorf("okx: empty response for pair %s", pairName)
	}

	return decimal.NewFromString(data.Data[0].Last)
}

func (r *OkxGetPriceClientImpl) GetName() string {
	return "okx"
}

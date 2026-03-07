package exchange_clients

import (
	"encoding/json"
	"fmt"
	"net/http"

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

type OkxGetPriceClientImpl struct{}

func NewOkxGetPriceClientImpl() *OkxGetPriceClientImpl {
	return &OkxGetPriceClientImpl{}
}

func (r *OkxGetPriceClientImpl) GetExchangePrice(pairName string) (decimal.Decimal, error) {
	baseUrl := fmt.Sprintf("https://www.okx.com/api/v5/market/ticker?instId=%s", pairName)

	resp, err := http.Get(baseUrl)
	if err != nil {
		return decimal.Zero, err
	}

	defer resp.Body.Close()

	var data OkxGetPriceClientResponse
	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return decimal.Zero, err
	}

	return decimal.NewFromString(data.Data[0].Last)
}

func (r *OkxGetPriceClientImpl) GetName() string {
	return "okx"
}

package usecases

import (
	"fmt"
	"price-tracker-service/src/domain"

	"github.com/shopspring/decimal"
)

type GetPriceFromExchangeImpl struct {
	clients []domain.ExchangeClient
}

func NewGetPriceFromExchangeImpl(clients []domain.ExchangeClient) *GetPriceFromExchangeImpl {
	return &GetPriceFromExchangeImpl{clients}
}

func (impl *GetPriceFromExchangeImpl) GetPriceFromExchange(pairName string, exchangeName string) (decimal.Decimal, error) {
	for _, client := range impl.clients {
		if client.GetName() == exchangeName {
			val, err := client.GetExchangePrice(pairName)
			if err != nil {
				return decimal.Zero, err
			}

			return val, nil
		}
	}
	return decimal.Zero, fmt.Errorf("exchange %s does not exist", exchangeName)
}

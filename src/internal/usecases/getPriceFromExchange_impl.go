package usecases

import (
	"fmt"
	"log"
	"price-tracker-service/src/domain"

	"github.com/shopspring/decimal"
)

type GetPriceFromExchangeUsecasesImpl struct {
	clients     []domain.ExchangeClient
	redisClient domain.RedisGetPriceExchangeClient
}

func NewGetPriceFromExchangeUsecasesImpl(
	clients []domain.ExchangeClient,
	redisClient domain.RedisGetPriceExchangeClient) *GetPriceFromExchangeUsecasesImpl {
	return &GetPriceFromExchangeUsecasesImpl{clients, redisClient}
}

func (impl *GetPriceFromExchangeUsecasesImpl) GetPriceFromExchange(pairName string, exchangeName string) (decimal.Decimal, error) {
	priceFromRedis, err := impl.redisClient.GetPriceExchange(pairName)
	if err != nil && priceFromRedis.IsZero() {
		log.Println(err)
	}
	if !priceFromRedis.IsZero() {
		return priceFromRedis, nil
	}

	for _, client := range impl.clients {
		if client.GetName() == exchangeName {
			val, err := client.GetExchangePrice(pairName)
			if err != nil {
				return decimal.Zero, err
			}

			err = impl.redisClient.SetPriceExchange(pairName, val)
			if err != nil {
				log.Println(err)
			}

			return val, nil
		}
	}
	return decimal.Zero, fmt.Errorf("exchange %s does not exist", exchangeName)
}

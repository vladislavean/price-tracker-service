package usecases

import (
	"context"
	"fmt"
	"log/slog"
	"price-tracker-service/src/domain"

	"github.com/shopspring/decimal"
)

type GetPriceFromExchangeUsecasesImpl struct {
	clients     map[string]domain.ExchangeClient
	redisClient domain.RedisGetPriceExchangeClient
}

func NewGetPriceFromExchangeUsecasesImpl(
	clients map[string]domain.ExchangeClient,
	redisClient domain.RedisGetPriceExchangeClient) *GetPriceFromExchangeUsecasesImpl {
	return &GetPriceFromExchangeUsecasesImpl{clients, redisClient}
}

func (impl *GetPriceFromExchangeUsecasesImpl) GetPriceFromExchange(ctx context.Context, pairName string, exchangeName string) (decimal.Decimal, error) {
	priceFromRedis, err := impl.redisClient.GetPriceExchange(pairName)
	if err != nil && priceFromRedis.IsZero() {
		slog.Warn("redis get failed, falling back to exchange", "pair", pairName, "error", err)
	}
	if !priceFromRedis.IsZero() {
		return priceFromRedis, nil
	}

	var client = impl.clients[exchangeName]
	if client == nil {
		return decimal.Zero, fmt.Errorf("exchange %s does not exist", exchangeName)
	}

	val, err := client.GetExchangePrice(ctx, pairName)
	if err != nil {
		return decimal.Zero, err
	}

	err = impl.redisClient.SetPriceExchange(pairName, val)
	if err != nil {
		slog.Warn("redis get failed, falling back to exchange", "pair", pairName, "error", err)
	}

	return val, nil
}

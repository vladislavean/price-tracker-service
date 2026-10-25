package usecases

import (
	"context"
	"fmt"
	"log/slog"
	"price-tracker-service/src/domain"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type GetPriceFromExchangeUsecasesImpl struct {
	clients     map[string]domain.ExchangeClient
	redisClient domain.RedisGetPriceExchangeClient
	repository  domain.LogRepository
}

func NewGetPriceFromExchangeUsecasesImpl(
	clients map[string]domain.ExchangeClient,
	redisClient domain.RedisGetPriceExchangeClient,
	repository domain.LogRepository) *GetPriceFromExchangeUsecasesImpl {
	return &GetPriceFromExchangeUsecasesImpl{clients, redisClient, repository}
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

	// logs
	f, _ := strconv.ParseFloat(val.String(), 64)
	if err := impl.repository.InsertLog(ctx, &domain.PriceChangesLog{
		Client:     exchangeName,
		PairName:   pairName,
		Value:      f,
		CreateDate: time.Now(),
	}); err != nil {
		slog.Warn("error write log", "err", err)
	}

	return val, nil
}

package redisintegration

import (
	"context"
	"fmt"
	"log/slog"
	"price-tracker-service/src/domain"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type RedisPriceExchangeClientImpl struct {
	client *redis.Client
}

func NewRedisPriceExchangeClientImpl(config *domain.RedisClientConfig) *RedisPriceExchangeClientImpl {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	return &RedisPriceExchangeClientImpl{client: client}
}

func (r *RedisPriceExchangeClientImpl) GetPriceExchange(pairName string) (decimal.Decimal, error) {
	conn := r.client

	if conn == nil {
		err := fmt.Errorf("cannot get price redis connection")
		return decimal.Zero, err
	}
	defer func(conn *redis.Client) {
		err := conn.Close()
		if err != nil {
			slog.Error("error closing redis connection", "error", err)
		}
	}(conn)

	ctx := context.Background()
	result := conn.Get(ctx, pairName)

	decimalString, err := result.Result()
	if err != nil {
		return decimal.Zero, err
	}
	dcml, err := decimal.NewFromString(decimalString)
	if err != nil {
		return decimal.Zero, err
	}
	return dcml, nil
}

func (r *RedisPriceExchangeClientImpl) SetPriceExchange(pairName string, price decimal.Decimal) error {
	conn := r.client

	if conn == nil {
		err := fmt.Errorf("cannot get price redis connection")
		return err
	}

	defer func(conn *redis.Client) {
		err := conn.Close()
		if err != nil {
			slog.Error("error closing redis connection", "error", err)
		}
	}(conn)

	ctx := context.Background()

	conn.Set(ctx, pairName, price.String(), time.Minute)
	return nil
}

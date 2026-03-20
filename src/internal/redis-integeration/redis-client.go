package redis_integeration

import (
	"context"
	"fmt"
	"log"
	"price-tracker-service/src/domain"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type RedisPriceExchangeClientImpl struct {
	config *domain.RedisClientConfig
}

func NewRedisPriceExchangeClientImpl(config *domain.RedisClientConfig) *RedisPriceExchangeClientImpl {
	return &RedisPriceExchangeClientImpl{config: config}
}

func (r *RedisPriceExchangeClientImpl) GetPriceExchange(pairName string) (decimal.Decimal, error) {
	conn := r.getConnection()
	defer func(conn *redis.Client) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing connection to redis: %v", err)
		}
	}(conn)
	if conn == nil {
		err := fmt.Errorf("cannot get price redis connection")
		return decimal.Zero, err
	}

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
	conn := r.getConnection()
	defer func(conn *redis.Client) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing connection to redis: %v", err)
		}
	}(conn)
	if conn == nil {
		err := fmt.Errorf("cannot get price redis connection")
		return err
	}
	ctx := context.Background()

	conn.Set(ctx, pairName, price.String(), time.Minute)
	return nil
}

func (r *RedisPriceExchangeClientImpl) getConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.config.Addr,
		Password: "",
		DB:       r.config.DB,
		Protocol: 2,
	})

	return rdb
}

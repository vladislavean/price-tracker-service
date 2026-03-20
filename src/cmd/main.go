package main

import (
	"fmt"
	"log"
	"price-tracker-service/src/config"
	"price-tracker-service/src/domain"
	exchangeclients "price-tracker-service/src/internal/exchange-clients"
	"price-tracker-service/src/internal/handlers"
	redisintegeration "price-tracker-service/src/internal/redis-integeration"
	"price-tracker-service/src/internal/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	databaseConfig, err := config.LoadDatabaseConfig()
	if err != nil {
		panic(err)
	}

	exchangeConfig, err := config.LoadExchangeConfig()
	if err != nil {
		panic(err)
	}

	redisConfig, err := config.LoadRedisConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(databaseConfig.Database)

	binanceClient := exchangeclients.NewBinanceGetPriceClientImpl(exchangeConfig)
	bybitClient := exchangeclients.NewByBitGetPriceClientImpl(exchangeConfig)
	okxClient := exchangeclients.NewOkxGetPriceClientImpl(exchangeConfig)
	clients := []domain.ExchangeClient{binanceClient, okxClient, bybitClient}

	redisClient := redisintegeration.NewRedisPriceExchangeClientImpl(redisConfig)

	coreImpl := usecases.NewGetPriceFromExchangeUsecasesImpl(clients, redisClient)

	handler := handlers.NewGetPriceHandler(coreImpl)

	api := NewControllers(handler)
	api.RegisterControllers(g)

	go func() {
		err := g.Run(":4400")
		if err != nil {
			log.Fatal(err)
		}
	}()
}

package main

import (
	"fmt"
	"price-tracker-service/src/config"
	"price-tracker-service/src/domain"
	exchangeclients "price-tracker-service/src/internal/exchangeclients"
	"price-tracker-service/src/internal/handlers"
	redisintegeration "price-tracker-service/src/internal/redisintegration"
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

	redisClient := redisintegeration.NewRedisPriceExchangeClientImpl(redisConfig)

	binanceClient := exchangeclients.NewBinanceGetPriceClientImpl(exchangeConfig)
	bybitClient := exchangeclients.NewByBitGetPriceClientImpl(exchangeConfig)
	okxClient := exchangeclients.NewOkxGetPriceClientImpl(exchangeConfig)
	clients := []domain.ExchangeClient{binanceClient, okxClient, bybitClient}

	clientsMap := make(map[string]domain.ExchangeClient, len(clients))
	for _, client := range clients {
		clientsMap[client.GetName()] = client
	}

	coreImpl := usecases.NewGetPriceFromExchangeUsecasesImpl(clientsMap, redisClient)

	handler := handlers.NewGetPriceHandler(coreImpl)

	api := NewControllers(handler)
	api.RegisterControllers(g)

	g.Run(":4400")
}

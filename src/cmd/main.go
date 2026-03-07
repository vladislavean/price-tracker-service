package main

import (
	"fmt"
	"price-tracker-service/src/config"
	"price-tracker-service/src/domain"
	exchange_clients "price-tracker-service/src/internal/exchange-clients"
	"price-tracker-service/src/internal/handlers"
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

	fmt.Println(databaseConfig.Database)

	binanceClient := exchange_clients.NewBinanceGetPriceClientImpl(exchangeConfig)
	bybitClient := exchange_clients.NewByBitGetPriceClientImpl(exchangeConfig)
	okxClient := exchange_clients.NewOkxGetPriceClientImpl(exchangeConfig)
	clients := []domain.ExchangeClient{binanceClient, okxClient, bybitClient}

	coreImpl := usecases.NewGetPriceFromExchangeUsecasesImpl(clients)

	handler := handlers.NewGetPriceHandler(coreImpl)

	api := NewControllers(handler)
	api.RegisterControllers(g)

	g.Run(":4400")
}

package main

import (
	"price-tracker-service/src/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Api struct {
	getPrice *handlers.BinanceGetPriceHandler
}

func NewControllers(getPrice *handlers.BinanceGetPriceHandler) *Api {
	return &Api{getPrice}
}

func (a *Api) RegisterControllers(r *gin.Engine) {
	r.GET("/get-price", a.getPrice.Handle)
}

package main

import (
	"price-tracker-service/src/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Api struct {
	getPrice *handlers.BinanceGetPriceHandler
}

func NewApi(getPrice *handlers.BinanceGetPriceHandler) *Api {
	return &Api{getPrice}
}

func (a *Api) RegisterApi(r *gin.Engine) {
	r.GET("/get-price", a.getPrice.Handle)
}

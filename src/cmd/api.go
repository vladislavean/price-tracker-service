package main

import (
	"net/http"
	"price-tracker-service/src/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Api struct {
	getPrice *handlers.GetPriceHandler
}

func NewControllers(getPrice *handlers.GetPriceHandler) *Api {
	return &Api{getPrice}
}

func (a *Api) RegisterControllers(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/get-price", a.getPrice.Handle)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}

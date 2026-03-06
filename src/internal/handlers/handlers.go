package handlers

import (
	"net/http"
	"price-tracker-service/src/internal/usecases"

	"github.com/gin-gonic/gin"
)

type BinanceGetPriceRequestDTO struct {
	PairName     string `json:"pairName"`
	ExchangeName string `json:"exchangeName"`
}

type BinanceGetPriceHandler struct {
	u *usecases.GetPriceFromExchangeImpl
}

func NewGetPriceHandler(getPrice *usecases.GetPriceFromExchangeImpl) *BinanceGetPriceHandler {
	return &BinanceGetPriceHandler{getPrice}
}

func (u *BinanceGetPriceHandler) Handle(c *gin.Context) {
	var dto BinanceGetPriceRequestDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	price, err := u.u.GetPriceFromExchange(dto.PairName, dto.ExchangeName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"price": price})
}

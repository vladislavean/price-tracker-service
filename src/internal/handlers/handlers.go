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
	usecases *usecases.GetPriceFromExchangeUsecasesImpl
}

func NewGetPriceHandler(getPrice *usecases.GetPriceFromExchangeUsecasesImpl) *BinanceGetPriceHandler {
	return &BinanceGetPriceHandler{getPrice}
}

func (h *BinanceGetPriceHandler) Handle(c *gin.Context) {
	var dto BinanceGetPriceRequestDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	price, err := h.usecases.GetPriceFromExchange(dto.PairName, dto.ExchangeName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"price": price})
}

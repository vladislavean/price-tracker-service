package handlers

import (
	"net/http"
	"price-tracker-service/src/domain"

	"github.com/gin-gonic/gin"
)

type GetPriceRequest struct {
	PairName     string `json:"pairName"`
	ExchangeName string `json:"exchangeName"`
}

type GetPriceHandler struct {
	usecases domain.PriceFromExchangeGetter
}

func NewGetPriceHandler(getPrice domain.PriceFromExchangeGetter) *GetPriceHandler {
	return &GetPriceHandler{getPrice}
}

func (h *GetPriceHandler) Handle(c *gin.Context) {
	var req GetPriceRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	price, err := h.usecases.GetPriceFromExchange(c.Request.Context(), req.PairName, req.ExchangeName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"price": price})
}

package domain

import (
	"time"
)

type PriceChangesLog struct {
	Client     string    `mapstructure:"client"`
	PairName   string    `mapstructure:"pairname"`
	Value      float64   `mapstructure:"price_value"`
	CreateDate time.Time `mapstructure:"create_date"`
}

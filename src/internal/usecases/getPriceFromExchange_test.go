package usecases

import (
	"context"
	"errors"
	"price-tracker-service/src/domain"
	"price-tracker-service/src/internal/redisintegration"
	"testing"

	"github.com/shopspring/decimal"
)

type mockExchangeClient struct {
	name        string
	priceByPair map[string]decimal.Decimal
	err         error
}

func (m *mockExchangeClient) GetName() string {
	return m.name
}

func (m *mockExchangeClient) GetExchangePrice(ctx context.Context, pair string) (decimal.Decimal, error) {
	if m.err != nil {
		return decimal.Zero, m.err
	}
	if v, ok := m.priceByPair[pair]; ok {
		return v, nil
	}
	return decimal.Zero, errors.New("pair not found")
}

func Test_GetPriceFromExchangeUsecasesImpl_GetPriceFromExchange(t *testing.T) {
	ethPrice := decimal.NewFromFloat(1234.56)
	btcPrice := decimal.NewFromFloat(77777.77)

	clients := []domain.ExchangeClient{
		&mockExchangeClient{
			name: "binance",
			priceByPair: map[string]decimal.Decimal{
				"ETHUSDT": ethPrice,
			},
		},
		&mockExchangeClient{
			name: "bybit",
			priceByPair: map[string]decimal.Decimal{
				"BTCUSDT": btcPrice,
			},
		},
	}
	var clientsMap = make(map[string]domain.ExchangeClient, len(clients))
	for _, client := range clients {
		clientsMap[client.GetName()] = client
	}

	config := &domain.RedisClientConfig{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	impl := &GetPriceFromExchangeUsecasesImpl{
		clients:     clientsMap,
		redisClient: redisintegration.NewRedisPriceExchangeClientImpl(config),
	}
	tests := []struct {
		name        string
		pair        string
		exchange    string
		setup       func()
		wantPrice   decimal.Decimal
		wantErr     bool
		wantErrText string
	}{
		{
			name:      "success_binance_eth",
			pair:      "ETHUSDT",
			exchange:  "binance",
			setup:     func() {},
			wantPrice: ethPrice,
			wantErr:   false,
		},
		{
			name:      "success_bybit_btc",
			pair:      "BTCUSDT",
			exchange:  "bybit",
			setup:     func() {},
			wantPrice: btcPrice,
			wantErr:   false,
		},
		{
			name:      "exchange_not_found",
			pair:      "ETHUSDT",
			exchange:  "kraken",
			setup:     func() {},
			wantPrice: decimal.Zero,
			wantErr:   true,
		},
		{
			name:     "client_returns_error",
			pair:     "ETHUSDT",
			exchange: "binance",
			setup: func() {
				// портим первого клиента
				m := clients[0].(*mockExchangeClient)
				m.err = errors.New("upstream error")
			},
			wantPrice: decimal.Zero,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := clients[0].(*mockExchangeClient)
			m.err = nil

			tt.setup()
			ctx := context.Background()
			got, err := impl.GetPriceFromExchange(ctx, tt.pair, tt.exchange)

			if tt.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tt.wantErr && tt.wantErrText != "" {
				t.Fatalf("error %q does not contain %q", err.Error(), tt.wantErrText)
			}

			if !got.Equal(tt.wantPrice) {
				t.Fatalf("unexpected price: got %s, want %s", got.String(), tt.wantPrice.String())
			}
		})
	}
}

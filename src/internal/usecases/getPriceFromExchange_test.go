package usecases

import (
	"errors"
	"price-tracker-service/src/domain"
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

func (m *mockExchangeClient) GetExchangePrice(pair string) (decimal.Decimal, error) {
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
	impl := &GetPriceFromExchangeUsecasesImpl{
		clients: clients,
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

			got, err := impl.GetPriceFromExchange(tt.pair, tt.exchange)

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

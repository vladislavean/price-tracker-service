package domain

type DatabaseConfig struct {
	Host     string `mapstructure:"database_host"`
	User     string `mapstructure:"database_user"`
	Password string `mapstructure:"database_password"`
	Database string `mapstructure:"database_name"`
}

type ExchangeClientConfig struct {
	BinanceBaseUrl string `mapstructure:"binance_base_url"`
	ByBitBaseUrl   string `mapstructure:"bybit_base_url"`
	OkxBaseUrl     string `mapstructure:"okx_base_url"`
}

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

type RedisClientConfig struct {
	Addr     string `mapstructure:"redis_addr"`
	Password string `mapstructure:"redis_password"`
	DB       int    `mapstructure:"redis_db"`
}

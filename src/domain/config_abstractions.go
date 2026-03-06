package domain

type DatabaseConfig struct {
	Host     string `mapstructure:"database_host"`
	User     string `mapstructure:"database_user"`
	Password string `mapstructure:"database_password"`
	Database string `mapstructure:"database_name"`
}

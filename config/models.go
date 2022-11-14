package config

type ViperConfig struct {
	Db  DatabaseConfig
	Api ApiConfig
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	Port     string `mapstructure:"DB_PORT"`
}

type ApiConfig struct {
	Port string `mapstructure:"PORT"`
}

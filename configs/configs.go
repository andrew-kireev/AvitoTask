package configs

type Config struct {
	StatisticServerAddr string `toml:"statistic_server_addr"`
	LogLevel            string `toml:"log_level"`
	StatisticPostgresBD string `toml:"statistic_bd"`
}

func NewConfig() *Config {
	return &Config{
		StatisticServerAddr: ":8080",
		LogLevel:            "DEBUG",
		StatisticPostgresBD: "",
	}
}

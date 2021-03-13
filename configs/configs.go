package configs

type Config struct {
	StatisticServerAddr string `toml:"music_server_addr"`
	LogLevel            string `toml:"log_level"`
	MusicPostgresBD     string `toml:"music_bd"`
}

func NewConfig() *Config {
	return &Config{
		StatisticServerAddr: ":8080",
		LogLevel:        "DEBUG",
		MusicPostgresBD: "host=localhost port=5432 dbname=music_service sslmode=disable",
	}
}

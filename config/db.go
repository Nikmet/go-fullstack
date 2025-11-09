package config

type DBConfig struct {
	Url string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Url: getString("DSN", ""),
	}
}

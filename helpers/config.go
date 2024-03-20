package helpers

import "github.com/matias-inc/goenv"

type _Config struct {
	AppAddr string
	DbDSN   string
}

var config *_Config = nil

func GetConfig() *_Config {
	if config == nil {
		config = &_Config{
			AppAddr: goenv.Config("APP_ADDR", "127.0.0.1:8000", goenv.CastString),
			DbDSN:   goenv.Config("DB_DSN", "jhon:secr3t@tcp(localhost)/daddy", goenv.CastString),
		}
	}

	return config
}

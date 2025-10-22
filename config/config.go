package config

var Values *Config

type Config struct {
	Http HttpConfig `config:"http"`
}

type HttpConfig struct {
	Port int `config:"port"`
}

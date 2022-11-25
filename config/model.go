package config

type OAuthConfig struct {
	ApiKey    string `yaml:"api_key"`
	ApiSecret string `yaml:"api_secret"`
}

type TwitterConfig struct {
	Enabled        bool        `yaml:"enabled"`
	Authentication OAuthConfig `yaml:"authentication"`
}

type ProviderConfig struct {
	Twitter TwitterConfig `yaml:"twitter"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type Config struct {
	Providers ProviderConfig `yaml:"providers"`
	Server    ServerConfig   `yaml:"server"`
}

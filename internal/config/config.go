package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	defaultHttpPort               = "8000"
	defaultHttpRWTimeout          = 10 * time.Second
	defaultHttpMaxHeaderMegabytes = 1

	EnvLocal = "local"
	Prod     = "prod"
)

type (
	Config struct {
		Environment string
		HTTP        HTTPConfig
	}
	HTTPConfig struct {
		Host               string        `mapStructure:"host"`
		Port               string        `mapStructure:"port"`
		ReadTimeout        time.Duration `mapStructure:"readTimeout"`
		WriteTimeout       time.Duration `mapStructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapStructure:"maxHeaderBytes"`
	}
)

func Init(configDir string) (*Config, error) {
	populateDefaults()

	if err := parseEnv(); err != nil {
		return nil, err
	}

	if err := parseConfigFile(configDir, viper.GetString("env")); err != nil {
		return nil, err
	}

	var cfg Config

	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil

}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	return nil
}

func parseConfigFile(folder, env string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == EnvLocal {
		return nil
	}

	viper.SetConfigName(env)

	return viper.MergeInConfig()
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = viper.GetString("host")

	cfg.Environment = viper.GetString("env")
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.max_header_megabytes", defaultHttpMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHttpRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHttpRWTimeout)
}

func parseEnv() error {
	if err := parseHostFromEnv(); err != nil {
		return err
	}

	return parsePasswordFromEnv()
}

func parseHostFromEnv() error {
	viper.SetEnvPrefix("http")

	return viper.BindEnv("host")
}

func parsePasswordFromEnv() error {
	viper.SetEnvPrefix("password")

	return viper.BindEnv("salt")
}

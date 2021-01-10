package configuration

import (
	"io"

	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		ListenPort   int    `envconfig:"APP_PORT" default:"3000"`
		DatabaseUser string `envconfig:"DB_USER"`
		DatabasePass string `envconfig:"DB_PASS"`
		DatabaseHost string `envconfig:"DB_HOST"`
		DatabasePort int    `envconfig:"DB_PORT"`
		DatabaseName string `envconfig:"DB_NAME"`
	}
)

var (
	globalConfig Config
)

func Usage(output io.Writer) {
	if err := envconfig.Usagef("", &globalConfig, output, envconfig.DefaultTableFormat); err != nil {
		panic(err.Error())
	}
}

func Load() {
	envconfig.MustProcess("", &globalConfig)
}

func Get() Config {
	return globalConfig
}

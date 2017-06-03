package config

import(
	"github.com/kelseyhightower/envconfig"
)

var DB DBConfig

type DBConfig struct {
	Host     string `required:"true"`
	Port	 string `default:"3306"`
	Name     string `required:"true"`
	User     string `required:"true"`
	Pass string `required:"true"`
	ListenPort string `required:"true"`
}

func LoadFromEnv() error {
	if err := envconfig.Process("db", &DB); err != nil {
		return err
	}

	return nil
}
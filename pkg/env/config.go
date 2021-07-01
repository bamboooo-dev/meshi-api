package env

import (
	"github.com/bamboooo-dev/meshi-api/app/interface/mysql"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MeshiMySQL mysql.Config
}

func LoadConfig() (*Config, error) {
	cfg := new(Config)

	var m mysql.Config
	err := envconfig.Process("meshi_db", &m)
	if err != nil {
		return nil, err
	}

	cfg.MeshiMySQL = m

	return cfg, nil
}

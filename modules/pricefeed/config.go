package pricefeed

import (
	"gopkg.in/yaml.v3"

	"github.com/forbole/bdjuno/types"
)

// Config contains the configuration about the pricefeed module
type Config struct {
	Tokens []types.Token `toml:"tokens"`
}

func ParseConfig(bz []byte) (*Config, error) {
	type T struct {
		Config *Config `yaml:"pricefeed"`
	}
	var cfg T
	err := yaml.Unmarshal(bz, &cfg)
	return cfg.Config, err
}
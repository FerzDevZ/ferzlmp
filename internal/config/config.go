package config

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
)

type Config struct {
	ApachePath string `mapstructure:"apache_path" yaml:"apache_path"`
	MySQLPath  string `mapstructure:"mysql_path" yaml:"mysql_path"`
	PHPPath    string `mapstructure:"php_path" yaml:"php_path"`
	ProjectRoot string `mapstructure:"project_root" yaml:"project_root"`
	PortApache int    `mapstructure:"port_apache" yaml:"port_apache"`
	PortMySQL  int    `mapstructure:"port_mysql" yaml:"port_mysql"`
	// Per-project version support
	ProjectVersions map[string]struct {
		PHPPath   string `yaml:"php_path"`
		MySQLPath string `yaml:"mysql_path"`
	} `yaml:"project_versions"`
}

func LoadConfig(configFile string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return &c, nil
}

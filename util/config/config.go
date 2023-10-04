package config

import (
	"os"
	"strings"

	toml "github.com/pelletier/go-toml"
	"github.com/spf13/cast"
)

type Config struct {
	configFilePath string
	configTree     *toml.Tree
	envKeyReplacer *strings.Replacer
}

var c *Config

func New() *Config {
	c := new(Config)
	c.configFilePath = "./config.toml"
	return c
}

func init() {
	c = New()
}

func SetConfigPath(configFilePath string) {
	c.SetConfigPath(configFilePath)
}

func (c *Config) SetConfigPath(configFilePath string) {
	if configFilePath != "" {
		c.configFilePath = configFilePath
	}
}

func SetEnvKeyReplacer(r *strings.Replacer) {
	c.SetEnvKeyReplacer(r)
}

func (c *Config) SetEnvKeyReplacer(r *strings.Replacer) {
	c.envKeyReplacer = r
}

func ReadInConfig() error {
	return c.ReadInConfig()
}

func (c *Config) ReadInConfig() error {
	configTree, err := toml.LoadFile(c.configFilePath)
	if err != nil {
		return err
	}
	c.configTree = configTree
	return nil
}

func Set(key string, value interface{}) {
	c.Set(key, value)
}

func (c *Config) Set(key string, value interface{}) {
	c.configTree.Set(key, value)
}

func (c *Config) getEnv(key string) (string, bool) {
	//if c.envKeyReplacer != nil {
	//	key = c.envKeyReplacer.Replace(key)
	//}
	key = strings.ToUpper(key)
	val, ok := os.LookupEnv(key)
	return val, ok
}

func (c *Config) Get(key string) interface{} {
	if val, ok := c.getEnv(key); ok {
		return val
	}
	if c.configTree == nil {
		return nil
	}
	return c.configTree.Get(key)
}

func GetString(key string) string {
	return c.GetString(key)
}

func (c *Config) GetString(key string) string {
	return cast.ToString(c.Get(key))
}

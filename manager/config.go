package manager

import (
	"github.com/BurntSushi/toml"
	"log"
)

func ParseConfig(data []byte) (*Config, error) {
	conf := &Config{}
	_, err := toml.Decode(string(data), conf)
	if err != nil {
		log.Println("Conf Parse Decode Error: ", err)
		return nil, err
	}
	return conf, nil
}

type Config struct {
	MaxWorkerNumber int                `toml:"max_worker_number"`
	Jobs            []ConfigJob        `toml:"job"`
	Log             ConfigLog          `toml:"job_log"`
	StatusServer    ConfigStatusServer `toml:"status_server"`
}

type ConfigJob struct {
	Host     string            `toml:"host"`
	Version  string            `toml:"version"`
	Image    string            `toml:"image"`
	Name     string            `toml:"name"`
	Command  Command           `toml:"command"`
	Env      []string          `toml:"env"`
	Volumes  []string          `toml:"volumes"`
	Interval TimeDuration      `toml:"interval"`
}

type ConfigLog struct {
	Type          CONFIG_LOG_TYPE `toml:"type"`
	RedisAddr     string          `toml:"redis_addr"`
	RedisPassword string          `toml:"redis_password"`
	RedisDB       int             `toml:"redis_db"`
	RedisKey      string          `toml:"redis_key"`
}

type CONFIG_LOG_TYPE string

const (
	CONFIG_LOG_TYPE_REDIS = "redis"
)

type ConfigStatusServer struct {
	Listen string `toml:"listen"`
	Port   string `toml:"port"`
}

func (c *ConfigStatusServer) GetAddr() string {
	return c.Listen + ":" + c.Port
}

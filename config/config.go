package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

//Config struct using viper
type Config struct {
	v *viper.Viper
}

//New Create a new config
func New() *Config {
	c := Config{
		v: viper.New(),
	}

	c.v.SetEnvPrefix("")
	c.v.AutomaticEnv()

	return &c
}

func (c *Config) GetString(name string, defaultvalue string) string {
	val := c.v.GetString(name)
	if val == "" {
		val = defaultvalue
	}
	return val
}

func (c *Config) GetInt(name string, defaultvalue int) int {
	val := c.v.GetInt(name)
	if val == 0 {
		val = defaultvalue
	}
	return val
}

func (c *Config) GetMaxThreads() int {
	return c.GetInt("MAX_THREADS", 10)
}

//GetTempFolder GetTempFolder
func (c *Config) GetTempFolder() string {
	return c.GetString("TEMPFOLDER", "/temp")
}

//GetAPIPort gets the main API port
func (c *Config) GetAPIPort() int {
	return c.GetInt("API_PORT", 30507)
}

//GetAPIHost gets the main API host
func (c *Config) GetAPIHost() string {
	return c.GetString("API_HOST", "0.0.0.0")
}

//APIHostPort The API host:port
func (c *Config) APIHostPort() string {
	return fmt.Sprintf("%s:%d", c.GetAPIHost(), c.GetAPIPort())
}

//GetLogLevel The log level
func (c *Config) GetLogLevel() string {
	return c.GetString("LOGLEVEL", "error")
}

//GetTemporalHost The Temporal host
func (c *Config) GetTemporalHost() string {
	return c.v.GetString("TEMPORAL_HOST")
}

//GetTemporalPort The Temporal Port
func (c *Config) GetTemporalPort() int64 {
	return c.v.GetInt64("TEMPORAL_PORT")
}

//TemporalHostPort Temporal Host:Port
func (c *Config) TemporalHostPort() string {
	return fmt.Sprintf("%s:%d", c.GetTemporalHost(), c.GetTemporalPort())

}

//GetTemporalTasklistName The Temporal task list name
func (c *Config) GetTemporalTasklistName() string {
	return c.GetString("TEMPORAL_TASKLISTNAME", "ptfdtasklist")
}

func (c *Config) DaprHostPort() string {
	return c.GetString("DAPR_HOST_PORT", "localhost:3500")

}

func (c *Config) RequestTimeout() int {
	timeout := c.v.GetInt("REQ_TIMEOUT")

	if timeout <= 0 {
		timeout = 60
	}
	return timeout

}

func (c *Config) RedisHostPort() string {
	return c.v.GetString("REDIS_HOST_PORT")

}

func (c *Config) RedisPwd() string {
	return c.v.GetString("REDIS_PWD")

}

func (c *Config) RMQUrl() string {
	str := "amqp://" + c.v.GetString("RMQ_USERNAME") + ":" + c.v.GetString("RMQ_PASSWORD") + "@" + c.v.GetString("RMQ_HOST") + ":" + c.v.GetString("RMQ_PORT") + "/"
	return str

}

func (c *Config) NatsURL() string {
	return c.v.GetString("NATS_URL")

}

func (c *Config) NatsNKeyPath() string {
	return c.GetString("NATS_NKEY_PATH", "/etc/nkey/nkey")

}

func (c *Config) DFEnv() string {
	return strings.ToLower(c.GetString("DF_ENV", "DEV"))
}

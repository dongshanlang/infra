package conf

import (
	"time"
)

var (
	confPath string
	Conf     *Config
)

// Config config
type Config struct {
	Debug     bool
	Log       *Log
	SqlMap    map[string]*Mysql `mapstructure:"sqlMap"`
	Redis     *Redis
	RPCServer *RPCServer
}

type Log struct {
	Dir string `yaml:"dir"`
}

type RPCServer struct {
	Network           string   `yaml:"network"`
	Addr              string   `yaml:"addr"`
	Timeout           time.Duration `yaml:"timeout"`
	IdleTimeout       time.Duration `yaml:"idleTimeout"`
	MaxLifeTime       time.Duration `yaml:"maxLifeTime"`
	ForceCloseWait    time.Duration `yaml:"ForceCloseWait"`
	KeepAliveInterval time.Duration `yaml:"KeepAliveInterval"`
	KeepAliveTimeout  time.Duration `yaml:"KeepAliveTimeout"`
	JaegerAddr        string   `yaml:"jaegerAddr"`
}

type Mysql struct {
	DSN         string   `yaml:"dsn"`
	MaxConn     int      `yaml:"max_connection"`
	MaxIdle     int      `yaml:"max_idle_connection"`
	MaxLifeTime time.Duration `yaml:"max_life_time"`
}

type Redis struct {
	Network      string   `yaml:"network"`
	Addr         string   `yaml:"addr"`
	Password     string   `yaml:"password"`
	DB           int      `yaml:"db"`
	DialTimeout  time.Duration `yaml:"dialTimeout"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	PoolSize     int      `yaml:"poolsize"`
	MinIdleConns int      `yaml:"minIdleConns"`
	IdleTimeout  time.Duration `yaml:"idleTimeout"`
}

package app

import (
	"github.com/spf13/pflag"
)

type ServerRunOptions struct {
	IP    string
	Port  string
	Redis string
}

func NewServerRunOptions() *ServerRunOptions {
	s := &ServerRunOptions{

	}
	return s
}

func (o *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	if o == nil {
		return
	}

	fs.StringVar(&o.IP, "serving-address", "127.0.0.1", "The IP serving railway query")
	fs.StringVar(&o.Redis, "redis-address", "127.0.0.1:6379", "The IP is connection to redis")
	fs.StringVar(&o.Port, "port", "8080", "The port serving railway query service")
}

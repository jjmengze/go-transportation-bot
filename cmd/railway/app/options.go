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

	fs.StringVar(&o.IP, "serving-address", o.IP, "The IP serving railway query")
	fs.StringVar(&o.Redis, "redis-address", o.Redis, "The IP is connection to redis")
	fs.StringVar(&o.Port, "port", o.Port, "The port serving railway query service")
}

package config

import (
	"flag"
)

type Config struct {
	LocalAddr  string
	RemoteAddr string
}

func Load() *Config {
	local := flag.String("l", ":8080", "Local address to listen on")
	remote := flag.String("r", "example.com:80", "Remote address to forward to")
	flag.Parse()

	return &Config{
		LocalAddr:  *local,
		RemoteAddr: *remote,
	}
}

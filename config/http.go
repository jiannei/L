package config

import "time"

type Http struct {
	Addr         string        `default:"127.0.0.1"`
	Port         string        `default:"8080"`
	ReadTimeout  time.Duration `default:"120"`
	WriteTimeout time.Duration `default:"120"`
}

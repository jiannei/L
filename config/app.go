package config

import "time"

type App struct {
	Name          string
	Mode          string        `default:"debug"`
	CancelTimeout time.Duration `default:"5"`
}

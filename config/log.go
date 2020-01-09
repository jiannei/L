package config

type Log struct {
	Switch     string `default:"off"`
	Path       string `default:"./storage/logs/"`
	FileName   string `default:"L.log"`
	Level      string `default:"info"`
	MaxSize    int    `default:"64"`
	MaxAge     int    `default:"7"`
	MaxBackups int    `default:"30"`
	Compress   bool   `default:"true"`
}

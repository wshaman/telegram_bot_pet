package config

import "os"

type Config struct {
	Token string
}

func Load() Config {
	c := Config{
		Token: os.Getenv("TG_TOKEN"),
	}
	return c
}

package config

import "time"

type Config struct {
	Server struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
	}

	Alerting struct {
		Interval time.Duration `toml:"interval"`
	}

	CPU struct {
		MaxUsage  float32 `toml:"max_usage"`
		AlertName string  `toml:"alert_name"`
		SendTo    string  `toml:"send_to"`
	}

	Memory struct {
		MaxUsage  float32 `toml:"max_usage"`
		AlertName string  `toml:"alert_name"`
		SendTo    string  `toml:"send_to"`
	}
}

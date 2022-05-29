package main

import (
	"encoding/json"
	"log"

	"github.com/rzaliyev/config"
)

type Config struct {
	APIkey string `json:"key"`
	APIurl string `json:"api,omitempty"`
}

func GetConfig() *Config {
	bs, err := config.Get(".config/news", "config")
	if err != nil {
		log.Fatal(err)
	}

	cfg := &Config{}
	if err = json.Unmarshal(bs, cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}

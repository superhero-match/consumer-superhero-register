package main

import (
	"github.com/consumer-superhero-register/cmd/consumer/reader"
	"github.com/consumer-superhero-register/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	r, err := reader.NewReader(cfg)
	if err != nil {
		panic(err)
	}

	err = r.Read()
	if err != nil {
		panic(err)
	}
}
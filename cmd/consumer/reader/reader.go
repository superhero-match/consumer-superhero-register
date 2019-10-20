package reader

import (
	"github.com/consumer-superhero-register/internal/config"
	"github.com/consumer-superhero-register/internal/consumer"
	"github.com/consumer-superhero-register/internal/db"
	"github.com/consumer-superhero-register/internal/es"
)

// Reader holds all the data relevant.
type Reader struct {
	DB       *db.DB
	ES       *es.ES
	Consumer *consumer.Consumer
}

// NewReader configures Reader.
func NewReader(cfg *config.Config) (r *Reader, err error) {
	dbs, err := db.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	e, err := es.NewES(cfg)
	if err != nil {
		return nil, err
	}

	c := consumer.NewConsumer(cfg)

	return &Reader{
		DB:       dbs,
		ES:       e,
		Consumer: c,
	}, nil
}

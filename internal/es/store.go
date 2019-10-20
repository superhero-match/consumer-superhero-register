package es

import (
	"context"
	"fmt"

	"github.com/consumer-superhero-register/internal/es/model"
)

// StoreSuperhero saves newly registered superhero in Elasticsearch.
func(es *ES) StoreSuperhero(s *model.Superhero) error {
	resp, err := es.Client.Index().
		Index(es.Index).
		BodyJson(s).
		Do(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Indexed Superhero %s to index %s, type %s\n", resp.Id, resp.Index, resp.Type)

	return nil
}
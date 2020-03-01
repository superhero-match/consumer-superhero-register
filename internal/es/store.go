/*
  Copyright (C) 2019 - 2020 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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
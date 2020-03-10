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
package reader

import (
	"context"
	"encoding/json"
	"fmt"
	dbm "github.com/superhero-match/consumer-superhero-register/internal/db/model"

	"gopkg.in/olivere/elastic.v7"

	chm "github.com/superhero-match/consumer-superhero-register/internal/cache/model"
	"github.com/superhero-match/consumer-superhero-register/internal/consumer/model"
	esm "github.com/superhero-match/consumer-superhero-register/internal/es/model"
)

// Read consumes the Kafka topic and stores the newly registered superhero to DB and Elasticsearch.
func (r *Reader) Read() error {
	ctx := context.Background()

	for {
		fmt.Print("before FetchMessage")
		m, err := r.Consumer.Consumer.FetchMessage(ctx)
		fmt.Print("after FetchMessage")
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		fmt.Printf(
			"message at topic/partition/offset \n%v/\n%v/\n%v: \n%s = \n%s\n",
			m.Topic,
			m.Partition,
			m.Offset,
			string(m.Key),
			string(m.Value),
		)

		var s model.Superhero
		if err := json.Unmarshal(m.Value, &s); err != nil {
			_ = r.Consumer.Consumer.Close()
			if err != nil {
				fmt.Println("Unmarshal")
				fmt.Println(err)
				err = r.Consumer.Consumer.Close()
				if err != nil {
					return err
				}

				return err
			}
		}

		err = r.DB.StoreSuperhero(dbm.Superhero{
			ID:                    s.ID,
			Email:                 s.Email,
			Name:                  s.Name,
			SuperheroName:         s.SuperheroName,
			MainProfilePicURL:     s.MainProfilePicURL,
			Gender:                s.Gender,
			LookingForGender:      s.LookingForGender,
			Age:                   s.Age,
			LookingForAgeMin:      s.LookingForAgeMin,
			LookingForAgeMax:      s.LookingForAgeMax,
			LookingForDistanceMax: s.LookingForDistanceMax,
			DistanceUnit:          s.DistanceUnit,
			Lat:                   s.Lat,
			Lon:                   s.Lon,
			Birthday:              s.Birthday,
			Country:               s.Country,
			City:                  s.City,
			SuperPower:            s.SuperPower,
			AccountType:           s.AccountType,
			FirebaseToken:         s.FirebaseToken,
			IsDeleted:             s.IsDeleted,
			DeletedAt:             s.DeletedAt,
			IsBlocked:             s.IsBlocked,
			BlockedAt:             s.BlockedAt,
			UpdatedAt:             s.UpdatedAt,
			CreatedAt:             s.CreatedAt,
		}, )
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		err = r.ES.StoreSuperhero(&esm.Superhero{
			ID:                    s.ID,
			Email:                 s.Email,
			Name:                  s.Name,
			SuperheroName:         s.SuperheroName,
			MainProfilePicURL:     s.MainProfilePicURL,
			Gender:                s.Gender,
			LookingForGender:      s.LookingForGender,
			Age:                   s.Age,
			LookingForAgeMin:      s.LookingForAgeMin,
			LookingForAgeMax:      s.LookingForAgeMax,
			LookingForDistanceMax: s.LookingForDistanceMax,
			DistanceUnit:          s.DistanceUnit,
			Location: elastic.GeoPoint{
				Lat: s.Lat,
				Lon: s.Lon,
			},
			Birthday:    s.Birthday,
			Country:     s.Country,
			City:        s.City,
			SuperPower:  s.SuperPower,
			AccountType: s.AccountType,
			IsDeleted:   s.IsDeleted,
			DeletedAt:   nil,
			IsBlocked:   s.IsBlocked,
			BlockedAt:   nil,
			UpdatedAt:   s.UpdatedAt,
			CreatedAt:   s.CreatedAt,
		}, )
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		err = r.Cache.SetToken(chm.FirebaseMessagingToken{
			Token:       s.FirebaseToken,
			SuperheroID: s.ID,
			CreatedAt:   s.CreatedAt,
		})
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		err = r.Consumer.Consumer.CommitMessages(ctx, m)
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}
	}
}

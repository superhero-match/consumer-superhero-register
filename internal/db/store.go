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
package db

import "github.com/consumer-superhero-register/internal/db/model"

// StoreSuperhero saves newly registered Superhero.
func(db *DB) StoreSuperhero (s model.Superhero) error {
	_, err := db.stmtInsertNewSuperhero.Exec(
		s.ID,
		s.Email,
		s.Name,
		s.SuperheroName,
		s.MainProfilePicURL,
		s.Gender,
		s.LookingForGender,
		s.Age,
		s.LookingForAgeMin,
		s.LookingForAgeMax,
		s.LookingForDistanceMax,
		s.DistanceUnit,
		s.Lat,
		s.Lon,
		s.Birthday,
		s.Country,
		s.City,
		s.SuperPower,
		s.AccountType,
		s.FirebaseToken,
		s.IsDeleted,
		s.DeletedAt,
		s.IsBlocked,
		s.BlockedAt,
		s.UpdatedAt,
		s.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

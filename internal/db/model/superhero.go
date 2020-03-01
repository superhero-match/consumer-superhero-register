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
package model

type Superhero struct {
	ID                    string  `db:"id"`
	Email                 string  `db:"email"`
	Name                  string  `db:"name"`
	SuperheroName         string  `db:"superhero_name"`
	MainProfilePicURL     string  `db:"main_profile_pic_url"`
	Gender                int     `db:"gender"`
	LookingForGender      int     `db:"looking_for_gender"`
	Age                   int     `db:"age"`
	LookingForAgeMin      int     `db:"looking_for_age_min"`
	LookingForAgeMax      int     `db:"looking_for_age_max"`
	LookingForDistanceMax int     `db:"looking_for_distance_max"`
	DistanceUnit          string  `db:"distance_unit"`
	Lat                   float64 `db:"lat"`
	Lon                   float64 `db:"lon"`
	Birthday              string  `db:"birthday"`
	Country               string  `db:"country"`
	City                  string  `db:"city"`
	SuperPower            string  `db:"superpower"`
	AccountType           string  `db:"account_type"`
	FirebaseToken         string  `db:"firebase_token"`
	IsDeleted             bool    `db:"is_deleted"`
	IsBlocked             bool    `db:"is_blocked"`
	UpdatedAt             string  `db:"updated_at"`
	CreatedAt             string  `db:"created_at"`
	BlockedAt             string  `db:"blocked_at"`
	DeletedAt             string  `db:"deleted_at"`
}

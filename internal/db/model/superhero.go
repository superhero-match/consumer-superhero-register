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
	IsDeleted             bool    `db:"is_deleted"`
	IsBlocked             bool    `db:"is_blocked"`
	UpdatedAt             string  `db:"updated_at"`
	CreatedAt             string  `db:"created_at"`
	BlockedAt             string  `db:"blocked_at"`
	DeletedAt             string  `db:"deleted_at"`
}

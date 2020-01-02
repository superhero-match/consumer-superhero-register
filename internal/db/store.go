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

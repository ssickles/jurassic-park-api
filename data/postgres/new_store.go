package postgres

import "jurassic-park-api/data"

func NewStore() (data.Store, error) {
	db, err := CreateDbConnection()
	if err != nil {
		return data.Store{}, err
	}

	return data.Store{
		Dinosaurs: Dinosaurs{Db: db},
	}, nil
}

package repository

import (
	"database/sql"
	model "effectiveMobileTest/internal/model"
)

type Repository struct {
	db sql.DB
}

func (r *Repository) Insert(car model.Car) error {
	query := "insert into Cars (regNum,mark,model,owner) values ($1,$2,$3,$4)"
	if _, err := r.db.Exec(
		query,
		car.RegNum,
		car.Mark,
		car.Model,
		car.Owner,
	); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Get(limit, offset int) (cars []model.Car, e error) {
	query := "select regNum,mark,model,owner from Cars LIMIT $1 OFFSET $1"

	rows, err := r.db.Query(
		query,
		limit,
		offset,
	)

	if err != nil {
		e = err
	}
	defer rows.Close()
	for rows.Next() {
		var car model.Car
		if err := rows.Scan(
			&car.RegNum,
			&car.Mark,
			&car.Model,
			&car.Owner,
		); err != nil {
			e = err
			return
		} else {
			cars = append(cars, car)
		}
	}
	return
}

func (r *Repository) Delete(regNum string) error {
	query := "delete from Cars where regNum = $1"
	if _, err := r.db.Exec(
		query,
		regNum,
	); err != nil {
		return err
	}
	return nil
}
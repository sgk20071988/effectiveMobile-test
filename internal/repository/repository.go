package repository

import (
	"database/sql"
	model "effectiveMobileTest/internal/model"

	"github.com/huandu/go-sqlbuilder"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) Insert(car model.Car) error {
	queryBuilder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	queryBuilder.InsertInto("Car")
	queryBuilder.Cols("mark", "model", "owner", "regNum")
	queryBuilder.Values(car.Mark, car.Model, car.Owner, car.RegNum)
	sql, args := queryBuilder.Build()
	if _, err := r.DB.Exec(
		sql,
		args,
	); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetCars(limit, offset int, conds map[string]string) (cars []model.Car, e error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.From("Cars")
	sb.Select("regNum", "mark", "model", "owner")
	sb.Limit(limit)
	sb.Offset(offset)
	for col, like := range conds {
		sb.Where(
			sb.Like(col, like),
		)
	}
	query, args := sb.Build()

	rows, err := r.DB.Query(
		query,
		args,
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

func (r *Repository) DeleteCar(regNum string) error {
	db := sqlbuilder.PostgreSQL.NewDeleteBuilder()
	db.DeleteFrom("Cars")
	db.Where(
		db.EQ("regNum", regNum),
	)
	query, args := db.Build()
	if _, err := r.DB.Exec(
		query,
		args,
	); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(regNum string, update map[string]string) error {
	ub := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	ub.Update("Cars")
	for col, val := range update {
		ub.Set(
			ub.Assign(col, val),
		)
	}
	ub.Where(
		ub.Equal("regNum", regNum),
	)

	query, args := ub.Build()
	if _, err := r.DB.Exec(
		query,
		args,
	); err != nil {
		return err
	}
	return nil
}

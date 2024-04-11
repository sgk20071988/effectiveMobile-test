package repository

import (
	"database/sql"
	model "effectiveMobileTest/internal/model"
	"strconv"

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

func (r *Repository) GetCar(regNum string) (car model.Car, e error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.From("Cars")
	sb.Select("regNum", "mark", "model", "owner")
	sb.Where(
		sb.Equal("regNum", regNum),
	)
	query, args := sb.Build()

	rows, err := r.DB.Query(
		query,
		args,
	)

	if err != nil {
		e = err
		return
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&car.RegNum,
			&car.Mark,
			&car.Model,
			&car.Owner,
		); err != nil {
			e = err
			return
		}
	}
	return

}

func (r *Repository) GetCars(params map[string]string, filters map[string]string) (cars []model.Car, e error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.From("Cars")
	sb.Select("regNum", "mark", "model", "owner")
	limit, ok := params["limit"]
	if ok {
		liminInt, err := strconv.Atoi(limit)
		if err != nil {
			e = err
			return
		}
		sb.Limit(liminInt)
	}
	offset, ok := params["offset"]
	if ok {
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			e = err
			return
		}
		sb.Offset(offsetInt)
	}

	for col, like := range filters {
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

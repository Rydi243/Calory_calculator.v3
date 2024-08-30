package store

import (
	"Calorie_calculator/postgres/public/model"
	"Calorie_calculator/postgres/public/table"
	"context"
	"database/sql"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return Store{db: db}
}

func (s *Store) InsertProduct(m *model.Calorie) (*model.Calorie, error) {
	var res model.Calorie

	stmt := table.Calorie.INSERT(table.Calorie.AllColumns.Except(table.Calorie.ID)).
		MODEL(m).
		RETURNING(table.Calorie.AllColumns)

	err := stmt.QueryContext(context.Background(), s.db, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *Store) GetAllProduct() ([]*model.Calorie, error) {
	res := make([]*model.Calorie, 0)

	stmt := table.Calorie.SELECT(table.Calorie.AllColumns)

	err := stmt.QueryContext(context.Background(), s.db, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

//	func (s *Store) GetCalProd(m *model.Calorie) (int32, error) {
//		var res model.Calorie
//		var resit int32
//		stmt := table.Calorie.SELECT(
//			table.Calorie.Calorie,
//		).
//			WHERE(table.Calorie.Product.EQ(postgres.String(m.Product)))
//		err := stmt.Query(s.db, &res)
//		if err != nil {
//			return 0, err
//		}
//		resit = res.Calorie
//		return resit, err
//	}
//
// (float64, float64, error)
func (s *Store) Calc(m *model.Calorie) (float64, error) {
	var ch model.Calorie
	stmt := table.Calorie.SELECT(
		table.Calorie.Calorie,
	).
		WHERE(table.Calorie.Product.EQ(postgres.String(m.Product)))
	err := stmt.Query(s.db, &ch)
	if err != nil {
		return 0, err
	}

	fmt.Println(ch)

	res := (float64(ch.Calorie) / 100) * float64(m.Weight)

	return res, err
	//77 / 100 * 150 =115.5
	//330/100 * 100 = 330
	//115.5 + 330 = 445.5

	//77/100 * 360 = 277.2
	//112/100 * 275 = 308
	//227.2 + 308 =535.2

}

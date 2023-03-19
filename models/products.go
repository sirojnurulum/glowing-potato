package models

type Products struct {
	Id    int     `db:"id"`
	Name  string  `db:"name"`
	Price float64 `db:"price"`
}

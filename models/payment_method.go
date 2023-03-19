package models

type PaymentMethod struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	IsActive bool   `db:"is_active"`
}

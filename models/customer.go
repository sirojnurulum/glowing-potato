package models

type Customer struct {
	Id           int    `db:"id"`
	CustomerName string `db:"customer_name"`
}

type CustomerAddress struct {
	Id         int    `db:"id"`
	CustomerId int    `db:"customer_id"`
	Address    string `db:"address"`
}

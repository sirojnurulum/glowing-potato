package objects

type Customer struct {
	Id           int    `json:"id"`
	CustomerName string `json:"customer_name"`
}

type CustomerAddress struct {
	Id         int    `json:"id"`
	CustomerId int    `json:"customer_id"`
	Address    string `json:"address"`
}

package models

type Product struct {
	Uid       string   `json:"uid,omitempty"`
	Id        string   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Price     int      `json:"price,omitempty"`
	QueryDate string   `json:"query_date,omitempty"`
	DType     []string `json:"dgraph.type,omitempty"`
}

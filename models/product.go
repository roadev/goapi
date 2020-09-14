package models

type Product struct {
	Uid       string   `json:"uid,omitempty"`
	Id        int      `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Price     string   `json:"price,omitempty"`
	QueryDate string   `json:"query_date,omitempty"`
	DType     []string `json:"dgraph.type,omitempty"`
}

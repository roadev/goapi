package models

type Product struct {
	Uid   string   `json:"uid,omitempty"`
	Name  string   `json:"name,omitempty"`
	Price string   `json:"price,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

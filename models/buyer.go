package models

type Buyer struct {
	Uid   string   `json:"uid,omitempty"`
	Name  string   `json:"name,omitempty"`
	Age   string   `json:"age,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

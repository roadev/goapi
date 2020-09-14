package models

type Buyer struct {
	Uid   string   `json:"uid,omitempty"`
	Name  string   `json:"name,omitempty"`
	Age   int      `json:"age,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

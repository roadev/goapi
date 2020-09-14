package models

type Buyer struct {
	Uid       string   `json:"uid,omitempty"`
	Id        int      `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Age       int      `json:"age,omitempty"`
	QueryDate string   `json:"query_date,omitempty"`
	DType     []string `json:"dgraph.type,omitempty"`
}

type BuyerResponse struct {
	Buyers []Buyer `json:"buyers"`
}

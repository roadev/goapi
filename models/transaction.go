package models

type Transaction struct {
	Uid        string   `json:"uid,omitempty"`
	Id         int      `json:"id,omitempty"`
	BuyerId    string   `json:"buyer_id,omitempty"`
	Ip         string   `json:"ip,omitempty"`
	Device     string   `json:"device,omitempty"`
	ProductIds []Buyer  `json:"product_ids,omitempty"`
	QueryDate  string   `json:"query_date,omitempty"`
	DType      []string `json:"dgraph.type,omitempty"`
}

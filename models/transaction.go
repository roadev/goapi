package models

type Transaction struct {
	Uid        string    `json:"uid,omitempty"`
	Id         string    `json:"id,omitempty"`
	BuyerId    string    `json:"buyer_id,omitempty"`
	Ip         string    `json:"ip,omitempty"`
	Device     string    `json:"device,omitempty"`
	ProductIds []string  `json:"product_ids,omitempty"`
	Products   []Product `json:"products,omitempty"`
	QueryDate  string    `json:"query_date,omitempty"`
	DType      []string  `json:"dgraph.type,omitempty"`
}

type TransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
}

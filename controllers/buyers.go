package controllers

import (
	"context"
	"fmt"
	"github.com/dgraph-io/dgo/v200"
	// "log"
	"net/http"
	// "github.com/roadev/goapi/server"
)

type BuyerController struct {
}

func GetAllBuyers(dgraphClient *dgo.Dgraph, ctx context.Context, w http.ResponseWriter) {
	query := `{
		buyers(func: has(name)) {
			uid
			id
			name
			age
		}
	}`
	// // variables := map[string]string {"$id1": response.Uids[]

	response, err := dgraphClient.NewTxn().Query(ctx, query)

	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}

	fmt.Println(response)
	// fmt.Println(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response.Json))

}

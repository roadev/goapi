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
		all(func: has(name)) {
			uid
			name
			age
		}
	}`
	// // variables := map[string]string {"$id1": response.Uids[]

	response, _ := dgraphClient.NewTxn().Query(context.Background(), query)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(query)
	fmt.Println(response)

	w.Write([]byte("pong"))

}

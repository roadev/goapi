package controllers

import (
	"context"
	"fmt"
	"github.com/dgraph-io/dgo/v200"
	"io/ioutil"
	"log"
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

func LoadBuyers(date int) {
	response, err := http.Get(fmt.Sprintf("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers?%d", date))

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	fmt.Println(responseString)

	// fmt.Println("Response: ", response.Body)

}

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	// "github.com/roadev/goapi/models"
	"github.com/roadev/goapi/utils"
	"io/ioutil"
	"log"
	"net/http"
	// "strings"
)

type TransactionController struct {
}

func GetAllTransactionsByBuyer(dgraphClient *dgo.Dgraph, ctx context.Context, w http.ResponseWriter, buyerId string) {
	// buyerQuery := fmt.Sprintf(`
	// 	{
	// 		buyer(func eq(id@en, %s))
	// 		@filter(has(age)) {
	// 			uid
	// 			id
	// 			name
	// 			age
	// 		}
	// 	}`,
	// 	buyerId,
	// )

	// transactionsQuery := `{
	// 	transactions(func: has(device)) {
	// 		uid
	// 		id
	// 		ip
	// 		device
	// 		buyer_id
	// 		producto_ids
	// 	}
	// }`

	// transactionsResponse, err := dgraphClient.NewTxn().Query(ctx, transactionsQuery)

	// if err != nil {
	// 	// log.Fatal(err)
	// 	fmt.Println(err)
	// }

	// query := `{
	// 	products(func: has(price)) {
	// 		uid
	// 		id
	// 		name
	// 		price
	// 	}
	// }`

	// productsResponse, err := dgraphClient.NewTxn().Query(ctx, query)

	// if err != nil {
	// 	// log.Fatal(err)
	// 	fmt.Println(err)
	// }

	// fmt.Println(productsResponse)

	// fmt.Println(response)
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(response.Json))

}

func LoadTransactions(dgraphClient *dgo.Dgraph, ctx context.Context, w http.ResponseWriter, date string) {

	response, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions?date=" + date)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	parsedTransactionList := utils.TransformTransactionsData(string(responseData))

	out, _ := json.Marshal(parsedTransactionList)

	// fmt.Println(out)

	// var products []models.Product

	// if err := json.Unmarshal(responseData, &products); err != nil {
	// 	panic(err)
	// }

	mutation := &api.Mutation{
		CommitNow: true,
		SetJson:   out,
	}

	res, err := dgraphClient.NewTxn().Mutate(ctx, mutation)

	if err != nil {
		rawJsonError := fmt.Sprintf(`
		{
			"error": "There was a problem importing the products data"
		}`,
			date,
		)

		jsonData := []byte(rawJsonError)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(jsonData))

		log.Fatal(err)
	} else {
		rawJson := fmt.Sprintf(`
		{
			"message": "%d transactions have been imported for the given datetime",
			"query_date": "%s"
		}`,
			len(parsedTransactionList),
			date,
		)
		jsonData := []byte(rawJson)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	}

	fmt.Println(res)

	// out, _ := json.MarshalIndent(buyers, "", "    ")
	// // fmt.Println(buyers[0].Name)

	// responseString := string(out)
	// fmt.Println(responseString)

	// fmt.Println("Response: ", response.Body)

}

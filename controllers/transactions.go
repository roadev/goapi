package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	unique "github.com/mpvl/unique"
	"github.com/roadev/goapi/models"
	"github.com/roadev/goapi/utils"
	"io/ioutil"
	"log"
	"net/http"
	// "strings"
)

type TransactionController struct {
}

func GetAllTransactionsByBuyer(dgraphClient *dgo.Dgraph, ctx context.Context, w http.ResponseWriter, buyerId string) {
	// 	buyerQuery := fmt.Sprintf(`
	// 		{
	// 			buyer(func eq(id@en, %s))
	// 			@filter(has(age)) {
	// 				uid
	// 				id
	// 				name
	// 				age
	// 			}
	// 		}`,
	// 		buyerId,
	// 	)

	// 	fmt.Println(buyerQuery)

	var transactions models.TransactionResponse

	transactionsQuery := fmt.Sprintf(`
		{
			transactions(func: eq(buyer_id, %s))
			@filter(has(device)) {
				uid
				id
				ip
				device
				buyer_id
				product_ids
			}
		}`,
		buyerId,
	)

	transactionsResponse, err := dgraphClient.NewTxn().Query(ctx, transactionsQuery)

	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}

	fmt.Println(transactionsResponse)

	if err := json.Unmarshal(transactionsResponse.Json, &transactions); err != nil {
		panic(err)
	}

	if len(transactions.Transactions) == 0 {
		out, _ := json.Marshal(transactions)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(out))
	}

	// fmt.Println(transactions.Transactions[0])

	// var transactionsWithProducts []models.Transaction
	var productIds []string

	for _, transaction := range transactions.Transactions {
		productIds = append(productIds, transaction.ProductIds...)
	}

	unique.Strings(&productIds)

	// fmt.Println(productIds)

	productsQuery := fmt.Sprintf(
		`{
			products(func: eq(id, %s)) {
				uid
				id
				name
				price
			}
		}`,
		productIds,
	)

	var products models.ProductResponse

	// fmt.Println(len(productIds))

	productsResponse, err := dgraphClient.NewTxn().Query(ctx, productsQuery)

	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}

	if err := json.Unmarshal(productsResponse.Json, &products); err != nil {
		panic(err)
	}

	// fmt.Println(products)

	for i := 0; i < len(transactions.Transactions); i++ {
		// var parsedProducts []models.Product

		for j := 0; j < len(transactions.Transactions[i].ProductIds); j++ {

			for k := 0; k < len(products.Products); k++ {
				if products.Products[k].Id == transactions.Transactions[i].ProductIds[j] {
					transactions.Transactions[i].Products = append(transactions.Transactions[i].Products, products.Products[k])
				}
			}
		}

	}

	// fmt.Println(transactions.Transactions[0].Products)

	out, _ := json.Marshal(transactions)
	// fmt.Println(string(out))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out))

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

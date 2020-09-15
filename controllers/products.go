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

type ProductController struct {
}

func GetAllProducts(dgraphClient *dgo.Dgraph, ctx context.Context, w http.ResponseWriter) {
	query := `{
		products(func: has(price)) {
			uid
			id
			name
			price
		}
	}`

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

func LoadProducts(dgraphClient *dgo.Dgraph, ctx context.Context, w http.ResponseWriter, date string) {
	response, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products?date=" + date)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	parsedProductsList := utils.TransformProductsData(string(responseData))

	out, _ := json.Marshal(parsedProductsList)

	fmt.Println(out)

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
		log.Fatal(err)
	}

	fmt.Println(res)

	rawJson := fmt.Sprintf(`
		{
			"message": "Products have been imported for the given datetime",
			"query_date": "%s"
		}`,
		date,
	)

	jsonData := []byte(rawJson)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonData))

	// out, _ := json.MarshalIndent(buyers, "", "    ")
	// // fmt.Println(buyers[0].Name)

	// responseString := string(out)
	// fmt.Println(responseString)

	// fmt.Println("Response: ", response.Body)

}

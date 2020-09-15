package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/roadev/goapi/models"
	"io/ioutil"
	"log"
	"net/http"
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

func LoadBuyers(dgraphClient *dgo.Dgraph, ctx context.Context, w http.ResponseWriter, date int) {
	response, err := http.Get(fmt.Sprintf("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers?%d", date))

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var buyers []models.Buyer

	if err := json.Unmarshal(responseData, &buyers); err != nil {
		panic(err)
	}

	// pb, err := json.Marshal(buyers)

	mutation := &api.Mutation{
		CommitNow: true,
		SetJson:   responseData,
	}

	res, err := dgraphClient.NewTxn().Mutate(ctx, mutation)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	rawJson := fmt.Sprintf(`
	{
		"message": "Buyers have been imported for the given datetime"
		"query_date": "%d"
	}`, date)

	jsonData := []byte(rawJson)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonData))

	// out, _ := json.MarshalIndent(buyers, "", "    ")
	// // fmt.Println(buyers[0].Name)

	// responseString := string(out)
	// fmt.Println(responseString)

	// fmt.Println("Response: ", response.Body)

}

package database

// "github.com/roadev/goapi/models/buyer"

import (
	"context"
	"encoding/json"
	"fmt"
	dgo "github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/roadev/goapi/models"
	"google.golang.org/grpc"
	"log"
	// "strings"
)

func NewDatabaseConnection() (*dgo.Dgraph, context.Context) {
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC", err)
	}

	defer conn.Close()
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	operation := &api.Operation{}

	operation.Schema = `		
		name: string @index(exact) .
		id: int @index(int) .
		age: int .
		price: float .
		Buyer: [uid] .
		Products: [uid] .
		Transactions: [uid] .
		ip: string .
		device: string .
		query_date: dateTime .

		type Buyer {
			id: int
			name: string
			age: int
			query_date: dateTime
			Transactions: [Transaction]
		}

		type Transaction {
			id: int
			Buyer: [Buyer]
			ip: string
			device: string
			query_date: dateTime
			Products: [Product]
		}

		type Product {
			id: int
			name: string
			query_date: dateTime
			price: float
		}
	`

	// buyer := models.Buyer{
	// 	Id:   2,
	// 	Uid:  "_:juan",
	// 	Name: "Juan",
	// 	Age:  28,
	// }

	query := `{
		buyers(func: has(name)) {
			uid
			id
			name
			age
		}
	}`

	// query := `{
	// 	buyers(func: has(name)) {
	// 		name
	// 	}
	// }`

	ctx := context.Background()

	// if err := dgraphClient.Alter(ctx, operation); err != nil {
	// 	log.Fatal("Alter error: ", err)
	// }

	// pb, err := json.Marshal(buyer)

	// mutation := &api.Mutation{
	// 	CommitNow: true,
	// 	SetJson:   pb,
	// }

	response, err := dgraphClient.NewTxn().Query(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(query)
	// fmt.Println(response.Json)

	var buyer models.BuyerResponse

	// err = json.Unmarshal(response.Json, &buyer)
	// // decoder := json.NewDecoder(strings.NewReader(response.Json))
	// // err := decoder.Decode(&buyer)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	if err := json.Unmarshal(response.GetJson(), &buyer); err != nil {
		panic(err)
	}

	// fmt.Println(buyer)

	out, _ := json.MarshalIndent(buyer, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(out)

	fmt.Printf("%s\n", out)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// req := &api.Request{CommitNow: true, Mutations: []*api.Mutation{mutation}}

	// response, err := dgraphClient.NewTxn().Mutate(ctx, mutation)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(response.Json))

	return dgraphClient, ctx

}

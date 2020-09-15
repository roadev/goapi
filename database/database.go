package database

// "github.com/roadev/goapi/models/buyer"

import (
	"context"
	// "encoding/json"
	// "fmt"
	dgo "github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	// "github.com/roadev/goapi/models"
	"google.golang.org/grpc"
	"log"
)

func NewDatabaseConnection() (*dgo.Dgraph, context.Context) {
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC", err)
	}

	// defer conn.Close()
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	operation := &api.Operation{}

	operation.Schema = `		
		name: string @index(exact) .
		id: string @index(exact) .
		age: int .
		price: int .
		Buyer: [uid] .
		Products: [uid] .
		Transactions: [uid] .
		ip: string .
		device: string .
		query_date: dateTime .

		type Buyer {
			id: string
			name: string
			age: int
			query_date: dateTime
			Transactions: [Transaction]
		}

		type Transaction {
			id: string
			Buyer: [Buyer]
			ip: string
			device: string
			query_date: dateTime
			Products: [Product]
		}

		type Product {
			id: string
			name: string
			query_date: dateTime
			price: int
		}
	`

	// buyer := models.Buyer{
	// 	Id:   2,
	// 	Uid:  "_:juan",
	// 	Name: "Juan",
	// 	Age:  28,
	// }

	// query := `{
	// 	buyers(func: has(name)) {
	// 		uid
	// 		id
	// 		name
	// 		age
	// 	}
	// }`

	ctx := context.Background()

	if err := dgraphClient.Alter(ctx, operation); err != nil {
		log.Fatal("Alter error: ", err)
	}

	// pb, err := json.Marshal(buyer)

	// mutation := &api.Mutation{
	// 	CommitNow: true,
	// 	SetJson:   pb,
	// }

	// req := &api.Request{CommitNow: true, Mutations: []*api.Mutation{mutation}}

	// response, err := dgraphClient.NewTxn().Mutate(ctx, mutation)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// response, err := dgraphClient.NewTxn().Query(ctx, query)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(query)

	// var buyer models.BuyerResponse

	// if err := json.Unmarshal(response.GetJson(), &buyer); err != nil {
	// 	panic(err)
	// }

	// out, _ := json.MarshalIndent(buyer, "", "    ")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%s\n", out)

	return dgraphClient, ctx

}

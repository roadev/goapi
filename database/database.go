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
)

func NewDatabaseConnection() {
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC", err)
	}

	defer conn.Close()
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	operation := &api.Operation{}

	operation.Schema = `		
		name: string @index(exact) .
		age: int .
		price: float .
		Buyer: [uid] .
		Products: [uid] .
		Transactions: [uid] .
		ip: string .
		device: string .

		type Buyer {
			name: string
			age: int
			Transactions: [Transaction]
		}

		type Transaction {
			Buyer: [Buyer]
			ip: string
			device: string
			Products: [Product]
		}

		type Product {
			name: string
			price: float
		}
	`

	buyer := models.Buyer{
		Uid:  "_:juan",
		Name: "Juan",
		Age:  28,
	}

	ctx := context.Background()
	if err := dgraphClient.Alter(ctx, operation); err != nil {
		log.Fatal("Alter error: ", err)
	}

	pb, err := json.Marshal(buyer)

	mutation := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	if err != nil {
		log.Fatal(err)
	}

	// req := &api.Request{CommitNow: true, Mutations: []*api.Mutation{mutation}}

	response, err := dgraphClient.NewTxn().Mutate(ctx, mutation)

	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)

}

// name: string @index(exact) .
// 		age: int .
// 		price float .
// 		buyer_id int .

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
		buyer_id: string @index(exact) .
		products: [int] .
		ip: string .
		device: string .
		query_date: dateTime .

		type Buyer {
			id: string
			name: string
			age: int
			query_date: dateTime
		}

		type Transaction {
			id: string
			buyer_id: string
			ip: string
			device: string
			query_date: dateTime
			products: [int]
		}

		type Product {
			id: string
			name: string
			query_date: dateTime
			price: int
		}
	`

	ctx := context.Background()

	if err := dgraphClient.Alter(ctx, operation); err != nil {
		log.Fatal("Alter error: ", err)
	}

	return dgraphClient, ctx

}

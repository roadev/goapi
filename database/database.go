package database

// "github.com/roadev/goapi/models/buyer"

import (
	"context"
	"fmt"
	dgo "github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
	"log"
)

func NewDatabaseConnection() {
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC", err)
	}

	fmt.Println("Hiiiii!!!!")

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

	ctx := context.Background()
	if err := dgraphClient.Alter(ctx, operation); err != nil {
		log.Fatal("Alter error: ", err)
	}

	// mutation = &api.Mutation{
	// 	CommitNow: true,
	// }

	// pb, err := json.Marshal(p)

}

// name: string @index(exact) .
// 		age: int .
// 		price float .
// 		buyer_id int .

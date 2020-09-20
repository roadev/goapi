package utils

import (
	"fmt"
	"github.com/roadev/goapi/models"
	"strconv"
	"strings"
)

func TransformProductsData(s string) []models.Product {

	var products []models.Product
	trimed := strings.TrimSpace(s)

	ss := strings.Split(trimed, "\n")

	// fmt.Println(ss)
	// m = make(map[string]string)
	// init := 3

	for i := 0; i < len(ss); i++ {
		// fmt.Println(ss)
		raw := strings.Split(ss[i], "'")
		fmt.Println(raw)
		price, _ := strconv.Atoi(raw[2])
		p := models.Product{
			Id:    raw[0],
			Name:  raw[1],
			Price: price,
		}
		products = append(products, p)

	}
	// fmt.Println(products)

	return products
}

func TransformTransactionsData(s string) []models.Transaction {

	var transactions []models.Transaction
	trimed := strings.TrimSpace(s)

	ss := strings.Split(trimed, "��")

	// fmt.Println(ss)
	// m = make(map[string]string)
	// init := 3

	for i := 0; i < len(ss); i++ {
		// fmt.Println(ss)
		raw := strings.Split(ss[i], "�")

		rawProductIdsA := strings.Replace(raw[4], "(", "", 1)
		rawProductIds := strings.Replace(rawProductIdsA, ")", "", 1)
		productIds := strings.Split(rawProductIds, ",")

		// fmt.Println(productIds)

		// var productIds string = raw[4]
		p := models.Transaction{
			Id:         raw[0],
			BuyerId:    raw[1],
			Ip:         raw[2],
			Device:     raw[3],
			ProductIds: productIds,
		}

		transactions = append(transactions, p)

	}
	// fmt.Println(products)

	return transactions
}

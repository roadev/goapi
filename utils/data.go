package utils

import (
	"fmt"
	"github.com/roadev/goapi/models"
	"regexp"
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

func cleanUnicode(s string) string {
	return strings.Replace(s, "\u0000", "", -1)
}

func findIpFromString(s string, n int) []int {
	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	return re.FindAllStringIndex(s, n)[0]
}

func TransformTransactionsData(s string) []models.Transaction {

	var transactions []models.Transaction
	// trimed := strings.TrimSpace(s)

	trimed := strings.TrimSpace(s)
	split1 := strings.Split(trimed, "#")
	split2 := split1[1:len(split1)]
	// this regex tests a valid IP in a string

	for i := 0; i < len(split2); i++ {
		ipIndexRange := findIpFromString(split2[i][20:len(split2[i])-1], 1)
		//ipStart := ipIndexRange[0][0:1]
		fmt.Println(ipIndexRange)
		deviceString := split2[i][20+ipIndexRange[1] : strings.Index(split2[i], "(")]
		productIds := strings.Split(split2[i], "(")[1]

		t := models.Transaction{
			Id:         split2[i][0:12],
			BuyerId:    cleanUnicode(split2[i][12:20]),
			Ip:         split2[i][20 : len(split2[i])-1][ipIndexRange[0]:ipIndexRange[1]],
			Device:     cleanUnicode(deviceString),
			ProductIds: strings.Split(strings.Replace(cleanUnicode(productIds[0:len(productIds)-1]), ")", "", 1), ","),
		}

		transactions = append(transactions, t)
	}
	// fmt.Println(products)

	return transactions
}

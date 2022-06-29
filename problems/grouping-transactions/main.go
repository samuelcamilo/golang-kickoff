package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	id int8
	paymentDate string
	paymentMethod string
	paymentValue float64
}

// retorno nomeado aumenta a performance, pois n será necessário advinhar o que será retornado...
func seed() (arr []Transaction) {
	return []Transaction {
		{id: 1, paymentDate: time.Now().UTC().Format("2006-01-02"), paymentMethod: "AA", paymentValue: 2400.00},
		{id: 2, paymentDate: time.Now().UTC().Format("2006-01-02"), paymentMethod: "AB", paymentValue: 100.00},
		{id: 3, paymentDate: time.Now().UTC().Format("2006-01-02"), paymentMethod: "AA", paymentValue: 2500.00},
		{id: 4, paymentDate: time.Now().UTC().Format("2006-01-02"), paymentMethod: "AA", paymentValue: 200.00},
		{id: 5, paymentDate: time.Now().UTC().Format("2006-01-02"), paymentMethod: "AB", paymentValue: 2500.00},
		{id: 6, paymentDate: time.Now().UTC().Format("2006-01-02"), paymentMethod: "AB", paymentValue: 2100.00},
		{id: 7, paymentDate: time.Now().UTC().Format("2006-01-02"), paymentMethod: "AB", paymentValue: 500.00},
		{id: 8, paymentDate: time.Now().UTC().Format("2006-01-02"), paymentMethod: "AB", paymentValue: 1500.00},
	}
}

var transactions []Transaction = seed()

func main() {
	dict := make(map[string][]Transaction)
	
	for i := range transactions {
		if _, ok := dict[transactions[i].paymentMethod]; !ok {
			dict[transactions[i].paymentMethod] = []Transaction {transactions[i]}
		} else {
			dict[transactions[i].paymentMethod] = append(dict[transactions[i].paymentMethod], transactions[i])
		}
	}

	fmt.Println(dict)
}
package main

import (
	"fmt"
	"strconv"
)

func main() {
	items := [][]string{{"Cake","0.99","7.5"},
		{"Milk","2.75","1.5"},
		{"Butter","0.87","2"}}

	var salesTaxTotal float32

	for _, item := range items {
		cost, _ :=  strconv.ParseFloat(item[1], 32)
		salesTaxRate, _ :=  strconv.ParseFloat(item[2], 32)
		salesTaxTotal += float32(cost * salesTaxRate * 0.01)
	}

	fmt.Println("Sales Tax Total:", salesTaxTotal)
}

package main

import "fmt"

func main() {
	// client side data
	income := 1000
	loanAmount := 1000
	loanTerm := 24
	// bank side data
	creditScore := 500
	// by default, should be changed if there some violation to be approved
	approved := true

	// ToDo: implement it in the function ...
	var monthlyPayment, rate, totalCost float64

	if loanTerm%12 == 0 {
		if creditScore >= 450 {
			rate = 0.15
		} else {
			rate = 0.2
		}

		totalCost = rate * float64(loanAmount)
		monthlyPayment = (totalCost + float64(loanAmount))/float64(loanTerm)

		if creditScore < 450 && 0.1 * float64(income) < monthlyPayment {
			approved = false
		}

		fmt.Printf("Application X\n" +
			"-----------\n"+
			"Credit Score: %v\n"+
			"Income: %v\n"+
			"Loan Amount: %v\n"+
			"Loan Term: %v\n"+
			"Monthly Payment: %v\n"+
			"Rate: %v\n"+
			"Total Cost: %v\n"+
			"Approved: %v\n",
			creditScore,
			income,
			loanAmount,
			loanTerm,
			monthlyPayment,
			int(rate*100),
			totalCost,
			approved)
	} else {
		fmt.Println("Opps! The loan should be divisible by 12")
	}
}

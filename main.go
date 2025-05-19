package main

import "fmt"

func  main()  { 
	transactions := []float64{}
	for{
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions= append(transactions, transaction)		
	}

	balance := calculateBalance(transactions)

	fmt.Printf("Ваш баланс: %.2f\n", balance)
	
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Println("Введите транзакцию: ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	sum := 0.0
	
	for _, value := range transactions {
		sum += value
	}

	return sum
}
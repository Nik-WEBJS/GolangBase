package main

import (
	"fmt"
	"math"
)

func main() { 
	fmt.Println("Калькулятор инлекса массы тела")
	userHeight, userKg := getUserInput()
	IMT := calculateIMT(userHeight, userKg)
	outputResult(IMT)
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Print(result)
}

func calculateIMT(userHeight, userKg float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight / 100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Print("Введите ваш рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес в килограммах: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}
package main

import (
	"errors"
	"fmt"
	"math"
)

func main() { 
	for{
		fmt.Println("Калькулятор индекса массы тела")
		userHeight, userKg := getUserInput()
		IMT, err := calculateIMT(userHeight, userKg)
		if err != nil {
			fmt.Println("Ошибка, неверно указан вес или рост")
			continue
		}
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
		switch {
	case IMT < 16:
		fmt.Println("У Вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У Вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У Вас нормальная масса тела")
	case IMT < 30:
		fmt.Println("У Вас избыточная масса тела")
	case IMT < 35:
		fmt.Println("У Вас ожирение первой степени")
	case  IMT < 40:
		fmt.Println("У Вас ожирение второй степени")
	default: fmt.Println("У Вас ожирение третьей степени")
	}
}

func calculateIMT(userHeight, userKg float64) (float64, error) {
	const IMTPower = 2
	if userKg <= 0 || userHeight <=0 {
		return 0, errors.New("ERROR")
	}
	IMT := userKg / math.Pow(userHeight / 100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Print("Введите ваш рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите ваш вес в килограммах: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Println(("Вы хотите сделать еще расчёт (y/n): "))
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
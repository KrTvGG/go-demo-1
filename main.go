package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("__ Калькулятор индекса массы тела __")
	var userHeight, userKg float64
	userHeight = takePrompt("Укажите свой рост в сантиметрах")
	userKg = takePrompt("Укажите свой вес в кг")
	IMT := userKg / math.Pow(userHeight / 100, IMTPower)
	outputResult(IMT)
}

func outputResult(IMT float64) {
	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
}

func takePrompt(text string) float64 {
	fmt.Print(text + ": ")
	var res float64
	fmt.Scan(&res)
	return res
}

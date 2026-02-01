package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Print("__ Калькулятор индекса массы тела __\n")
	var userHeight, userKg float64
	userHeight = takePrompt("Укажите свой рост в метрах")
	userKg = takePrompt("Укажите свой вес в кг")
	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print("Ваш индекс массы тела: ")
	fmt.Println(IMT)
}

func takePrompt(text string) float64 {
	fmt.Print(text + ": ")
	var res float64
	fmt.Scan(&res)
	return res
}

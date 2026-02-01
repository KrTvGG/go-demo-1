package main

import (
	"fmt"
	"math"
)

/** Для возведения в степень */
const IMTPower = 2
/** Делитель высоты, чтобы из сантиметров получить метры */
const userHeightDivider = 100

func main() {
	fmt.Println("__ Калькулятор индекса массы тела __")
	var userHeight, userKg float64
	userHeight = takePrompt("Укажите свой рост в сантиметрах")
	userKg = takePrompt("Укажите свой вес в кг")
	IMT := calculateIMT(userHeight, userKg)
	outputResult(IMT)
}

/** Выводит результат */
func outputResult(IMT float64) {
	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
}

/** Расчёт индекса массы тела */
func calculateIMT(userHeight, userKg float64) float64 {
	IMT := userKg / math.Pow(userHeight / userHeightDivider, IMTPower)
	return IMT
}

/** Выводит получаемый текст и сканирует результат из консоли */
func takePrompt(text string) float64 {
	fmt.Print(text + ": ")
	var res float64
	fmt.Scan(&res)
	return res
}

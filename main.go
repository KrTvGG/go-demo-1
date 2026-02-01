package main

import (
	"errors"
	"fmt"
	"math"
)

/** Для возведения в степень */
const IMTPower = 2
/** Делитель высоты, чтобы из сантиметров получить метры */
const userHeightDivider = 100

func main() {
	defer func() {
		r := recover()
		if recover() != nil {
			fmt.Println("Recover ", r)
		}
	}()
	fmt.Println("__ Калькулятор индекса массы тела __")
	for {
		var userHeight, userKg float64
		userHeight = takePrompt("Укажите свой рост в сантиметрах")
		userKg = takePrompt("Укажите свой вес в кг")
		IMT, err := calculateIMT(userHeight, userKg)
		if (err != nil) {
			panic("Не заданы параметры для расчёта!")
		}
		outputResult(IMT)
		outputConclusion(IMT)
		isRepeat := checkRepeatCalcelate()
		if (!isRepeat) {
			break
		}
	}
}

/** Повторять ли расчёт */
func checkRepeatCalcelate() bool {
	fmt.Println("Вы хотите повторить? (y / n)")
	var repeat string
	fmt.Scan(&repeat)
	if (repeat == "y" || repeat == "Y") {
		return true
	}
	return false
}

/** Вывод заключения по индексу массы тела */
func outputConclusion(IMT float64) {
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степено ожирения")
	}
}

/** Выводит результат */
func outputResult(IMT float64) {
	fmt.Printf("Ваш индекс массы тела: %.0f\n", IMT)
}

/** Расчёт индекса массы тела */
func calculateIMT(userHeight, userKg float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userKg / math.Pow(userHeight / userHeightDivider, IMTPower)
	return IMT, nil
}

/** Выводит получаемый текст и сканирует результат из консоли */
func takePrompt(text string) float64 {
	fmt.Print(text + ": ")
	var res float64
	fmt.Scan(&res)
	return res
}

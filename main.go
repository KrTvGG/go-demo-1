package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight, userKg float64 = 1.8, 100
	IMT := float64(userKg) / math.Pow(userHeight, 2)
	fmt.Println(IMT)
}

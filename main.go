package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	var userHeight, userKg float64 = 1.8, 100
	IMT := float64(userKg) / math.Pow(userHeight, IMTPower)
	fmt.Println(IMT)
}

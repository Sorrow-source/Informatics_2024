package lab4

import (
	"fmt"
	"math"
)

func ColculateFunction(a, b, x float64) float64 {
	if x == 0 {
		return 0
	}
	y := math.Pow(a+b*x, 2.5) / (1 + math.Log10(a+b*x))
	return y
}

func completeTaskA(a, b, xn, xk, xd float64) []float64 {
	var result []float64
	for i := xn; i < xk; i += xd {
		result = append(result, ColculateFunction(a, b, i))
	}
	return result
}

func completeTaskB(a, b float64, x []float64) []float64 {
	var result []float64
	for _, i := range x {
		result = append(result, ColculateFunction(a, b, i))
	}
	return result
}

func CompleteLab4() {
	var a float64 = 2.5
	var b float64 = 4.6
	var x []float64 = []float64{1.2, 1.28, 1.36, 1.46, 2.35}
	var xn float64 = 1.1
	var xk float64 = 3.6
	var xd float64 = 0.5
	var resultA []float64 = completeTaskA(a, b, xn, xk, xd)
	fmt.Println(resultA)
	var resultB []float64 = completeTaskB(a, b, x)
	fmt.Println(resultB)
}

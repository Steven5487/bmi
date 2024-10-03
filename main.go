package main

import "fmt"

var name string
var age int = 18
var job = "enginner"

var height, weight float64 = 171, 52
var (
	sport string
	times int
)

func main() {
	fmt.Println("hello, world!")

	name = "man"
	fmt.Println(name, age, job)
	school := "NKUST"
	fmt.Println(school)

	x := 10
	x++
	y := 5
	y--

	bmi := calculate(height, weight)
	fmt.Printf("The BMI of %s is: %.2f\n", name, bmi)
}

func calculate(height, weight float64) float64 {
	heightM := height / 100
	bmi := weight / (heightM * heightM)
	return bmi
}

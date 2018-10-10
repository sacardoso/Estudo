package main

import (
	"fmt"
)

//soma dos multiplos
//pega o numero 20
//ve quais sao os multiplos de 3 e 5
//3, 5, 6, 9, 10, 12, 15, 18
//soma esses numeros
// = 78

func main() {

	SumOfMultiples(3, 5, 20)
}

func SumOfMultiples(number1, number2, limit int) int {

	listOfMults := Multiples(number1, number2, limit)

	sum := 0
	for _, n := range listOfMults {
		sum += n
	}
	fmt.Println("A soma dos multiplos de", number1, " e ", number2, "(até ", limit, ") é:", sum)
	return sum
}

func Multiples(number1, number2, limit int) []int {

	mults := []int{}

	for i := 1; i <= (limit / number1); i++ {
		m1 := number1 * i
		if m1 < limit {
			mults = append(mults, m1)
		}
	}

	for i := 1; i <= (limit / number2); i++ {
		fmt.Println("The multiples:", mults)
		m2 := number2 * i
		if m2 < limit {
			for _, v := range mults {
				if m2 == v {
					m2 = 0
					//mult2 := &m2
					//*mult2 = 0
				}
			}
			mults = append(mults, m2)
		}
	}
	fmt.Println(mults)
	return mults
}

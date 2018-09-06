package main

import "fmt"

//number = numero de termos!
//numbers = lista com "number" termos

//MakeList cria uma lista de numeros com N(number) termos
func MakeList(number int) []int {
	result := make([]int, number)
	for i := 1; i <= number; i++ {
		result[i-1] = i
	}
	return result

}

//SquareList cria a lista do quadrado dos numeros(numbers = MakeList)
func SquareList(numbers []int) []int {
	result := make([]int, len(numbers))
	for i, n := range numbers {
		result[i] = n * n
	}
	return result
}

//SumList soma os numeros da MakeList
func SumList(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

//Process - Junta as funçoes, faz todo o processo
func Process(number int) int {
	numbers := MakeList(number)
	sum := SumList(numbers)
	squares := SquareList(numbers)

	sqrSum := sum * sum
	sumSqr := SumList(squares)

	dif := sqrSum - sumSqr

	return dif

}

func main() {
	result := Process(10)
	fmt.Println(result)
}

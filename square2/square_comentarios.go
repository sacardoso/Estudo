package main

import (
	"fmt"
)

// The MakeIntList() function returns an array of consecutive integers
// starting from 1 all the way to the `number` (including the number)
func MakeIntList(number int) []int { // cria uma lista com N(N=number) numeros (escolhido no process(N))
	result := make([]int, number)
	for i := 1; i <= number; i++ { // result[1-1=0] = 1 ; result[2-1=1] = 2 ; result[3-1=2] = 3 ......
		result[i-1] = i // faz isso pq nao quer que comece com 0, e sim com 1
	}
	return result
}

// The SquareList() function takes a slice of integers and returns an
// array of the quares of these integers
func SquareList(numbers []int) []int { // numbers = lista dos numeros criada na MakeIntList
	result := make([]int, len(numbers))

	for i, n := range numbers { // i --> relativo a posicao; n --> percorre o "range"dos numbers...  pega a lista(numbers) e percorre ela toda... se fizesse o for igual ao interior, faria return[i] = num[i]*num[i]
		result[i] = n * n // result[1]=1*1=1; result[2]=2*2=4; result[3]=3*3=9 .....
	}

	return result // mostra essa lista nova, com os quadrados
}

// The SumList() function takes a slice of integers and returns their sum
func SumList(numbers []int) int { // numbers=lista criada no MakeIntList
	result := 0                 // agora o resultado nao é mais uma lista, so um numero.. que vai sendo incrementado! por isso =0.
	for _, n := range numbers { // percorre os numeros para somar, mas nao quer a posicao, e sim o numero em si, por isso "tira"o i
		result += n // soma 1+2+3+4+5+6+7+8+......
	}

	return result
}

// Solve Project Euler #6 - Sum square difference
func Process(number int) int { // number=numero de termos que queremos
	numbers := MakeIntList(number) // variavel numbers recebe a lista
	sum := SumList(numbers)        // variavel sum recebe a soma dos numeros(numbers)
	squares := SquareList(numbers) //squares recebe a lista dos quadrados nos numbers

	sumOfSquares := SumList(squares) // ao inves de somar os numbers, soma os quadrados --> soma dos quadrados (1+4+9+16+25+....)
	squareOfSum := sum * sum         // quadrado da soma --> pega a soma dos numeros, e multiplica por ela mesma: quadrado da soma

	diff := squareOfSum - sumOfSquares // faz a diferenca

	return diff // retorna o resultado!
}

func main() {
	result := Process(10)
	fmt.Println(result) // printa o resultado do result, que é o diff
}

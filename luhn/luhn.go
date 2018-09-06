package main

import "fmt"

//para mostrar naquele "formato de cartao":
//s := []string{"4539", "1488", "0343", "6467"}
//	fmt.Println(strings.Join(s, " "))

func main() {

	number := []int{4, 5, 3, 9, 1, 4, 8, 8, 0, 3, 4, 3, 6, 4, 6, 7}

	newNumber := DoubleNumber(number)
	sum := SumDigits(newNumber)

	fmt.Println(VerifyIfNumberIsValid(sum))

}

//DoubleNumber dobra o numero com indices 0, 2, 4, 6, 8 .......
func DoubleNumber(number []int) []int {

	for i, n := range number { // 'n' vai ser os numeros do number(n[i]); exemplo: n[0]=4; n[1]=5.....

		if i%2 == 0 {

			number[i] = n * 2

			if number[i] > 9 {
				result := number[i] - 9
				number[i] = result
			}
		}
	}

	fmt.Println(number)
	return number
}

//SumDigits vai somar os numeros do NOVO slice number
func SumDigits(newNumber []int) int {

	total := 0
	for _, v := range newNumber {
		total += v
	}
	fmt.Println(total)
	return total
}

//VerifyIfNumberIsValid verifica se o numero é valido
func VerifyIfNumberIsValid(sum int) string {

	if sum%10 == 0 {
		return "The number is valid!"
	}

	return "The number is not valid!"

}

/*
	//number := []int{4, 5, 3, 9, 1, 4, 8, 8, 0, 3, 4, 3, 6, 4, 6, 7}

	//N0 := number[0]
	dupN0 := number[0] * 2

	if dupN0 > 9 {
		result := dupN0 - 9
		dupN0 = result
	}

	//N2 := number[2]
	dupN2 := number[2] * 2

	if dupN2 > 9 {
		result := dupN0 - 9
		dupN0 = result
	}

	//N4 := number[4]
	dupN4 := number[4] * 2

	if dupN4 > 9 {
		result := dupN4 - 9
		dupN4 = result
	}

	//N6 := number[6]
	dupN6 := number[6] * 2

	if dupN6 > 9 {
		result := dupN6 - 9
		dupN6 = result
	}

	//N8 := number[8]
	dupN8 := number[8] * 2

	if dupN8 > 9 {
		result := dupN8 - 9
		dupN8 = result
	}

	//N10 := number[10]
	dupN10 := number[10] * 2

	if dupN10 > 9 {
		result := dupN10 - 9
		dupN10 = result
	}

	//N12 := number[12]
	dupN12 := number[12] * 2

	if dupN12 > 9 {
		result := dupN12 - 9
		dupN12 = result
	}

	//N14 := number[14]
	dupN14 := number[14] * 2

	if dupN14 > 9 {
		result := dupN14 - 9
		dupN14 = result

	}

	fmt.Println(dupN0, "_", dupN2, "_", dupN4, "_", dupN6, "_", dupN8, "_", dupN10, "_", dupN12, "_", dupN14, "_")

	fmt.Println(dupN0, number[1], dupN2, number[3], dupN4, number[5], dupN6, number[7], dupN8, number[9], dupN10, number[11], dupN12, number[13], dupN14, number[15])
*/

/*func DoubleNumber([]int) []int{




}
*/

//Validating a number
//strings of length 1 or less are not valid
//spaces só no input, mas devem ser retirados antes de checar
//nenhum outro non-digit é permitido

/*
example:

4539 1488 0343 6467

1o passo: dobrar o segundo digito(comecando da direita)

4_3_ 1_8_ 0_4_ 6_6_   -- numeros para serem dobrados!!!!!

se qnd dobrarmos o numero, ele for maior que 9... diminui 9 do produto.

8_6_ 2_7_ 0_8_ 3_3_

resultado:

8569 2478 0383 3437

soma todos os numeros

8+5+6+9+......3+4+3+7 = 80

se o resultado for divisivel por 10, o numero é valido! (if sum%10 = 0 {fmt.Println("The number is valid")})

*/

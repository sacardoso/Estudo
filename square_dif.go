package main

import "fmt"

func main() {

	//N := 9
	//num := make([]int, N)
	//sqr := make([]int, N)
	var num1, num2 int

	fmt.Println("Digite os números:")

	fmt.Scan(&num1, &num2)

	//for i := 0; i < N; i++ {
	//	fmt.Scan(&num[i])
	//}

	//for i := 0; i < N; i++ {
	//	sqr[i] = num[i] * num[i]
	//	fmt.Print("O quadrado do número:", sqr[i])
	//}

	sqr1 := num1 * num1
	fmt.Println("O quadrado do número", num1, "é", sqr1)

	sqr2 := num2 * num2
	fmt.Println("O quadrado do número", num2, "é", sqr2)

	sumNum := num1 + num2
	sqrSum := sumNum * sumNum //quadrado da soma

	fmt.Println("O quadrado da soma dos números é:", sqrSum)

	sumSqr := sqr1 + sqr2 // soma dos quadrados

	fmt.Println("A soma dos quadrados é:", sumSqr)

	dif := sqrSum - sumSqr

	fmt.Println("A diferença do quadrado da soma e a soma dos quadrados é:", dif)

}

/*func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func soma(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
*/

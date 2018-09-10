package main

import (
	"fmt"
	"strings"
)

func main() {

	A := "1, 2, 3, 4"
	B := "1, 2, 3"

	fmt.Print(Sublist(A, B))
}

//Sublist verifica
func Sublist(A, B string) string {

	if (strings.Contains(B, A)) == true {
		if A == B {
			return fmt.Sprintln("A is equal to B")
		}
		return fmt.Sprintln("A is a sublist of B")
	}

	if (strings.Contains(A, B)) == true {
		if len(A)-2 > len(B) {
			return fmt.Sprintln("A is a superlist of B")
		}
	}

	return "A is not a superlist of, sublist of or equal to B"
}

/*
//Sublist verifica as listas e informa o que ela é
func Sublist(string, string) string {

	//A := "1, 2, 3, 4"
	//B := "1, 2, 3, 4"
	var A, B string

	if (strings.Contains(B, A)) == true {
		if A == B {
			return "A is equal to B"
		}
		return "A is a sublist of B"
	}

	if (strings.Contains(A, B)) == true {
		if len(A) > len(B) {
			return "A is a superlist of B"
		}
	}
	return "A is not a superlist of, sublist of or equal to B"
}
*/

/*

//Equal verifica se são iguais
func Equal(string) string {

	if (strings.Contains(B, A)) == true {
		if A == B {
			fmt.Println("A is equal to B")
			return
		}
	}
}
//Sublist mostra se e uma superlist
func Sublist(string) string {

	if (strings.Contains(B, A)) == true && A != B{
		fmt.Println("A is a sublist of B")
	}
}

func SuperList(string) string{

	if (strings.Contains(A, B)) == true {
		if len(A) > len(B) {
			fmt.Println("A is a superlist of B")
			return
		}
	}
}

*/

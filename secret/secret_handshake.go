package main

import (
	"fmt"
)

func main() {

	Secret(19)
}

func Secret(number int64) []string {

	handshake := []string{}
	numero := number

	if number >= 16 {
		number = number - 16
	}

	if number > 8 || number-8 == 0 {
		handshake = append(handshake, "jump")
		number = number - 8
	}

	if number > 4 || number-4 == 0 {
		handshake = append(handshake, "close your eyes")
		number = number - 4
	}

	if number > 2 || number-2 == 0 {
		handshake = append(handshake, "double blink")
		number = number - 2
	}

	if number > 1 || number-1 == 0 {
		handshake = append(handshake, "wink")
		number = number - 1
	}

	if numero < 16 {
		Reverse(handshake)
	}

	fmt.Println(handshake)
	return handshake
}

func Reverse(handshake []string) []string {
	chars := []string(handshake)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 { //inverter a posicao
		chars[i], chars[j] = chars[j], chars[i]
	}
	return chars
}

/*1 = wink = 1
10 = double blink = 2
100 = close your eyes = 4
1000 = jump = 8


10000 = Reverse the order of the operations in the secret handshake. = 16

*/

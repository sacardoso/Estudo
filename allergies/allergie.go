package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	allergie(515)
}

var typesOfAllergies = map[string]float64{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

//IntToBytes convertendo o numero para bytes e mostra na ordem que eu quero!!!!
func allergie(allergies int64) string {

	b := strconv.FormatInt(allergies, 2)  //converti o numero para binario, para saber as potencias de 2
	numberInBytes := strings.Split(b, "") //converte o numero para [], para conseguir pegar a posicao
	chars := []string(numberInBytes)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 { //inverter a posicao
		chars[i], chars[j] = chars[j], chars[i]
	}

	allergic := ""
	for i := range numberInBytes {

		if numberInBytes[i] == "1" {
			m := math.Pow(2, float64(i))

			for t, number := range typesOfAllergies {
				if m == number {
					allergic += t + " "
				}
			}
		}
	}
	fmt.Println(allergic)
	return allergic
}

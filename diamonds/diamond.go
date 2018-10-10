package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {

	fmt.Println(Diamond('G'))
}

func Diamond(letter rune) string {
	result := ConvertToByte(letter)
	return fmt.Sprintln(Gen(result))
}

func ConvertToByte(letter rune) byte {
	char := letter
	limit := byte(char)
	return limit
}

func Gen(limit byte) (string, error) {
	if limit < byte('A') || byte('Z') < limit { // menor que 65 e maior que 90 retorna esse erro
		return "", errors.New("Not a valid limit: " + string(limit))
	}
	length := int(limit - byte('A'))     // tamanho = 'tamanho' da letra - 65
	result := make([]string, 2*length+1) // result = []string com tamanho 2*tamanho + 1 --> pq sempre duplica as linhas debaixo, e conta +1 por causa do A
	for r := 0; r <= length; r++ {
		row := make([]rune, 2*length+1)
		for col, _ := range row {
			letter := '-'
			if length+r == col || length-r == col {
				letter = rune('A' + r) //65 + r
			}
			row[col] = letter
		}
		result[r] = string(row) // para repetir as linhas
		result[2*length-r] = string(row)
	}
	return strings.Join(result, "\n") + "\n", nil
}

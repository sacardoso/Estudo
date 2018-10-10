package main

import "fmt"

// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

func main() {

	s := person{name: "Sean", age: 50} // s está alocado no end de memoria 200

	sp := &s            // end de memoria do sp vai apontar para o end de memoria do s, que é 200
	fmt.Println(sp)     // sp = s = {Sean, 50}
	fmt.Println(sp.age) // sp.age = s.age = 50

	sp.age = 51     // como o end de memoria do sp.age está alocado no s.age,
	fmt.Println(s)  // qnd ele printar o s, ele vai estar modificado: {sean, 51}
	fmt.Println(sp) // sp tambem se modificará

	//MAS
	//Se nao tivesse o & no sp := &s

	/*
			sp := s // end de memoria do sp vai ser DIFERENTE do end de memoria do s
			ele só está copiando o valor de s! s nao vai ser modificado caso o sp mude!
		    fmt.Println(sp) // sp = s = {Sean, 50}
			fmt.Println(sp.age) // sp.age = s.age = 50

		    sp.age = 51 // como o end de memoria do sp só recebeu o 51, SOMENTE o sp.age será modificado.
			fmt.Println(s) // qnd ele printar o s, ele não vai estar modificado: {sean, 50}!!!
			isso pq qnd ele modifica o sp.age, ele n tem nenhuma relacao com o s.age, pois está em outro end de memoria
			entao o codigo lê o valor de s.age, q n foi modificado pq o sp nao aponta para o s!!!
			fmt.Println(sp) // sp está se modificará: {sean, 51}
	*/
}

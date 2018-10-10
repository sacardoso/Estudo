package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x antes", &x)  // endereço de memoria do x
	fmt.Println("x antes", x)   // valor do x, que inicialmente é = 0
	foo(&x)                     // depois que chama a funcao, o valor de x muda pq ele aponta para o valor de y, que na função muda para 43
	fmt.Println("x depois", &x) // mesma memoria de sempre
	fmt.Println("x depois", x)  // x agora tem o valor de y, que é 43
}

func foo(y *int) {
	fmt.Println("y antes", y)   // y = x = 0
	fmt.Println("y antes", *y)  // &y = &x --> memoria do y é a mesma do x, pq aponta para ele!!!!!
	*y = 43                     // quando Y recebe um valor, o x tambem vai receber, pois ele aponta para a memoria do x
	fmt.Println("y depois", y)  // agora y = 43 e x = 43
	fmt.Println("y depois", *y) // segue na mesma memoria
}

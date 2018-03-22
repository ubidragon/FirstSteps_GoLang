package main

import (
	"fmt"
)

func main() {
	suma(4, 4)
	sumaRecortada(5, 5)
	fmt.Println(sumaRetorno(2, 3))
}

func suma(x int, y int) {
	fmt.Println(x + y)
}

func sumaRecortada(x, y int) {
	fmt.Println(x + y)
}

func sumaRetorno(x int, y int) int {
	return x + y
}

package main

import "fmt"

// Obtém a série de potências do inteiro a e retorna a tupla do quadrado de
// e cubo de a. (aˆ2, aˆ3)
//func powerSeries(a int) (int, int) {
//	return a * a, a * a * a
//}

func powerSeries(a int) (square int, cube int) {
	square = a * a
	cube = square * a
	return
}

func main() {

	var square int
	var cube int
	square, cube = powerSeries(3)
	fmt.Println("quadrado: ", square, "Cubo: ", cube)

}

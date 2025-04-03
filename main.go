package main

import "fmt"

func main(){
	// slice
var numeros = []int{1, 2, 3, 4, 5}
fmt.Println(numeros)
numeros = append(numeros, 7, 8, 9)
fmt.Println(numeros)
}
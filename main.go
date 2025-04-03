package main

import "fmt"

func main(){
	var nome = []string{"murilo","nakaguma", "manu", "lucas", "gobetti"}
	fmt.Println(nome)
	rangeOne := nome[:2]
	fmt.Println(rangeOne)
	rangeTwo := nome[3:]
	fmt.Println(rangeTwo)
	fmt.Println(nome[2])
}
package main

import (
	"fmt"
)

func main()  {
	card:= newCard()
	fmt.Println(card)
}

/*
para las funciones se requiere seamos muy explicitos para los datos que se regreseran
*/
//newCard is a function return a string
func newCard()string  {
	return "5 de diamantes"
}
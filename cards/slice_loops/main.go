package main
/*
Aquí se podra ver dos diferentes tipos de estructuras basicas de Go
Array y Slice
un array tiene un tamaño determinado
pero un slice es un array que puede crecer o reducirce a voluntad
*/
import (
	"fmt"
)

func main()  {
	card:= []string{newCard(),"As de diamantes"}
	card= append(card,"seis de espadas")
	fmt.Println(card)
} 

func newCard()string  {
	return "5 de diamantes"
}
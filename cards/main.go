package main

import (
	"fmt"
)

func main()  {
	card:= mazo{newCard(),"As de diamantes"}
	card= append(card,"seis de espadas")
	for i, carta := range card{
		fmt.Println(i,"-->", carta)
	}
} 

func newCard()string  {
	return "5 de diamantes"
}
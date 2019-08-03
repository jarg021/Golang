package main

import(
	"fmt"
)

func main()  {
	//designando tipo de variable 
	var card string="as de espadas"
	//De esta forma confiamos en el compilador para descubrir el tipo de dato de esta forma solo se utiliza cuando se declara la variable
	card2:= "as de copas"
	fmt.Println(card, card2)
	card= "5 de diamantes"
	fmt.Println(card)
}
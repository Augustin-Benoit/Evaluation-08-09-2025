package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("7")
	fmt.Scan(&x)
	fmt.Print("63")
	fmt.Scan(&y)

	resultat := PGCD(x, y)
	fmt.Printf("Le plus grand diviseur commun de %d et %d est : %d\n", x, y, resultat)
}

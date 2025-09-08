package main

import (
	"fmt"
)

func CompterLettres(s string) int {
	compteur := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			compteur++
		}
	}
	return compteur
}

func main() {
	chaine := "Bonjour le monde ! 123 😊 Ça va ?"
	fmt.Println("Nombre de lettres latines :", CompterLettres(chaine))
}

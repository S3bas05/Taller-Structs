package main

import (
	"fmt"
	"strings"
)

func ContarPalabras(texto string) {
	palabras := strings.Fields(strings.ToLower(texto))
	conteo := make(map[string]int)

	for _, palabra := range palabras {
		conteo[palabra]++
	}

	fmt.Println("\nConteo de palabras:")

	mayor := 0
	palabraMayor := ""

	for palabra, cantidad := range conteo {
		fmt.Printf("%s: %d veces\n", palabra, cantidad)

		if cantidad > mayor {
			mayor = cantidad
			palabraMayor = palabra
		}
	}

	if len(conteo) == 0 {
		fmt.Println("No se ingresaron palabras.")
	} else {
		fmt.Printf("\nPalabra mas repetida: %s\n", palabraMayor)
		fmt.Printf("Cantidad de repeticiones: %d\n", mayor)
	}
}

func main() {
	var texto string

	fmt.Println("Ingrese una frase o texto:")
	fmt.Scanln(&texto)

	ContarPalabras(texto)
}

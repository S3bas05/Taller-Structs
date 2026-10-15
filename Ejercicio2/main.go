package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func MostrarProducto(p Producto) {
	fmt.Println("Nombre:", p.Nombre)
	fmt.Printf("Precio: $%.2f\n", p.Precio)
	fmt.Println("Cantidad:", p.Cantidad)
}

func CalcularTotal(p Producto) float64 {
	return p.Precio * float64(p.Cantidad)
}

func TieneStock(p Producto) bool {
	return p.Cantidad > 0
}

func main() {
	productos := []Producto{
		{"Arroz", 1.25, 10},
		{"Leche", 0.95, 5},
		{"Pan", 0.50, 0},
	}

	for _, producto := range productos {
		fmt.Println("---------------------")

		MostrarProducto(producto)

		total := CalcularTotal(producto)
		fmt.Printf("Valor total: $%.2f\n", total)

		if TieneStock(producto) {
			fmt.Println("Stock: Disponible")
		} else {
			fmt.Println("Stock: Sin stock")
		}
	}
}

package GoArea

import (
	"math"
)

// Pi e uma constante numerica
const Pi = 3.1416

// Circ calcula a area da circuferencia
func Circ(raio float64) float64{
	return math.Pow(raio,2) * Pi
}

//Rect calcula a area de um retangulo
func Rect(base, altura float64) float64{
	return base * altura
}

//Nao e visivel
func _TrianguloEq(base, altura float64) float64{
	return (base * altura) / 2
}

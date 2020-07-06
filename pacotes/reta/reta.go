package main

import "math"

// Inicializando com letra maiuscula e PUBLICO (visibilidade e fora do pacote)!
// Inicializando com letra minuscula e PRIVADO (visibilidade apenas dentro do pacote)!

// Por Exemplo
// Ponto - gerara algo publico
// ponto ou _Ponto - gerara algo privado

// Ponto representa uma coordenada no ponto cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// A distancia e responsavel por calcular a distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Teste de ponto dentro de polígono convexo
func insidePolygon(p Point, poly []Point) bool {
	n := len(poly)

	for i := 0; i < n; i++ {
		a := poly[i]
		b := poly[(i+1)%n]

		cross := (b.X-a.X)*(p.Y-a.Y) - (b.Y-a.Y)*(p.X-a.X)

		if cross < -1e-9 {
			return false
		}
	}

	return true
}

func main() {
	// Pentágono regular de lado aproximadamente 1
	r := 1.0

	var pent []Point

	// Gera vértices no sentido anti-horário
	for i := 0; i < 5; i++ {
		ang := math.Pi/2 + 2*math.Pi*float64(i)/5
		pent = append(pent, Point{
			X: r * math.Cos(ang),
			Y: r * math.Sin(ang),
		})
	}

	// Escolhe uma aresta do pentágono
	a := pent[0]
	b := pent[1]

	dx := b.X - a.X
	dy := b.Y - a.Y

	edgeLen := math.Hypot(dx, dy)

	ux := dx / edgeLen
	uy := dy / edgeLen

	// normal para dentro (ajustada para este pentágono)
	nx := -uy
	ny := ux

	bestSide := 0.0

	// busca simples
	for side := 0.0; side <= edgeLen; side += 0.0001 {

		p3 := Point{
			X: a.X + nx*side,
			Y: a.Y + ny*side,
		}

		p4 := Point{
			X: b.X + nx*side,
			Y: b.Y + ny*side,
		}

		if insidePolygon(p3, pent) &&
			insidePolygon(p4, pent) {

			bestSide = side
		}
	}

	fmt.Printf("Maior lado do quadrado: %.6f\n", bestSide)
	fmt.Printf("Área: %.6f\n", bestSide*bestSide)
}

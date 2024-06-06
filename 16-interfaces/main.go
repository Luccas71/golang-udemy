package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	base   float64
	altura float64
}

type circulo struct {
	raio float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da figura é %.2f\n", f.area())
}

func (r retangulo) area() float64 {
	return r.base * r.altura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {

	r := retangulo{10, 10}
	escreverArea(r)

	c := circulo{1}
	escreverArea(c)
}

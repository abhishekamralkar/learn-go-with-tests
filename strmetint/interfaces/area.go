package main

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func (c Circle) Area() float64 {
	return (math.Pi * c.Radius * c.Radius)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

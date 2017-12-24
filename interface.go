package main

import (
  "fmt"
  "math"
)

type geometry interface {
  area() float64
  perimeter() float64
}

type rect struct {
  width, height float64
}

type circle struct {
  radius float64
}

func (r rect) area() float64 {
  return r.width * r.height
}

func (r rect) perimeter() float64 {
  return r.width + 2 * r.height
}

func (c circle) area() float64 {
  return 2 * math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64{
  return 2 * math.Pi * c.radius
}

func calculate(g geometry) {
  fmt.Println(g)
  fmt.Println("Area:", g.area())
  fmt.Println("Circumference:", g.perimeter() )
}

func main(){
  r := rect{ width: 10, height: 7 }
  c := circle{ radius: 7 }
  calculate(r)
  calculate(c)
}

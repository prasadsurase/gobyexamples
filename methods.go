package main

import "fmt"

type rect struct{
  height int
  width int
}

type square struct {
  side int
}

func (r *rect) area() int{
  return r.height * r.width
}

func (s *square) area() int {
  return s.side * s.side
}

func main(){
  r := rect{height: 10, width: 5}
  s := square{ side: 5 }
  fmt.Println("area:", r.area())
  fmt.Println("area:", s.area())
}

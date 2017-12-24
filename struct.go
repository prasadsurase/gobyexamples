package main

import "fmt"

type User struct{
  name string
  age int
}

func main(){
  u := User{name: "Prasad", age: 31}
  fmt.Println(u)
  fmt.Println(u.name)
  fmt.Println(u.age)

  s := &u
  fmt.Println(s)
  fmt.Println(s.name)
  fmt.Println(s.age)
}

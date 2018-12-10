package main

import "fmt"

type Person struct {
  Name string
  Age  int
}

func main() {
  child := Person{}
  child.Name = "Felicity"
  child.Age = 2
  fmt.Println(child.Name, child.Age)
  
  child2 := Person{"Bethany", 1}
  fmt.Println(child2)
}

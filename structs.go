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
  
  child2 := Person{}
  child2.Name = "Bethany"
  child2.Age = 1
  fmt.Println(child2.Name, child2.Age)
}

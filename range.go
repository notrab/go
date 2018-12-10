package range

import "fmt"

func main() {
  ages := map[string]float64{"Felicity": 2, "Bethany": 1}
  
  for name, age := range ages {
    fmt.Println(name, age)
  }
}

package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

    array := [5]int{r1.Intn(25), r1.Intn(25), r1.Intn(25), r1.Intn(25), r1.Intn(25)}
   max := array[4]
   min := array[0]

fmt.Println(array)
  
for _, value := range array {
  if value < min {
    min=value
  }
}

fmt.Println("The smallest value is : ", min)

for _, value := range array{
  if value > max {
    max=value
  }
}

fmt.Println("The largest value is : ", max)

}
package main

import (
	"fmt"
	"math"
)

func main(){
  var a float64
  fmt.Scan(&a)
  if math.Pow(2,a) > math.Pow(a,2) {
    fmt.Println("Yes")
  }else{
    fmt.Println("No")
  }
}


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string{
  sc.Scan()
  return sc.Text()
}

func main(){

  var n int
  var cnt int
  var sl []int
  fmt.Scan(&n)
  inputs := strings.Split(nextLine()," ")
  sl = append(sl, 0,360)
  for _,nm := range inputs {
    cnt += nm
    sl = append(sl, cnt%360)
  }

}

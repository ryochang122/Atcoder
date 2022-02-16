package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string{
  sc.Scan()
  return sc.Text()
}

func atoi(s string) int {
  var i int
  i,_=strconv.Atoi(s)
  return i
}

func main(){

  var n int
  var cnt int
  var max int
  var sl []int

  fmt.Scan(&n)
  inputs := strings.Split(nextLine()," ")

  sl = append(sl, 0,360)
  for _,nm := range inputs {
    cnt += atoi(nm)
    cnt %= 360
    sl = append(sl, cnt)
  }

  sort.Slice(sl, func(i,j int) bool{return sl[i] > sl[j]})
  for i:=0;i<len(sl)-1;i++{
    if d := sl[i] - sl[i+1];max<d{
      max = d
    }
  }
  fmt.Println(max)
}

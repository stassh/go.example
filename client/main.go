package main


import (
  "fmt"
  "github.com/shopspring/decimal"
)

func main() {
  a, err := decimal.NewFromString("184467440073709551615")
  if err != nil {
    fmt.Println("Error %v", err)
    return;
  }
  b, _ := decimal.NewFromString("10")
  fmt.Println("a. Number A is %v", a)
  fmt.Println("b. Number B is %v", b)

  sum := decimal.Sum(a, b)

  fmt.Println("A + B is %v", sum)

}
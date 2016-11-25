package main

import (
  "fmt"
  "stringutil"
)
func main(){
  f, n := stringutil.FullName("abhishek", "gautam")
  fmt.Printf("Full Name: %v, character count: %v ", f, n)
  f1, n1 := stringutil.FullNameNakedReturn("abhishek", "gautam")
  fmt.Println("\nNaked Return:")
  fmt.Printf("Full Name: %v, character count: %v ", f1, n1)
}

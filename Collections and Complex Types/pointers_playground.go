package main

import(
  "fmt"
)

func main(){
  //a null pointer not pointing towards anything
  var p *int
  if p!=nil{
      fmt.Println("value of p:", *p)
  }else{
    fmt.Println("value of p is nil")
  }
  var v int = 42
  p = &v
  if p!=nil{
      fmt.Println("value of p:", *p)
  }else{
    fmt.Println("value of p in nil")
  }
  var number float64 = 43.45
  pointer := &number

  fmt.Println("Value of number is:", *pointer)
  *pointer += 2
  fmt.Println("Now value of number is:", *pointer)
}

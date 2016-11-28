package main

import (
  "fmt"
  "os"
  "errors"
)
func main() {
  myErrors := errors.New("FATAL! The file you are trying to access does not exist!")
  f, err := os.Open("/Users/Gautam/Desktop/phone.txt")
  if err!=nil {
    fmt.Println(myErrors)
  }else{
    fmt.Println(f)
  }

}

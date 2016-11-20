package main
import (
  "fmt"
)

func main(){
  fmt.Println("is Number equal to 5?",checkNumber(6))
  f, n := FullName("abhishek", "gautam")
  fmt.Printf("Full Name: %v, character count: %v ", f, n)
  f1, n1 := FullNameNakedReturn("abhishek", "gautam")
  fmt.Println("\nNaked Return:")
  fmt.Printf("Full Name: %v, character count: %v ", f1, n1)
}
//function with lowercase(Local to this package) returning
//a single value
func checkNumber(x int) bool {
  if x == 5{
    return true
  }else{
    return false
  }
}

//function with uppercase(not local to this package)returning more than value
func FullName(f, l string) (string, int) {
  full := f + " " + l
  length := len(full)
  return full, length
}

func FullNameNakedReturn(f, l string) (fn string, cc int) {
  fn = f + " " + l
  cc = len(fn)
  return
}

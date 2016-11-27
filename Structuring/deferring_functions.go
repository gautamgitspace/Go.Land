//Each statement prefixed by defer keyword is pushed on to a stack
//and is entertained in LIFO order.
//defer statemnt works only inside a function
package main
import (
  "fmt"
)
func main() {
  defer fmt.Println("1")
  defer fmt.Println("2")
  defer fmt.Println("3")
  defer fmt.Println("4")
  fmt.Println("Undeferred statem")
}

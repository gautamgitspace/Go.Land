package main
import "fmt"
func main(){
  //explicit typed declarations
  const decision bool = true
  var number int = 10
  var text string = "N.W.A"
  number = 11
  //implicit typed declarations
  greeting := "Hello"

  fmt.Println(decision)
  fmt.Println(number)
  fmt.Println(text)
  fmt.Println(greeting)
}

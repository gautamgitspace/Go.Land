package main
import(
  "fmt"
)
func main(){
  sum := 1
  //for loop behaves like a while loop in go
  for sum < 1000 {
    sum += sum
    fmt.Println("Sum:", sum)
    if sum > 200 {
      goto endofprogram
    }
    if sum > 500 {
      break
    }
  }
  endofprogram : fmt.Println("EOP")
}

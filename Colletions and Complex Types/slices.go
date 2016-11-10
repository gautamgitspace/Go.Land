//slices in go
package main
import (
  "fmt"
  "sort"
)
func main(){
  var colors = []string{"red", "blue", "green"}
  fmt.Println(colors)
  //add an item
  colors = append(colors, "purple")
  fmt.Println(colors)
  //remove first item
  colors = append(colors[1:])
  fmt.Println(colors)
  //remove last item
  colors = append(colors[:len(colors)-1])
  fmt.Println(colors)

  //declaring a slice with a type using make function
  // type, initial length, optional capacity
  numbers := make([]int, 5, 5)
  numbers[0] = 11
  numbers[1] = 2
  numbers[2] = 33
  numbers[3] = 14
  numbers[4] = 54
  fmt.Println(numbers)
  numbers = append(numbers, 555)
  //going past the capacity so more cap will be added in the background
  fmt.Println(numbers)
  fmt.Println(cap(numbers))
  sort.Ints(numbers)
  fmt.Println("ASC ORDER:", numbers)
  sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
  fmt.Println("DSC ORDER:", numbers)
}

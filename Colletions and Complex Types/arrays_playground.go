package main
import(
  "fmt"
)
func main(){
  var colors [4]string
  colors[0] = "Red Sox"
  colors[1] = "Mets"
  colors[2] = "Yankees"
  colors[3] = "Royals"

  fmt.Println(colors)
  fmt.Println(colors[1])
  fmt.Println(len(colors))
}

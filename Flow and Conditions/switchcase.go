//switch case doesn't need a break statement in Go
package main
import(
  "fmt"
  "math/rand"
  "time"
)
func main(){
  rand.Seed(time.Now().Unix())
  //dow := rand.Intn(6) + 1

  result := ""
  //this will restrict dow's scope to this block only
  switch dow := rand.Intn(6) + 1; dow{
    case 7:
      result = "Sunday"
    case 6:
      result = "Saturday"
    default:
      result = "Weekday"
  }
  fmt.Println(result)
}

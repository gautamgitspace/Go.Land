package main
import (
  "fmt"
)
func main() {
  headcount := map[string]bool {
  "Carmelo" : true,
  "Derrick" : true,
  "Dirk" : true,
  "Rajon" : true}
  //if there exists an entry, 'ok' will be true and isPresent will
  //contain the value associated with the key
isPresent, ok := headcount["Dirk"]
if ok {
  fmt.Println("Was Dirk Present?", isPresent)
}else{
  fmt.Println("No info found for Paul")
}
}

package main
import (
  "fmt"
  "sort"
)
func main(){
  states := make(map[string]string)
  /*ADD TO MAP*/
  states["NY"] = "New York"
  states["PA"] = "Pennsylvania"
  states["OH"] = "Ohio"
  states["OR"] = "Oregon"
  fmt.Println(states)

  //fetch from map
  Penn := states["PA"]
  //delete from map
  fmt.Println("fetch:", Penn)
  delete(states, "OH")
  fmt.Println(states)
  //iterate through a map
  for k, v := range states {
    fmt.Printf("%v: %v\n", k, v)
  }

  /*ordered iteration*/

  //make a slice to hold keys
  keys := make([]string, len(states))
  i := 0
  for k := range states {
    keys[i] = k
    i++
  }
  //sort the slice(alphabetical order) that holds keys
  sort.Strings(keys)
  for i:= range keys {
    fmt.Println(states[keys[i]])
  }

}

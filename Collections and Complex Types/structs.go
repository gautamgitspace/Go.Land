//Structs encapsulate both data and methods but Go doesn't have an inheritance model. There is no concept
//of super or sub-structs. Each struct is independent, with its own fields for data management and optionally
//its own methods.
package main
import (
  "fmt"
)
type Player struct {
  Team string
  Name string
  JerseyNumber int
}
func main(){
  Melo := Player{"New York Knicks", "Carmelo Anthony", 7}
  Mamba := Player{"LA Lakers", "Kobe Bryant", 24}

  fmt.Println(Melo)

  fmt.Printf("%+v\n", Melo)

  fmt.Printf("Name: %v\nTeam: %v\nJersey# %v\n", Melo.Name, Melo.Team, Melo.JerseyNumber)
  fmt.Printf("Name: %v\nTeam: %v\nJersey# %v\n", Mamba.Name, Mamba.Team, Mamba.JerseyNumber)
}

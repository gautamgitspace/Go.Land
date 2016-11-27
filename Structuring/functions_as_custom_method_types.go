//function as custom types, call by val and call by ref
package main

import (
  "fmt"
)

type Player struct{
  Team string
  Name string
  JerseyNumber int
}

func (p Player) displayJerseyNumber() {
  fmt.Println(p.JerseyNumber)
}

func (p Player) changeJerseyNumberByValue() {
  p.JerseyNumber = 34
  fmt.Println("Changing by Value:", p.JerseyNumber)
}

func (p *Player) changeJerseyNumberByRef() {
  p.JerseyNumber = 34
  fmt.Println("Changing by Ref:", p.JerseyNumber)
}
func main() {
  obj := Player{"New York Knicks", "Carmello Anthony", 7}
  //dump of the object
  fmt.Println("Before Name change:", obj)
  obj.displayJerseyNumber()

  obj.Name = "Melo"
  fmt.Println("After Name change:", obj)

  obj.changeJerseyNumberByValue()
  fmt.Println("After Change by value, original object doesn't get affected:", obj)

  obj.changeJerseyNumberByRef()
  fmt.Println("After Change by ref, original object gets affected:", obj)
}

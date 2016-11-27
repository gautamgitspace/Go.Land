package main
 import (
   "fmt"
 )

//CUSTOM interface WITH A SINGLE METHOD SIGNATURE
type Animal interface {
 Speak() string
}

//CUSTOM TYPE struct which implement Speak()
type Dog struct {
}
func (d Dog) Speak() string{
  return "woof!"
}

type Cat struct {
}
func (c Cat) Speak() string {
  return "Meow!"
}

type Cow struct {
}
func (c Cow) Speak() string {
  return "Moo!"
}


func main(){
  //explicit cast to interface - Animal
  poodle := Animal(Dog{})
  //obj dump
  fmt.Println(poodle)
  //slice of type Animal
  animals := []Animal{Dog{}, Cat{}, Cow{}}
  for _, animal := range animals {
    fmt.Println(animal.Speak())
  }
}

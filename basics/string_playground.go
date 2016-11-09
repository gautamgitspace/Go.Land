package main

import(
  "fmt"
  "strings"
)

func main(){
  str1 := "implicitly typed"
  var str2 string = "EXPLICITLY TYPED"
  str3 := "IMPLICITLY TYPED"
  fmt.Println(str1)
  fmt.Println(strings.ToLower(str2))
  fmt.Println(strings.Title(str1))
  fmt.Println("EQUALS CASE-SENSITIVE?", (str1 == str2))
  fmt.Println("EQUALS? NON-CASE-SENSITIVE?", strings.EqualFold(str1, str3))
  fmt.Println("CONTAINS EXP?", strings.Contains(str2, "EXP"))
}

package main
import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)

func  main(){
  // var s string
  // fmt.Scanln(&s)
  // fmt.Println(s)
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter Text: ")
  //ignore the error object by using a underscore
  str, _ := reader.ReadString('\n')
  fmt.Println(str)

  fmt.Print("Enter A Number: ")
  str1, _ := reader.ReadString('\n')
  f, err := strconv.ParseFloat(strings.TrimSpace(str1), 64)
  //now that you've declared the error object you need to deal with it
  if err!=nil{
    fmt.Println(err)
  }else{
    fmt.Println("Value of number", f)
  }
}

package main
import "fmt"

func main(){
  //these are implicit typed declarations(inferred from initial value)
  str1:= "This behavior"
  str2:= "is unique"
  str3:= "to Go"
  num:= 73
  isTrue:= true

  stringLength, err := fmt.Println(str1, str2, str3)

  if err == nil{
    fmt.Println("LENGTH OF THE STRING:",stringLength)
  }
fmt.Printf("VALUE OF NUMBER: %v\n", num)
fmt.Printf("VALUE OF BOOL: %v\n", isTrue)
fmt.Printf("VALUE OF NUMBER AS A FLOAT: %.2f\n", float64(num))

fmt.Printf("DATA TYPES: %T, %T, %T, %T and %T\n", str1, str2, str3, num, isTrue)

myString := fmt.Sprintf("DATA TYPES as var: %T, %T, %T, %T and %T", str1, str2, str3, num, isTrue)
fmt.Printf(myString)
}

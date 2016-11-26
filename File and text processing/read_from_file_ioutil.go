package main
import (
  "io/ioutil"
  "fmt"
)
func main() {
  file := "/Users/Gautam/Desktop/phones.txt"
  content, err := ioutil.ReadFile(file)
  checkError(err)
  fmt.Println("Dumping raw bytes:\n\n", content, "\n")

  result := string(content)
  fmt.Println("Reading content from dump:\n\n", result)
}
func checkError(err error) {
  if err!=nil {
    panic(err)
  }
  }

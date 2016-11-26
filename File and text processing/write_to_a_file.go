package main
import (
  "fmt"
  "os"
  "io"
  "io/ioutil"
)
func main() {
  content := "sample text"
  file, err := os.Create("/Users/Gautam/Work/Go.Land/File and text processing/file_string_created_by_golang.txt")
  checkError(err)
  defer file.Close()

  ln, err := io.WriteString(file, content)
  checkError(err)

  fmt.Printf("Done writing %v characters to file", ln)

  bytes := []byte(content)
  ioutil.WriteFile("/Users/Gautam/Work/Go.Land/File and text processing/file_bytes_create_by_golang.txt", bytes, 0644)

  err_rem := os.Remove("/Users/Gautam/Work/Go.Land/File and text processing/file_bytes_create_by_golang.txt")
  checkError(err_rem)

}

func checkError(err error) {
  if err!=nil{
    panic(err)
  }
}

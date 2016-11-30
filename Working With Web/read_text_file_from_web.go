package main
import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {
  url := "https://gist.github.com/gautamgitspace/27f259be1fae711106a08f172186faad"

  resp, err := http.Get(url)
  if err!=nil {
    panic(err)
  }
  fmt.Printf("Response type: %T\n", resp)
  defer resp.Body.Close()
  //READ CONTENTS
  bytes, err := ioutil.ReadAll(resp.Body)
  if err!=nil {
    panic(err)
  }
  content  := string(bytes)
  fmt.Print(content)
}

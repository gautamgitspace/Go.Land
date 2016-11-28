package main
import (
  "fmt"
  "net/http"
)

type Sample struct{}

//THIS FUNC WILL BE CALLED EACH TIME A REQ IS MADE
func (s Sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "<h1>Welcome to Go.Land Server!</h1>")
}

func main() {
  var s Sample
  err := http.ListenAndServe("localhost:4000", s)
  checkError(err)
}
func checkError(err error) {
  if err!=nil {
    panic(err)
  }
}

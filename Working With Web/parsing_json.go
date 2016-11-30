package main
import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strings"
  "math/big"
)

type Tour struct {
  Name, Price string
}

func main() {
  url := "http://services.explorecalifornia.org/json/tours.php"
  content := getContentFromServer(url)
  tours := walkWithMeJson(content)

  //FURTHER PARSE NAME AND PRICE
  for _, tour := range tours {
    price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
    fmt.Printf("%v($%.2f)\n", tour.Name, price)
  }
}

//READ CONTENTS
func getContentFromServer(url string) string {
  resp, err := http.Get(url)
  checkError(err)
  defer resp.Body.Close()
  bytes, err := ioutil.ReadAll(resp.Body)
  checkError(err)

  return string(bytes)
}

//ERROR CHECK METHOD
func checkError(err error) {
  if err!=nil {
    panic(err)
  }
}

//PARSE JSON
func walkWithMeJson(content string) []Tour {
  tours := make([]Tour, 0, 20)
  decoder := json.NewDecoder(strings.NewReader(content))

  _, err := decoder.Token()
  checkError(err)

  var tour Tour
  //loop till decoder has data
  for decoder.More() {
    err := decoder.Decode(&tour)
    checkError(err)
    tours = append(tours, tour)
  }
  return tours
}

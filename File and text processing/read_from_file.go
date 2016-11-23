package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    file, err := os.Open("/Users/Gautam/Desktop/phones.txt")
    if err != nil {
        fmt.Print(err)
    }
    defer file.Close()
    container := ""
    var tokens []string


    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
      container = scanner.Text()
      tokens = strings.Split(container, ",")
    }
    fmt.Println("Last Container:", container)
    fmt.Println("Last token[2]:", strings.TrimSpace(tokens[2]))
}

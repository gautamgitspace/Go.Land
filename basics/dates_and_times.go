package main
import(
  "fmt"
  "time"
)
func main(){
  t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
  fmt.Printf("Go Launched at %s\n", t)

  now := time.Now()
  fmt.Printf("Time Now Is: %s\n", now)
  fmt.Println("The Month is:", t.Month())

  tomorrow := now.AddDate(0,0,1)
  fmt.Printf("tommorrow is %v, %v, %v %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())
  longFormat := "Monday, 1/2/06"
  fmt.Println("tomorrow is", tomorrow.Format(longFormat))
}

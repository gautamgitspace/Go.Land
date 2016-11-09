package main
import "fmt"
import "math/big"
import "math"

func main(){
i1, i2, i3 := 12, 45, 68
sum_i := i1 + i2 + i3
fmt.Println("INTEGER SUM:", sum_i)

f1, f2, f3 := 23.2, 6.7, 0.145
sum_f := f1 + f2 + f3
fmt.Println("FLOAT SUM:", sum_f)

var b1, b2, b3, bigSum big.Float
b1.SetFloat64(23.3)
b1.SetFloat64(23.6)
b1.SetFloat64(23.1)

bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
fmt.Printf("BIG SUM = %.10g\n", &bigSum)

circleRadius := 15.8
circumference := circleRadius * math.Pi
fmt.Printf("Circumference: %.2f\n", circumference)
}

package main
import "math"
import "fmt"

var Pi float64

func init() {
	Pi = 4 * math.Atan(1)
}

func main() {
	fmt.Printf("hello %f\n", Pi)
}



package main
import "fmt"

func main(){
	slice1 := make([]int, 5, 10)
	fmt.Printf("slice len = %d\n", len(slice1))
	fmt.Printf("slice cap = %d\n", cap(slice1))

	slice2 := new([100]int)[0:50]
	fmt.Printf("slice2 len = %d\n", len(slice2))
	fmt.Printf("slice2 cap = %d\n", cap(slice2))
}


package main
import "fmt"

func main(){
	var num int = 9
	switch {
	case num < 0:
		fmt.Println("Number is negetive")
	case num >0 && num<10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}
}


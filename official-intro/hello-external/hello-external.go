package main

import	"fmt"
import	"rsc.io/quote"

func main(){
	fmt.Println(quote.Hello())
}
/**
Put your own code in a module for tracking dependencies.
When your code imports packages from another module, a go.mod file lists the specific modules and versions providing those packages. That file stays with your code, including in your source code repository.

To create a go.mod file, run the go mod init command, giving it the name of the module your code will be in (here, just use "hello"):
**/

/**
use command to change proxy
go env -w GOPROXY=https://goproxy.cn
**/

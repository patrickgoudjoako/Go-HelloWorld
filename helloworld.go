package main

import (
	"fmt"
	"runtime"
	"reflect"
)

var (
	name, course 	string
	module			float64
)

func main() {
	fmt.Println("Hello from", runtime.GOOS)	
	fmt.Println("Name is ", name, " and is of type", reflect.TypeOf(name))	
	fmt.Println("Module is ", module, " and is of type", reflect.TypeOf(module))	

}
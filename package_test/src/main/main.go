package main

import "fmt"
import m "test_package"

func main() {
	fmt.Println("main func")
	lib_func1()
	m.Test_package_func1()
}
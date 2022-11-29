package main

import "fmt"

func Hello1(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello1("world"))
}

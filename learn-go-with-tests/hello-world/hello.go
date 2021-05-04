package main

import "fmt"

const prefix = "Hello, "
const suffix = "!"

func Hello(name string) string {
	return prefix + name + suffix
}

func main() {
	fmt.Println(Hello("world"))
}

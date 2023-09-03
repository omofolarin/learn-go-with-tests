package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case "Spanish":
		return "Hola, " + name

	case "French":
		return "Bonjour, " + name
	default:
		return "Hello, " + name
	}
}

func main() {
	fmt.Println(Hello("folarin", ""))
}

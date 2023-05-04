package say_hello_modul

import "fmt"

func SayHello() string {
	return "Hello World!"
}

func SayHelloName(name string) string {
	if name != "" {
		return fmt.Sprintf("Hello, %s", name)
	}
	return "Hello, World"
}

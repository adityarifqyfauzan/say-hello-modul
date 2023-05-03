package say_hello_modul

func SayHello() string {
	return "Hello World!"
}

func sayHelloName(name string) string {
	if name != "" {
		return "Hello " + name
	}
	return "Hello, World"
}

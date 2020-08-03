package hello

// HelloWorld says Hello
func HelloWorld() (string, int) {
	hello := "Hello World"
	return hello, len(hello)
}

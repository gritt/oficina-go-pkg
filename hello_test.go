package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	// given
	want := "Hello World"

	// when
	got := HelloWorld()

	// then
	if got != want {
		t.Errorf("HelloWorld() = %v, want %v", got, want)
	}
}

package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	// given
	wantHello := "Hello World"
	wantLen := 11

	// when
	gotHello, gotLen := HelloWorld()

	// then
	if gotHello != wantHello {
		t.Errorf("HelloWorld() = %v, want %v", gotHello, wantHello)
	}
	if gotLen != wantLen {
		t.Errorf("HelloWorld() = %d, want %d", gotLen, wantLen)
	}
}

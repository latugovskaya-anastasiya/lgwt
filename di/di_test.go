package main_test

import (
	"bytes"
	"testing"

	"github.com/latugovskaya-anastasiya/lgwt/di"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	main.Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

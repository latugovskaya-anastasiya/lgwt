package integers_test

import (
	"fmt"
	"testing"

	"github.com/latugovskaya-anastasiya/lgwt/integers"
)

func TestAdder(t *testing.T) {
	sum := integers.Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
func ExampleAdd() {
	sum := integers.Add(1, 6)
	fmt.Println(sum)
	// Output: 7
}

package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestDivided(t *testing.T) {
	divided := Divide(10, 5)
	expected := 2

	if divided != expected {
		t.Errorf("got %d, want %d", divided, expected)
	}

}

// func Benchmark(b *testing.B){
//  for i := 0; i < b.N; i++{
// 	 Repeat("a")
//  }
// }

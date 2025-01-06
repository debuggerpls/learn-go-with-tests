package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' 3 times", func(t *testing.T) {
		got := Repeat("a", 3)
		want := "aaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("repeat 'b' 5 times", func(t *testing.T) {
		got := Repeat("b", 5)
		want := "bbbbb"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	// runs the benchmark b.N times
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	output := Repeat("ab", 3)
	fmt.Println(output)
	// Output: ababab
}

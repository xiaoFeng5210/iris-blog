package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expceted := 4
	if sum != expceted {
		t.Errorf("expected '%d' but got '%d'", expceted, sum)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

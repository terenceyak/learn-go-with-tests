package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkReport(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

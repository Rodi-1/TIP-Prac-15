package stringsx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClip(t *testing.T) {
	for _, c := range []struct {
		name, s string
		max     int
		want    string
	}{{"empty", "", 5, ""}, {"zero", "hello", 0, ""}, {"negative", "hello", -1, ""}, {"equal", "hello", 5, "hello"}, {"more", "hello", 6, "hello"}, {"truncate", "hello", 3, "hel"}} {
		t.Run(c.name, func(t *testing.T) { assert.Equal(t, c.want, Clip(c.s, c.max)) })
	}
}
func FuzzClip(f *testing.F) {
	f.Add("hello", 3)
	f.Add("", -1)
	f.Add("Привет", 5)
	f.Fuzz(func(t *testing.T, s string, max int) {
		got := Clip(s, max)
		want := max
		if want < 0 {
			want = 0
		}
		if want > len(s) {
			want = len(s)
		}
		if len(got) != want || got != s[:want] {
			t.Fatalf("Clip length/prefix invariant")
		}
	})
}

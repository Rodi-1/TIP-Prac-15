package mathx

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSum_Table(t *testing.T) {
	for _, c := range []struct {
		name       string
		a, b, want int
	}{{"positive", 2, 3, 5}, {"negative", -10, -5, -15}, {"mixed", 10, -5, 5}, {"zero", 0, 0, 0}} {
		t.Run(c.name, func(t *testing.T) {
			if got := Sum(c.a, c.b); got != c.want {
				t.Fatalf("got %d want %d", got, c.want)
			}
		})
	}
}
func TestDivide_OkAndError(t *testing.T) {
	for _, c := range []struct {
		name       string
		a, b, want int
	}{{"positive", 10, 2, 5}, {"negative", -10, 2, -5}, {"truncate", 7, 2, 3}, {"zero numerator", 0, 2, 0}} {
		t.Run(c.name, func(t *testing.T) { got, e := Divide(c.a, c.b); require.NoError(t, e); assert.Equal(t, c.want, got) })
	}
	_, e := Divide(10, 0)
	require.Error(t, e)
	assert.ErrorIs(t, e, ErrDivideByZero)
}
func TestMustDivide(t *testing.T) {
	require.Equal(t, 5, MustDivide(10, 2))
	require.Panics(t, func() { MustDivide(10, 0) })
}

var benchmarkResult int

func BenchmarkSum(b *testing.B) {
	a, c := 123, 456
	for i := 0; i < b.N; i++ {
		benchmarkResult = Sum(a, c)
	}
}

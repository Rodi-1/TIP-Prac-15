package mathx

import "errors"

var ErrDivideByZero = errors.New("divide by zero")

func Sum(a, b int) int { return a + b }
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func MustDivide(a, b int) int {
	result, e := Divide(a, b)
	if e != nil {
		panic(e)
	}
	return result
}

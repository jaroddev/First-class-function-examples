package functional

import "math"

type Calculator struct {
	accumulator float64
}

// No more hardcoded operation
// Rather we are encapsulating behavior into a function
// And grow the possibilities at runtime
// Problem, we still cannot correctly implement some operations
// See WrongSqrt
type wrongOperationFunc func(float64, float64) float64

func (calculator *Calculator) DoWrong(wrongOperation wrongOperationFunc, value float64) float64 {
	calculator.accumulator = wrongOperation(calculator.accumulator, value)
	return calculator.accumulator
}

func WrongAdd(a, b float64) float64 {
	return a + b
}

func WrongSub(a, b float64) float64 {
	return a - b
}

func WrongMul(a, b float64) float64 {
	return a * b
}

func WrongDiv(a, b float64) float64 {
	if b == 0 {
		panic("Bad value ! 0 can't divide anything !!")
	}
	return a / b
}

// Well we can actually try to manage
// But it will one day reach its limit
// And it will cost us plenty of effort to fix it
func WrongSqrt(n, _ float64) float64 {
	return math.Sqrt(n)
}

func wrongMain() {
	cal := Calculator{
		accumulator: 1,
	}

	cal.DoWrong(WrongMul, 9)
	cal.DoWrong(WrongAdd, 3)

	// Here whatever we pass, will be ignored
	// Useless constant that can confuse developpers
	cal.DoWrong(WrongSqrt, 0)
}

// To Fix the problem, we use a functional pattern
// The First Class Function pattern
type operationFunc func(float64) float64

func (cal *Calculator) Do(operation operationFunc) float64 {
	cal.accumulator = operation(cal.accumulator)
	return cal.accumulator
}

func Add(n float64) operationFunc {
	return func(accumulator float64) float64 {
		return accumulator + n
	}
}

func Sub(n float64) operationFunc {
	return func(accumulator float64) float64 {
		return accumulator - n
	}
}

func Mul(n float64) operationFunc {
	return func(accumulator float64) float64 {
		return accumulator * n
	}
}

func Div(n float64) operationFunc {
	return func(accumulator float64) float64 {
		if n == 0 {
			panic("Bad value ! 0 can't divide anything !!")
		}
		return accumulator / n
	}
}

// Magic, it no longer takes no argument
func Sqrt() operationFunc {
	return func(f float64) float64 {
		return math.Sqrt(f)
	}
}

// Magic, it no longer takes no argument
func Pow(power int) operationFunc {
	return func(f float64) float64 {
		return math.Pow(f, float64(power))
	}
}

func main() {
	cal := Calculator{
		accumulator: 1,
	}

	cal.Do(Mul(9))
	cal.Do(Add(3))

	// Here too, it no longer takes no argument
	cal.Do(Sqrt())
}

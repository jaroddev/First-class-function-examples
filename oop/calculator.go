package oop

type Calculator struct {
	accumulator float64
}

const (
	OP_ADD = 1 << iota
	OP_SUB
	OP_MUL
	OP_DIV
)

// Normally, as the state was updated
// the method should not return anything but it does
// Here this function does a lot of thing and should be better handled
// Hard to read once we make it grow
func (calculator *Calculator) Do(operation int, v float64) float64 {
	switch operation {
	case OP_ADD:
		calculator.accumulator += v
	case OP_SUB:
		calculator.accumulator -= v
	case OP_MUL:
		calculator.accumulator *= v
	case OP_DIV:
		if v == 0 {
			panic("Bad value ! 0 can't divide anything !!")
		}
		calculator.accumulator /= v
	default:
		panic("unhandled operation")
	}
	return calculator.accumulator
}

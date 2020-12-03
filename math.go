package tinylisp

// Addition implements addition for TinyLisp.
type Addition struct{}

// Arity returns infiniteArity for +.
func (a *Addition) Arity() int {
	return infiniteArity
}

// Call implements the addition.
func (a *Addition) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	if len(arguments) < 2 {
		return nil, &executionError{line, "'+' requires at least two arguments"}
	}

	var sum float64 = 0

	for _, arg := range arguments {
		if num, ok := arg.(float64); ok {
			sum += num
		} else {
			return nil, &executionError{line, "'+' is only defined for numbers"}
		}
	}

	return sum, nil
}

// Subtraction implements subtraction for TinyLisp.
type Subtraction struct{}

// Arity returns infiniteArity for +.
func (s *Subtraction) Arity() int {
	return infiniteArity
}

// Call implements the subtraction.
func (s *Subtraction) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	if len(arguments) < 2 {
		return nil, &executionError{line, "'-' requires at least two arguments"}
	}

	var result float64

	if num, ok := arguments[0].(float64); ok {
		result = num
	} else {
		return nil, &executionError{line, "'-' is only defined for numbers"}
	}

	for i, arg := range arguments {
		if i == 0 {
			continue
		}

		if num, ok := arg.(float64); ok {
			result -= num
		} else {
			return nil, &executionError{line, "'-' is only defined for numbers"}
		}
	}

	return result, nil
}

// Multiplication implements multiplication for TinyLisp.
type Multiplication struct{}

// Arity returns infiniteArity for *.
func (m *Multiplication) Arity() int {
	return infiniteArity
}

// Call implements the multiplication.
func (m *Multiplication) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	if len(arguments) < 2 {
		return nil, &executionError{line, "'*' requires at least two arguments"}
	}

	var result float64

	if num, ok := arguments[0].(float64); ok {
		result = num
	} else {
		return nil, &executionError{line, "'*' is only defined for numbers"}
	}

	for i, arg := range arguments {
		if i == 0 {
			continue
		}

		if num, ok := arg.(float64); ok {
			result *= num
		} else {
			return nil, &executionError{line, "'*' is only defined for numbers"}
		}
	}

	return result, nil
}

// Division implements division for TinyLisp.
type Division struct{}

// Arity returns infiniteArity for /.
func (d *Division) Arity() int {
	return infiniteArity
}

// Call implements the division.
func (d *Division) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	if len(arguments) < 2 {
		return nil, &executionError{line, "'/' requires at least two arguments"}
	}

	var result float64

	if num, ok := arguments[0].(float64); ok {
		result = num
	} else {
		return nil, &executionError{line, "'/' is only defined for numbers"}
	}

	for i, arg := range arguments {
		if i == 0 {
			continue
		}

		if num, ok := arg.(float64); ok {
			if num != 0 {
				result /= num
			} else {
				return nil, &executionError{line, "Division by zero"}
			}
		} else {
			return nil, &executionError{line, "'/' is only defined for numbers"}
		}
	}

	return result, nil
}

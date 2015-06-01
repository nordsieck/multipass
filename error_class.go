package multipass

type ErrorClass struct {
	id          uint64
	errors      map[error]struct{}
	multiPasses map[uint64]struct{}
}

func NewErrorClass(es ...error) ErrorClass {
	ec := ErrorClass{
		id:          id(),
		errors:      map[error]struct{}{},
		multiPasses: map[uint64]struct{}{},
	}

	for _, e := range es {
		switch e := e.(type) {
		case Error:
			ec.multiPasses[e.id] = struct{}{}
		case *Error:
			ec.multiPasses[e.id] = struct{}{}
		case nil: // do nothing
		default:
			ec.errors[e] = struct{}{}
		}
	}
	return ec
}

func (e *ErrorClass) Contains(err error) bool {
	ok := false
	switch err := err.(type) {
	case nil: // do nothing
	case Error:
		_, ok = e.multiPasses[err.id]
	case *Error:
		_, ok = e.multiPasses[err.id]
	default:
		_, ok = e.errors[err]
	}
	return ok
}

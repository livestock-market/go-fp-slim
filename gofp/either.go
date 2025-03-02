package gofp

type Either[R any] struct {
	err    Option[XError]
	value  R
	isLeft bool
}

func Left[R any](err XError) Either[R] {
	return Either[R]{
		err:    Some(err),
		isLeft: true,
	}
}

func Right[R any](value R) Either[R] {
	return Either[R]{
		value:  value,
		isLeft: false,
		err:    None[XError](),
	}
}

func (e Either[R]) IsLeft() bool {
	return e.isLeft
}

func (e Either[R]) IsRight() bool {
	return !e.isLeft
}

func (e Either[R]) Left() XError {
	if e.isLeft && e.err.HasValue() {
		return e.err.Get()
	}

	panic("Either is not left, Left is empty")
}

func (e Either[R]) Right() R {
	if e.isLeft {
		panic("Cannot get Right value from Left")
	}
	return e.value
}

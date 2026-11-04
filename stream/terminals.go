package stream

func (s Stream[T]) ToList() []T {
	result := make([]T, 0)
	if s.next == nil {
		return result
	}
	for {
		item, ok := s.next()
		if !ok {
			return result
		}
		result = append(result, item)
	}
}

func (s Stream[T]) Reduce[U any](identity U, f func(U, T) U) U {
	value := identity
	next := s.next
	if next == nil {
		return value
	}

	for {
		item, ok := next()
		if !ok {
			return value
		}
		value = f(value, item)
	}
}

func (s Stream[T])

//func (s Stream[T]) Sum() T {
//	var zero T
//	return s.Reduce(zero, Add)
//}

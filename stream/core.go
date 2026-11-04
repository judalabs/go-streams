package stream

type Number interface {
	int | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

type Stream[T any] struct {
	next func() (T, bool)
}

func From[T any](list []T) Stream[T] {
	i := 0
	return Stream[T]{
		next: func() (T, bool) {
			if i >= len(list) {
				var zero T
				return zero, false
			}
			item := list[i]
			i++
			return item, true
		},
	}
}

func (s Stream[T]) Map[U any](function func(T) U) Stream[U] {
	return Stream[U]{
		next: func() (U, bool) {
			if s.next == nil {
				var zero U
				return zero, false
			}
			item, ok := s.next()
			if !ok {
				var zero U
				return zero, false
			}
			return function(item), true
		},
	}
}

func (s Stream[T]) Filter(predicate func(T) bool) Stream[T] {
	return Stream[T]{
		next: func() (T, bool) {
			if s.next == nil {
				var zero T
				return zero, false
			}
			for {
				item, ok := s.next()
				if !ok {
					var zero T
					return zero, false
				}
				if predicate(item) {
					return item, true
				}
			}
		},
	}
}

func (s Stream[T]) Peek(consumer func(T)) Stream[T] {
	return Stream[T]{
		next: func() (T, bool) {
			if s.next == nil {
				var zero T
				return zero, false
			}
			item, ok := s.next()
			if !ok {
				var zero T
				return zero, false
			}
			consumer(item)
			return item, true
		},
	}
}
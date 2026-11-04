package stream

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Add[T Number](a, b T) T {
	return a + b
}

func Subtract[T Number](a, b T) T {
	return a - b
}
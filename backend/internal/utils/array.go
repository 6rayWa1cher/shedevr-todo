package utils

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func MapWithError[T, V any](ts []T, fn func(T) (V, error)) ([]V, error) {
	result := make([]V, len(ts))
	var err error
	for i, t := range ts {
		result[i], err = fn(t)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

package slice

import "errors"

type Num interface {
	int | int32 | int64 | uint32 | uint64 | float32 | float64
}

func MinValue[T Num](s []T) (T, error) {
	if len(s) == 0 {
		return 0, errors.New("empty slice")
	}

	min := s[0]
	for _, v := range s {
		if min > v {
			min = v
		}

	}
	return min, nil
}

func MaxValue[T Num](s []T) (T, error) {
	if len(s) == 0 {
		return 0, errors.New("empty slice")
	}

	max := s[0]
	for _, v := range s {
		if max < v {
			max = v
		}

	}
	return max, nil
}

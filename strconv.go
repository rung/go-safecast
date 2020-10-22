package safecast

import "strconv"

func Atoi32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

func Atoi16(s string) (int16, error) {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(i), nil
}

func Atoi8(s string) (int8, error) {
	i, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(i), nil
}
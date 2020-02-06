package tool

import "strconv"

func String2Uint(str string) (uint, error) {
	tmp, err := strconv.ParseUint(str, 10, 32<<(^uint(0)>>63))
	return uint(tmp), err
}

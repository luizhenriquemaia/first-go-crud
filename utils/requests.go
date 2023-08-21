package utils

import "strconv"

func ParseParamId(id string) int64 {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}
	return idInt
}

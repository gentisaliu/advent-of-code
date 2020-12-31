package util

import "strconv"

func ConvertIntToBinary(intValue, length int) string {
	binary := strconv.FormatInt(int64(intValue), 2)
	return LeftPadToLength(binary, "0", length)
}

func ConvertBinaryToInt(binaryString string) int64 {
	n, _ := strconv.ParseInt(binaryString, 2, 64)
	return n
}

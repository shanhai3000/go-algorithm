package algorithm

import (
	"strconv"
	"strings"
)

/**
each number's length < 110
*/
func StringMultiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num1Arr := make([]uint8, len(num1))
	num2Arr := make([]uint8, len(num2))
	ans := make([]uint8, len(num1)*len(num2)+1)
	num1Len := len(num1) - 1
	num2Len := len(num2) - 1
	for i := 0; i <= num1Len; i++ {
		num1Arr[i] = char2Uint8(num1[num1Len-i])
	}
	for i := 0; i <= num2Len; i++ {
		num2Arr[i] = char2Uint8(num2[num2Len-i])
	}
	var tmp uint8
	ansLen := 0
	for i := 0; i <= num1Len; i++ {
		for j := 0; j <= num2Len; j++ {
			tmp = num2Arr[j] * num1Arr[i]
			append0(ans, tmp, i+j, &ansLen)
		}
	}
	var builder strings.Builder
	for i := ansLen; i >=0; i-- {
		builder.WriteString(strconv.Itoa(int(ans[i])))
	}
	return builder.String()
}
func append0(arr []uint8, num uint8, shift int, ansLen *int) {
	if *ansLen < shift{
		*ansLen = shift
	}
	sum := arr[shift]+num
	if sum< 10 {
		arr[shift] = sum
	} else {
		arr[shift] = sum % 10
		append0(arr, sum/10, shift+1, ansLen)
	}
}
func char2Uint8(c uint8) uint8 {
	return c - '0'
}

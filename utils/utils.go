package utils

const MyMinInt int = -int(^uint(0)>>1) - 1

func Btoi(a byte) int {
	return int(a - "0"[0])
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxSlice(nums []int) int {
	max := int(^uint(0) >> 1)
	for _, i := range nums {
		if i > max {
			max = i
		}
	}
	return max
}

func Abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

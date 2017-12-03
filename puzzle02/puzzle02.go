package puzzle01

import (
	"log"
	"strconv"

	"github.com/parafiend/advent2017/base"
)

const id = "2"

type Puzzle struct{}

func (p Puzzle) Test() string {
	test := []string{
		"5 1 9 5",
		"7 5 3",
		"2 4 6 8",
	}
	return strconv.Itoa(checksum(test))
}

func (p Puzzle) Phase1() string {
	return "doin phase1"
}

func (p Puzzle) Phase2() string {
	return "doin phase2"
}

func (p Puzzle) String() string {
	return "puzz" + id
}

func checksum(input []string) int {
	var result int
	for _, v := range input {
		minVal := int(^uint(0) >> 1)
		maxVal := 0
		for j := 0; j < len(v); j++ {
			if v[j] == " "[0] {
				continue
			}
			curr := btoi(v[j])
			minVal = min(minVal, curr)
			maxVal = max(maxVal, curr)
		}
		diff := maxVal - minVal
		log.Println(minVal, maxVal, diff)
		result += diff
	}
	log.Println(result)
	return result
}

func btoi(a byte) int {
	return int(a - "0"[0])
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func init() {
	base.Register(id, Puzzle{})
}

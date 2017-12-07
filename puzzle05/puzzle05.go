package puzzle05

import (
	"log"
	"strconv"
	"strings"

	"github.com/parafiend/advent2017/base"
)

const id = "5"

type Puzzle struct{}

func (p Puzzle) Test() string {
	testInput := `0
	3
	0
	1
	-3`
	log.Println(strconv.Itoa(countJumps(testInput)))
	return strconv.Itoa(countJumps2(testInput))
}

func (p Puzzle) Phase1() string {
	return strconv.Itoa(countJumps(rawJumps))
}

func (p Puzzle) Phase2() string {
	return strconv.Itoa(countJumps2(rawJumps))
}

func (p Puzzle) String() string {
	return "puzz" + id
}

func countJumps(input string) int {
	jumps := parseInput(input)
	count := 0
	for idx := 0; idx < len(jumps); count++ {
		jump := jumps[idx]
		jumps[idx]++
		idx += jump
		log.Println(idx, jump)
	}
	return count
}

func countJumps2(input string) int {
	jumps := parseInput(input)
	count := 0
	for idx := 0; idx < len(jumps) && idx >= 0; count++ {
		jump := jumps[idx]
		switch {
		case jump >= 3:
			jumps[idx]--
		default:
			jumps[idx]++
		}
		idx += jump
		if count%10000 == 0 {
			log.Println(idx, jump)
		}
	}
	return count
}

func parseInput(input string) []int {
	pieces := strings.Fields(input)
	var jumps []int
	for _, v := range pieces {
		if jump, err := strconv.Atoi(v); err == nil {
			jumps = append(jumps, jump)
		}
	}
	log.Println(input, pieces, jumps)
	return jumps
}

func init() {
	base.Register(id, Puzzle{})
}

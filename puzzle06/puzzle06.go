package puzzle06

import (
	"log"
	"strconv"
	"strings"

	"github.com/parafiend/advent2017/base"
)

const id = "6"

type Puzzle struct{}

func (p Puzzle) Test() string {
	testInput := `0 2 7 0`
	count, delta := countRedistributes(testInput)
	log.Println(count, delta)
	return "lalalal"
}

func (p Puzzle) Phase1() string {
	count, _ := countRedistributes(rawBanks)
	return strconv.Itoa(count)
}

func (p Puzzle) Phase2() string {
	_, delta := countRedistributes(rawBanks)
	return strconv.Itoa(delta)
}

func (p Puzzle) String() string {
	return "puzz" + id
}

func countRedistributes(input string) (int, int) {
	banks := parseInput(input)
	count := 0
	delta := 0
	seen := make(map[string]int)
	for {
		banks = redistribute(banks)
		log.Println(banks)
		hash := getHash(banks)
		count++
		if last, ok := seen[hash]; ok {
			delta = count - last
			break
		}
		seen[hash] = count
	}
	return count, delta
}

func getHash(banks []int) string {
	var hash string
	for _, v := range banks {
		hash += strconv.Itoa(v) + ","
	}
	return hash
}

func parseInput(input string) []int {
	var banks []int
	for _, v := range strings.Fields(input) {
		if bank, err := strconv.Atoi(v); err == nil {
			banks = append(banks, bank)
		}
	}
	return banks
}

func redistribute(banks []int) []int {
	src, count := findBiggest(banks)
	banks[src] = 0
	for i := 1; i <= count; i++ {
		target := (src + i) % len(banks)
		banks[target]++
	}
	return banks
}

func findBiggest(banks []int) (int, int) {
	idx, count := 0, -1
	for i, v := range banks {
		if v > count {
			idx = i
			count = v
		}
	}
	return idx, count
}

func init() {
	base.Register(id, Puzzle{})
}

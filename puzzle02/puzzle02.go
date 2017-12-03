package puzzle02

import (
	"log"
	"strconv"
	"strings"

	"github.com/parafiend/advent2017/base"
)

const id = "2"

type Puzzle struct{}

func (p Puzzle) Test() string {
	test := `5 1 9 5
		7 5 3
		2 4 6 8`
	log.Println(strconv.Itoa(checksum(test)))

	test2 := `5 9 2 8
	9 4 7 3
	3 8 6 5`
	log.Println(checksum2(test2))
	return "what"
}

func (p Puzzle) Phase1() string {
	return strconv.Itoa(checksum(Spreadsheet))
}

func (p Puzzle) Phase2() string {
	return strconv.Itoa(checksum2(Spreadsheet))
}

func (p Puzzle) String() string {
	return "puzz" + id
}

func toCells(input string) [][]string {
	rows := strings.Split(input, "\n")
	var cells [][]string
	for _, v := range rows {
		row := strings.Fields(v)
		cells = append(cells, row)
		//log.Println(cells)
	}
	return cells
}

func checksum(input string) int {
	sheet := toCells(input)
	var result int
	for _, cells := range sheet {
		minVal := int(^uint(0) >> 1)
		maxVal := 0
		for _, cell := range cells {
			curr, _ := strconv.Atoi(cell)
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

func checksum2(input string) int {
	sheet := toCells(input)
	var result int
	for _, cells := range sheet {
		var big, little int
		var found bool
		for i, cell := range cells {
			curr, _ := strconv.Atoi(cell)
			for j := i + 1; j < len(cells); j++ {
				candidate, _ := strconv.Atoi(cells[j])

				if curr%candidate == 0 || candidate%curr == 0 {
					big, little = max(curr, candidate), min(curr, candidate)
					found = true
					break
				}
			}
			if found {
				divide := big / little
				log.Println(little, big, divide)
				result += divide
				break
			}
		}
	}
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

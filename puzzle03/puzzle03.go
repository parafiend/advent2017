package puzzle01

import (
	"log"
	"strconv"

	"github.com/parafiend/advent2017/base"
	"github.com/parafiend/advent2017/utils"
)

const id = "3"

type Puzzle struct{}

func (p Puzzle) Test() string {
	log.Println(cartesianize(1))
	log.Println(cartesianize(4))
	log.Println(cartesianize(9))
	log.Println(cartesianize(11))
	log.Println(cartesianize(25))
	log.Println(cartesianize(26))

	log.Println(distance(cartesianize(1)))
	log.Println(distance(cartesianize(12)))
	log.Println(distance(cartesianize(23)))
	log.Println(distance(cartesianize(1024)))

	log.Println(findSumAbove(1))
	log.Println(findSumAbove(4))
	log.Println(findSumAbove(800))

	return "doin a test"
}

func (p Puzzle) Phase1() string {
	return strconv.Itoa(distance(cartesianize(312051)))
}

func (p Puzzle) Phase2() string {
	return strconv.Itoa(findSumAbove(312051))
}

func distance(x int, y int) int {
	return utils.Abs(x) + utils.Abs(y)
}

func cartesianize(cell int) (int, int) {
	if cell == 1 {
		return 0, 0
	}

	ring := 1
	for i := 1; i < cell; i += 2 {
		if cell <= i*i {
			break
		}
		ring++
	}
	square := (ring * 2) - 1
	baseX, baseY := (ring - 1), (ring - 2)
	start := ringMax(ring-1) + 1
	toWalk := cell - start

	//mode := toWalk / square
	//remainder := toWalk % square

	for i := 0; i < toWalk; i++ {
		switch {
		case i < square-2:
			baseY--
		case i < 2*(square-1)-1:
			baseX--
		case i < 3*(square-1)-1:
			baseY++
		default:
			baseX++
		}
	}
	//log.Println(baseX, baseY, cell)
	return baseX, baseY
}

func findSumAbove(input int) int {
	var maxVal int
	const cellCount = 50
	cells := make([][]int, cellCount)
	for i := 0; i < len(cells); i++ {
		cells[i] = make([]int, cellCount)
	}
	center := cellCount / 2
	cells[center][center] = 1

	for i := 2; maxVal < input; i++ {
		x, y := cartesianize(i)
		var sum int
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if j == 0 && k == 0 {
					continue
				}
				targetX := x + j + center
				targetY := y + k + center
				sum += cells[targetX][targetY]
			}
		}
		maxVal = utils.Max(sum, maxVal)
		cells[x+center][y+center] = sum
	}
	return maxVal
}

func ringMax(ring int) int {
	ringSquare := 2*ring - 1
	return ringSquare * ringSquare
}

func (p Puzzle) String() string {
	return "puzz" + id
}

func init() {
	base.Register(id, Puzzle{})
}

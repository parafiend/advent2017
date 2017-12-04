package puzzle01

import (
	"log"

	"github.com/parafiend/advent2017/base"
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
	return "doin a test"
}

func (p Puzzle) Phase1() string {
	return "doin phase1"
}

func (p Puzzle) Phase2() string {
	return "doin phase2"
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
	log.Println(baseX, baseY, cell)
	return baseX, baseY
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

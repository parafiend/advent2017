package puzzle01

import "github.com/parafiend/advent2017/base"

const id = "3"

type Puzzle struct{}

func (p Puzzle) Test() string {
	return "doin a test"
}

func (p Puzzle) Phase1() string {
	return "doin phase1"
}

func (p Puzzle) Phase2() string {
	return "doin phase2"
}

func cartesianize(cell int) (int, int) {
	return 0, 0
}

func (p Puzzle) String() string {
	return "puzz" + id
}

func init() {
	base.Register(id, Puzzle{})
}

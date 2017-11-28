package puzzle01

import "github.com/parafiend/advent2017/base"

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

func (p Puzzle) String() string {
	return "puzz1"
}

func init() {
	base.Register("1", Puzzle{})
}

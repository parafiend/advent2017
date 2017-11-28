package base

var Puzzles = make(map[string]Puzzle)

type Puzzle interface {
	Test() string
	Phase1() string
	Phase2() string
}

func Register(id string, logic Puzzle) {
	Puzzles[id] = logic
}

func Get(id string) Puzzle {
	return Puzzles[id]
}

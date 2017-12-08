package puzzle07

import (
	"log"
	"strconv"
	"strings"

	"github.com/parafiend/advent2017/base"
)

const id = "7"

type Puzzle struct{}

func (p Puzzle) Test() string {
	testInput := `pbga (66)
	xhth (57)
	ebii (61)
	havc (66)
	ktlj (57)
	fwft (72) -> ktlj, cntj, xhth
	qoyq (66)
	padx (45) -> pbga, havc, qoyq
	tknk (41) -> ugml, padx, fwft
	jptl (61)
	ugml (68) -> gyxo, ebii, jptl
	gyxo (61)
	cntj (57)`
	stack := makeStack(testInput)
	log.Println(stack.Root.FindUnbalance())
	return stack.Root.Name
}

func (p Puzzle) Phase1() string {
	stack := makeStack(Programs)
	return stack.Root.Name
}

func (p Puzzle) Phase2() string {
	stack := makeStack(Programs)
	log.Println(stack.Root.FindUnbalance())
	return "doin phase2"
}

func (p Puzzle) String() string {
	return "puzz" + id
}

func makeStack(input string) *Stack {
	lines := strings.Split(input, "\n")
	stack := NewStack()
	for _, line := range lines {
		name, weight, children := parseElement(line)
		log.Println(name, weight, children)
		stack.CreateNode(name, weight, children)
	}
	return &stack
}

func parseElement(line string) (string, int, []string) {
	chunks := strings.Fields(line)
	name := chunks[0]
	weight, _ := parseWeight(chunks[1])
	children := make([]string, 0)
	if len(chunks) > 2 {
		for _, rawChild := range chunks[3:] {
			children = append(children, strings.Replace(rawChild, ",", "", -1))
		}
	}
	return name, weight, children
}

func parseWeight(input string) (int, error) {
	input = strings.Replace(input, "(", "", -1)
	input = strings.Replace(input, ")", "", -1)
	return strconv.Atoi(input)
}

func init() {
	base.Register(id, Puzzle{})
}

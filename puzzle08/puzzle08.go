package puzzle08

import (
	"log"
	"strconv"
	"strings"

	"github.com/parafiend/advent2017/base"
)

const id = "8"

type Puzzle struct{}

func (p Puzzle) Test() string {
	input := `b inc 5 if a > 1
	a inc 1 if b < 5
	c dec -10 if a >= 1
	c inc -20 if c == 10`
	registers := process(input)
	return strconv.Itoa(registers[getMaxRegister(&registers)])
}

func (p Puzzle) Phase1() string {
	registers := process(instructions)
	return strconv.Itoa(registers[getMaxRegister(&registers)])
}

func (p Puzzle) Phase2() string {
	return "doin phase2"
}

func (p Puzzle) String() string {
	return "puzz" + id
}

func process(input string) map[string]int {
	registers := make(map[string]int)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		bits := strings.Fields(line)
		if evaluate(&registers, bits[4:]) {
			update(&registers, bits[:3])
		}
	}
	return registers
}

func evaluate(registry *map[string]int, parts []string) bool {
	initialVal := getOrCreateRegister(registry, parts[0])
	condition, _ := strconv.Atoi(parts[2])
	var result bool
	switch parts[1] {
	case ">":
		result = initialVal > condition
	case "<":
		result = initialVal < condition
	case ">=":
		result = initialVal >= condition
	case "<=":
		result = initialVal <= condition
	case "==":
		result = initialVal == condition
	case "!=":
		result = initialVal != condition
	default:
		log.Println(parts)
	}
	log.Println(parts, result)
	return result
}

func update(registry *map[string]int, parts []string) {
	log.Println(registry, parts)
	initialVal := getOrCreateRegister(registry, parts[0])
	delta, _ := strconv.Atoi(parts[2])
	log.Println(delta, parts[1])
	switch parts[1] {
	case "inc":
		(*registry)[parts[0]] = initialVal + delta
	case "dec":
		(*registry)[parts[0]] = initialVal - delta
	default:
		log.Println(parts)
	}
	log.Println(registry)
}

func getOrCreateRegister(registry *map[string]int, name string) int {
	val, found := (*registry)[name]
	if !found {
		(*registry)[name] = val
	}
	return val
}

func getMaxRegister(registers *map[string]int) string {
	var result string
	max := -int(^uint(0)>>1) - 1
	log.Println(max, registers)
	for k, v := range *registers {
		if v > max {
			result = k
			max = v
		}
	}
	return result
}

func init() {
	base.Register(id, Puzzle{})
}

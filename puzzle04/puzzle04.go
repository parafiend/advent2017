package puzzle04

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/parafiend/advent2017/base"
)

const id = "4"

type Puzzle struct{}

func (p Puzzle) Test() string {
	log.Println(isPhraseValid("aa bb cc dd ee"))
	log.Println(isPhraseValid("aa bb cc dd aa"))
	log.Println(isPhraseValid("aa bb cc dd aaa"))
	return "doin a test"
}

func (p Puzzle) Phase1() string {
	var sumValid int
	for _, v := range strings.Split(Phrases, "\n") {
		if isPhraseValid(v) {
			sumValid++
		}
	}
	return strconv.Itoa(sumValid)
}

func (p Puzzle) Phase2() string {
	return "doin phase2"
}

func (p Puzzle) String() string {
	return "puzz" + id
}

func isPhraseValid(phrase string) bool {
	phrases := strings.Fields(phrase)
	sort.Strings(phrases)
	valid := true
	for i, v := range phrases {
		if i < len(phrases)-1 && phrases[i+1] == v {
			valid = false
			break
		}
	}
	return valid
}

func init() {
	base.Register(id, Puzzle{})
}

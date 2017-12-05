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
	var sumValid int
	for _, v := range strings.Split(Phrases, "\n") {
		if isPhraseValid2(v) {
			sumValid++
		}
	}
	return strconv.Itoa(sumValid)
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

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(input string) string {
	runes := []rune(input)
	sort.Sort(sortRunes(runes))
	return string(runes)
}

func isPhraseValid2(phrase string) bool {
	phrases := strings.Fields(phrase)
	for i, v := range phrases {
		x := sortString(v)
		phrases[i] = x
	}
	log.Println(phrases)
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

package puzzle09

import (
	"log"
	"strconv"

	"github.com/parafiend/advent2017/base"
)

const id = "9"

type StreamParser struct {
	GroupCount    int
	InGarbage     bool
	NextCancelled bool
}

func MakeParser() StreamParser {
	return StreamParser{0, false, false}
}

type Puzzle struct{}

func (p Puzzle) Test() string {
	testStrings := [6]string{"{{{},{},{{}}}}",
		"{{<ab>},{<ab>},{<ab>},{<ab>}}",
		"{{<a!>},{<a!>},{<a!>},{<ab>}}",
		"{{<!!>},{<!!>},{<!!>},{<!!>}}",
		"<{o\"i!a,<{i<a>",
		"<<<<>"}
	for _, test := range testStrings {
		score, garbage := ScoreGroups(test)
		log.Println("----", score, garbage)
	}
	return "doin a test"
}

func (p Puzzle) Phase1() string {
	score, garbage := ScoreGroups(input)
	return strconv.Itoa(score) + ", " + strconv.Itoa(garbage)
}

func (p Puzzle) Phase2() string {
	return "See phase1"
}

func (p Puzzle) String() string {
	return "puzz" + id
}

func ScoreGroups(stream string) (int, int) {
	p := MakeParser()
	score := 0
	garbage := 0
	for i, char := range stream {
		if p.InGarbage {
			switch {
			case p.NextCancelled:
				p.NextCancelled = false
			case char == '!':
				p.NextCancelled = true
			case char == '>':
				p.InGarbage = false
			default:
				garbage++
			}
		} else {
			switch char {
			case '{':
				p.GroupCount++
				score += p.GroupCount
				log.Println(i, p.GroupCount)
			case '}':
				p.GroupCount--
			case '<':
				p.InGarbage = true
			}
		}
	}
	return score, garbage
}

func init() {
	base.Register(id, Puzzle{})
}

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
	log.Println("----", ScoreGroups("{{{},{},{{}}}}"))
	log.Println("----", ScoreGroups("{{<ab>},{<ab>},{<ab>},{<ab>}}"))
	log.Println("----", ScoreGroups("{{<a!>},{<a!>},{<a!>},{<ab>}}"))
	log.Println("----", ScoreGroups("{{<!!>},{<!!>},{<!!>},{<!!>}}"))
	log.Println("----", ScoreGroups("<{o\"i!a,<{i<a>"))
	log.Println("----", ScoreGroups("<<<<>"))
	return "doin a test"
}

func (p Puzzle) Phase1() string {
	return strconv.Itoa(ScoreGroups(input))
}

func (p Puzzle) Phase2() string {
	return "doin phase2"
}

func (p Puzzle) String() string {
	return "puzz" + id
}

func ScoreGroups(stream string) int {
	p := MakeParser()
	score := 0
	for i, char := range stream {
		if p.InGarbage {
			switch {
			case p.NextCancelled:
				p.NextCancelled = false
			case char == '!':
				p.NextCancelled = true
			case char == '>':
				p.InGarbage = false
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
	return score
}

func init() {
	base.Register(id, Puzzle{})
}

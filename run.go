package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/parafiend/advent2017/base"
)

import _ "github.com/parafiend/advent2017/puzzle01"
import _ "github.com/parafiend/advent2017/puzzle02"
import _ "github.com/parafiend/advent2017/puzzle03"
import _ "github.com/parafiend/advent2017/puzzle04"

func main() {
	puzzle := flag.Int("puzzle", 0, "Puzzle ID to run")
	mode := flag.String("mode", "test", "Puzzle mode to run: test/1/2")

	flag.Parse()
	fmt.Println("puzz:", *puzzle)
	fmt.Println("mode:", *mode)

	var logic base.Puzzle
	logic = base.Get(strconv.Itoa(*puzzle))

	//TODO: Add nil check
	fmt.Println(logic)

	switch *mode {
	case "test":
		fmt.Println(logic.Test())
	case "1":
		fmt.Println(logic.Phase1())
	case "2":
		fmt.Println(logic.Phase2())
	default:
		fmt.Println("Error: Expected 'test', '1' or '2' for mode option")
	}
}

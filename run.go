package main

import "flag"
import "fmt"

func main() {
	puzzle := flag.Int("puzzle", 0, "Puzzle ID to run")
	mode := flag.String("mode", "test", "Puzzle mode to run: test/1/2")

	flag.Parse()
	fmt.Println("puzz:", *puzzle)
	fmt.Println("mode:", *mode)
}
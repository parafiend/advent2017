package puzzle10

import (
	"fmt"
	"log"
	"strconv"

	"github.com/parafiend/advent2017/base"
)

const id = "10"

type Ring struct {
	Size   int
	Skip   int
	Index  int
	Values []int
}

func MakeRing(size int) Ring {
	values := make([]int, size)
	for i := 0; i < size; i++ {
		values[i] = i
	}
	return Ring{size, 0, 0, values}
}

type Puzzle struct{}

func (p Puzzle) Test() string {
	ring := MakeRing(5)
	swaps := [4]int{3, 4, 1, 5}
	for _, swap := range swaps {
		ring.ProcessSwap(swap)
		//log.Println(ring.Values)
	}

	log.Println(FullHash(""))
	log.Println(FullHash("AoC 2017"))
	log.Println(FullHash("1,2,3"))
	return strconv.Itoa(ring.Values[0] * ring.Values[1])
}

func (p Puzzle) Phase1() string {
	ring := MakeRing(256)
	swaps := [16]int{225, 171, 131, 2, 35, 5, 0, 13, 1, 246, 54, 97, 255, 98, 254, 110}
	for _, swap := range swaps {
		ring.ProcessSwap(swap)
		//log.Println(ring.Values)
	}
	return strconv.Itoa(ring.Values[0] * ring.Values[1])
}

func (p Puzzle) Phase2() string {
	swaps := "225,171,131,2,35,5,0,13,1,246,54,97,255,98,254,110"
	return FullHash(swaps)

}

func (p Puzzle) String() string {
	return "puzz" + id
}

func FullHash(input string) string {
	ring := MakeRing(256)
	extraSwaps := [5]int{17, 31, 73, 47, 23}
	for i := 0; i < 64; i++ {
		for _, swap := range input {
			ring.ProcessSwap(int(swap))
			//log.Println(int(swap), ring.Values)
		}
		for _, swap := range extraSwaps {
			ring.ProcessSwap(swap)
		}
	}
	dense := make([]int, 16)
	output := ""
	for i := 0; i < 16; i++ {
		res := ring.Values[i*16]
		for j := 1; j < 16; j++ {
			res ^= ring.Values[(i*16)+j]
		}
		dense[i] = res
		output += fmt.Sprintf("%02x", res)
	}
	return output
}

func (ring *Ring) ProcessSwap(len int) {
	for i := 0; i < len/2; i++ {
		l := ring.RingifyIndex(ring.Index + i)
		r := ring.RingifyIndex(ring.Index + len - i - 1)
		temp := ring.Values[l]
		ring.Values[l] = ring.Values[r]
		ring.Values[r] = temp
	}
	ring.Index = ring.RingifyIndex(ring.Index + len + ring.Skip)
	ring.Skip++
}

func (r *Ring) RingifyIndex(idx int) int {
	for idx >= r.Size {
		idx -= r.Size
	}
	for idx < 0 {
		idx += r.Size
	}
	return idx
}

func init() {
	base.Register(id, Puzzle{})
}

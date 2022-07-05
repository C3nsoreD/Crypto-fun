package main

import (
	"fmt"
	"sort"
)

type Sequence []int

// required by sort.Interface
// func (s Sequence) Len() int {
// 	return len(s)
// }

// func (s Sequence) Less(i, j int) bool {
// 	return s[i] < s[j]
// }

// func (s Sequence) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

/* Above code not needed when using sort.IntSlice() */

// returns a copy of the sequence
func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

// method for printing -
func (s Sequence) String() string {
	s = s.Copy()
	sort.IntSlice(s).Sort() // IntSlice implements methods that lets us sort the Sequence... reducing code
	return fmt.Sprint([]int(s))
}

func main() {
	s := Sequence{1, 9, 8, 4, 20, 19}
	fmt.Print(s.String(), "\n")
}

package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2, 3, 4}

	// cap=len=4
	// *int -> int pointer
	target := make([]*int, 4)

	for _, s := range src {
		// go < 1.22
		// s := s
		target = append(target, &s)
	}

	// target = changeSlice(target)
	// changeSliceV2(&target)

	prettyPrint(target)
}

func prettyPrint(ss []*int) {
	for _, s := range ss {
		if s != nil {
			fmt.Println(*s)
		} else {
			fmt.Println("nil")
		}
	}
}

func changeSlice(ss []*int) []*int {
	lastElement := 0
	for _, s := range ss {
		if s != nil {
			lastElement += *s
		}
	}

	return append(ss, &lastElement)
}

func changeSliceV2(ss *[]*int) {
	if ss == nil {
		return
	}

	lastElement := 0
	for _, s := range *ss {
		if s != nil {
			lastElement += *s
		}
	}

	*ss = append(*ss, &lastElement)
}

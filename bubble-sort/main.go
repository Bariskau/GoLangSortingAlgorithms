package main

import (
	"fmt"
	"math/rand"
	"time"
)

type list []int

func main() {
	s := makeRange(0, 200)
	s.shuffle()
	s.bubbleSort()

	fmt.Println(s)

}

func (s list) bubbleSort() {
	length := len(s)

	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if s[j] > s[j+1] {
				t := s[j]
				s[j] = s[j+1]
				s[j+1] = t
			}
		}
	}
}

func makeRange(min int, max int) list {
	var duration []int

	for i := min; i <= max; i++ {
		duration = append(duration, i)
	}

	return duration
}

func (s list) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range s {
		newPosition := r.Intn(len(s) - 1)
		s[index], s[newPosition] = s[newPosition], s[index]
	}
}

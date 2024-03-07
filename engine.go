package main

import (
	"fmt"
	"math"
)

type searchLimits struct {
	depth    int
	nodes    uint64
	movetime int
	infinite bool
}

var limits searchLimits

func (s *searchLimits) init() {
	s.depth = 64
	s.nodes = math.MaxUint64
	s.movetime = 0
	s.infinite = false
}

func (s *searchLimits) setDepth(d int) {
	s.depth = d
}

func (s *searchLimits) setMoveTime(t int) {
	s.movetime = t
}

func engine() (toEngine chan string, fromEngine chan string) {
	toEngine = make(chan string)
	fromEngine = make(chan string)
	go func() {
		defer close(toEngine)
		for cmd := range fromEngine {
			switch cmd {
			case "stop":
				// Handle stop command
			case "quit":
				// Handle quit command
			case "go":
				// Implement your search algorithm here
			default:
				fmt.Println("Unknown command:", cmd)
			}
		}
	}()
	return toEngine, fromEngine
}

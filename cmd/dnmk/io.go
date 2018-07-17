package main

import (
	"os"
	"fmt"
)

const (
	Minimal = iota
	Normal = iota
	Verbose = iota
)

type IO struct {
	wantedLevel int
}

func (io *IO) Print(level int, text string) {
	if io.wantedLevel >= level {
		fmt.Println(text)
	}
}

func (io *IO) Error(level int, err error) {
	if io.wantedLevel >= level { 
		fmt.Fprintln(os.Stderr, err)
		if level == Minimal {
			os.Exit(1)
		}
	}
}

func NewIO(wantedLevel int) *IO {
	return &IO{
		wantedLevel: wantedLevel,
	}
}
package main

import "fmt"

type file struct {
	name string
}

func (f *file) print(ind string) {
	fmt.Println(ind + f.name)
}

func (f *file) clone() node {
	return &file{name: f.name + "_clone"}
}

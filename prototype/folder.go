package main

import "fmt"

type folder struct {
	childrens []node
	name      string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.childrens {
		i.print(indentation + "  ")
	}
}

func (f *folder) clone() node {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildrens []node
	for _, i := range f.childrens {
		copy := i.clone()
		tempChildrens = append(tempChildrens, copy)
	}
	cloneFolder.childrens = tempChildrens
	return cloneFolder
}

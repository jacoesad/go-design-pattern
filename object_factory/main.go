package main

import "github.com/jacoesad/go-design-pattern/object_factory/factory"

func main() {
	student := factory.GetStudent("Science")
	student.Call()
}

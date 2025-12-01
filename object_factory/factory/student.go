package factory

import "fmt"

type Student interface {
	Call()
}

func GetStudent(subject string) Student {
	switch subject {
	case "Science":
		return &ScienceStudent{}
	case "Commerce":
		return &CommerceStudent{}
	case "Arts":
		return &ArtStudent{}
	default:
		fmt.Println("Invalid Subject given")
		return nil
	}
}

type ScienceStudent struct{}

func (s *ScienceStudent) Call() {
	fmt.Println("This is a science student")
}

type ArtStudent struct{}

func (s *ArtStudent) Call() {
	fmt.Println("This is an arts student")
}

type CommerceStudent struct{}

func (s *CommerceStudent) Call() {
	fmt.Println("This is a commerce student")
}

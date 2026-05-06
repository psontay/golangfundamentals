package practice

import "fmt"

type Animal interface {
	Sound() string
	Name() string
}

type Dog struct {
}

func (d Dog) Sound() string {
	return "Dog sound"
}

func (d Dog) Name() string {
	return "Dog"
}

type Cat struct{}

func (c Cat) Sound() string {
	return "Cat sound"
}

func (c Cat) Name() string {
	return "Cat"
}

type Duck struct{}

func (d Duck) Sound() string {
	return "Duck sound"
}

func (d Duck) Name() string {
	return "Duck"
}

func printInfoAnimal(animals []Animal) {
	for _, a := range animals {
		fmt.Println(a.Sound())
		fmt.Println(a.Name())
	}
}

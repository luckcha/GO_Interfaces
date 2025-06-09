package main

import "fmt"

type Ability interface {
	Speak() string
	Eat() string
}

type Animal struct {
	Name string
	food string
}

func (A Animal) Speak() string {
	return fmt.Sprintf("%s says something", A.Name)
}

func (A Animal) Eat() string {
	return fmt.Sprintf("%s eats %s", A.Name, A.food)

}

func Allability(Ab Ability) {
	fmt.Println(Ab.Speak())
	fmt.Println(Ab.Eat())
}

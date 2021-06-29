package main

import "fmt"

type Animal interface {
	Sleeping()
	Eating()
}

type Cat struct {
	color string
}
func (c *Cat) Eating() {
	fmt.Println(c.color," cat is eating")
}
func (c *Cat) Sleeping() {
	fmt.Println(c.color," cat is sleeping")
}

type Dog struct {
	color string
}
func (d *Dog) Eating() {
	fmt.Println(d.color," dog is eating")
}
func (d *Dog) Sleeping() {
	fmt.Println(d.color," dog is sleeping")
}

func main(){
	c1 := Cat{"white"}
	d1 := Dog{"black"}
	var a1 Animal
	a1 = &c1
	a1.Sleeping()
	a1 = &d1
	a1.Eating()
}

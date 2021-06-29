package main

import "fmt"

type Person struct{
	name string
	age int
	sex string
	fight int
}
type Superman struct{
	strength int
	speed int
	Person
}

func (p *Person) setAge(age int){
	p.age = age
}

func (s *Superman) print(){
	fmt.Printf("%+v\n",s)
}
func main(){
	p1 := Person{"Poor man",30,"man",5 }
	s1 := Superman{100000, 19000, p1 }
	s1.setAge(50)
	s1.print()
}


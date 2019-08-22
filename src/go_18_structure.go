package main

import "fmt"

type Person struct {
	name  string
	age   int
	adult bool
}

type Animal struct {
	name string
	age  int
}

func main() {

	fmt.Println(Person{name: "foo", age: 1})
	fmt.Println(&Person{name: "foo", age: 1})
	fmt.Println(&Person{name: "foo", age: 1})

	person := Person{name: "bar", age: 23, adult: true}
	fmt.Println("========================get")
	fmt.Println(&person)
	fmt.Println(person.name)
	fmt.Println(person.age)

	person.name = "bar2222"
	person.age = 231

	fmt.Println("========================set")
	fmt.Println(&person)
	fmt.Println(person.name)
	fmt.Println(person.age)

	p1 := &person
	p1.name = "fooxxx"
	fmt.Println("========================pointer")
	fmt.Println(&p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

}

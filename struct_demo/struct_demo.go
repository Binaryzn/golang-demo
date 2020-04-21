package struct_demo

import (
	"aweProject/interface_demo"
	"fmt"
)

type Aniaml struct {
	Color string
}

type Dog struct {
	Aniaml
	Name string
	Age  int
	Id   int
}

type Cat struct {
	Aniaml
	Name string
	Age  int
	Id   int
}

func TestForStruct() {

	/*	var dog Dog
		dog.Age = 1
		dog.Id =2
		dog.Name = "tom"*/

	//	dog := Dog{Name: "ddd",Age: 12,Id: 16}

	dog := new(Dog)
	dog.Age = 12
	dog.Id = 22
	dog.Name = "tom"
	fmt.Println("dog is : ", dog)

}

func Action(b interface_demo.Behavir) string {
	b.Run()
	b.Eat()
	return ""
}

func (dog *Dog) Run() string {
	dog.Age = 66
	return "dog is run"

}

func (an *Dog) Eat() string {
	fmt.Println("dog is eating")
	return "dog is  eat"
}

func (dog *Cat) Run() string {
	dog.Age = 66
	return "cat is run "

}

func (an *Cat) Eat() string {
	fmt.Println("cat is eating")
	return "cat is eat "
}

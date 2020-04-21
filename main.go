package main

import (
	"aweProject/struct_demo"
	"fmt"
	"reflect"
)

func main() {
	//makeSlice()
	//makeMap()
	//makeChan()
	//	NewMap()
	//appendElementForSlice()
	//copyForSlice()
	//	deletFromMap()
	//receivePanic()

	//struct_demo.TestForStruct()

	//	dog := new(struct_demo.Dog)
	//	dog.Age = 15
	//	dog.Id = 16
	//	dog.Name = "ddd"
	//	dog.Color = "red"
	//	 dog.Run()
	//	dog.Eat()
	//	fmt.Println("i is",dog.Age)
	//var d interface_demo.Behavir
	d := new(struct_demo.Dog)
	d.Eat()
	d.Run()
	b := new(struct_demo.Cat)
	b.Eat()
	b.Run()
}
func receivePanic() {
	defer coverPanic()
	panic("i am panic")

}

func coverPanic() {
	message := recover()
	fmt.Println(message)
}

func deletFromMap() {
	myMap := make(map[int]string)
	myMap[0] = "dog"
	myMap[1] = "id-2"
	delete(myMap, 0)
	fmt.Println(myMap)
	delete(myMap, 1)
	fmt.Println(myMap)
}

func appendElementForSlice() {
	myIdSlice := make([]string, 2)
	myIdSlice[0] = "id-1"
	myIdSlice[1] = "id-2"
	myIdSlice = append(myIdSlice, "id-4")
	fmt.Println(myIdSlice)
	fmt.Println("len is :", cap(myIdSlice))
	fmt.Println("len is : ", len(myIdSlice))

}

func copyForSlice() {
	myIdSliceDst := make([]string, 2)
	myIdSliceDst[0] = "iddst-1"
	myIdSliceDst[1] = "iddst-2"
	myIdSliceSrc := make([]string, 2)
	myIdSliceSrc[0] = "idsrc-1"
	myIdSliceSrc[1] = "idsrc-2"
	copy(myIdSliceDst, myIdSliceSrc)
	fmt.Println(myIdSliceDst)
}

func NewMap() {
	mMap := new(map[int]string)

	makeMap := make(map[int]string)
	makeMap[0] = "ddd"
	fmt.Println("new map is :", reflect.TypeOf(mMap))
	fmt.Println("make map is :", reflect.TypeOf(makeMap))

}

func makeSlice() {
	mSlice := make([]string, 3)
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "tiger"
	fmt.Println(mSlice)
}

//创建map
func makeMap() {
	mMap := make(map[int]string)
	mMap[11] = "dog"
	mMap[100] = "cat"
	fmt.Println(mMap)
}

func makeChan() {
	mChan := make(chan int)
	close(mChan)
}

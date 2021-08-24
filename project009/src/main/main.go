// 接口
package main

import "fmt"

// 接口是方法的集合
// 一但struct实现了全部的方法，则接口实现
// 接口本身没有实例，实质上依靠实现接口的对象
// 即成为一个指向struct的指针


func main() {

	s := student{"ca",18}
	w := worker{"hcw",35}

	s.day(s)
	fmt.Println()
	// 等效于
	var inter business = s
	s.day(inter)
	fmt.Println()
	// 调用student的接口
	s.day(w)
	// 调用worker的接口




}


type business interface {
	do()
	eat(int)
	get()int
}


type worker struct {
	name string
	age int
}
func (a worker)do()  {
	fmt.Printf("worker say that he is %s and %d years old.\n",a.name,a.age)
}
func(a worker) eat(x int)  {
	fmt.Printf("worker %s eat %d kg.\n",a.name,x)
}
func (a worker)get() int {
	return 100
}
func (a worker)day(inter business)  {
	inter.do()
	inter.eat(12)
	fmt.Printf("get %d\n",inter.get())
}


type student struct {
	name string
	age int
}
func (a student)do()  {
	fmt.Printf("student say that he is %s and %d years old.\n",a.name,a.age)
}
func(a student) eat(x int)  {
	fmt.Printf("student %s eat %d kg.\n",a.name,x)
}
func (a student)get() int {
	return -10
}
func (a student)day(inter business)  {
	inter.do()
	inter.eat(2)
	fmt.Printf("get %d\n",inter.get())
}



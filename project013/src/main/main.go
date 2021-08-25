// 多态

package main

import "fmt"

func main() {

	s := student{"ca",18}
	w := worker{"hcw",35}


	// 1.多态参数
	// 方法随传入的实例的类型而改变
	s.day(s)
	s.day(w)

	// 2.多态数组
	// 某接口数组/切片可以容纳任何一个实现该接口的对象
	li := make([]business,4)
	li[0] = student{"ca",18}
	li[1] = worker{"zlc",60}
	li[2] = student{"zlh",18}
	li[3] = worker{"hcw",40}

	fmt.Println(li)
}


type business interface {
	do()
	eat(int)
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
func (a worker)day(inter business)  {
	inter.do()
	inter.eat(12)
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
func (a student)study() {
	fmt.Println("i'm studying")
}
func (a student)day(inter business)  {
	inter.do()
	inter.eat(2)
}



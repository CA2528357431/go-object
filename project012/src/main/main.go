// 类型断言

package main

import "fmt"

func main() {

	// struct-->interface
	// 可用接口中的方法

	w := worker{"hcw",35}
	var p business
	p = w
	p.eat(10)

	// interface-->struct
	// 由于interface宽泛，无法确认其来自于哪个struct
	// 使用类型断言
	// 若stu也实现接口，ww = p.(stu)则会报错
	var ww worker
	ww = p.(worker)
	fmt.Println(ww)




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
func(a worker) get(x int)  {
	fmt.Printf("worker get %d.\n",x)
}



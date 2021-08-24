// 有传址方法的接口

// 接口应定义为指向对象的指针

// 如果函数是传值，可传入指针或者原值；如果函数是传地址，只能传址
// 对象只能调用对象的方法，指针可以调用对象和指针的方法


package main

import "fmt"

func main() {

	w := worker{"hcw",35}

	var in valuebusiness = &w
	w.day(in)
	fmt.Println()
	// 等价于
	w.day(&w)
}


type valuebusiness interface {
	do()
	get()int
	eat(int)
}


type worker struct {
	name string
	age int
}
func (a *worker)do()  {
	fmt.Printf("worker say that he is %s and %d years old.\n",a.name,a.age)
}
func(a worker) eat(x int)  {
	fmt.Printf("worker %s eat %d kg.\n",a.name,x)
}
func (a worker)get() int {
	a.age+=1
	return 100
}
func (a worker)day(inter valuebusiness)  {
	inter.do()
	inter.eat(12)
	fmt.Printf("get %d\n",inter.get())
}






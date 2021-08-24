// 接口继承

// 想实现子接口必须先实现父接口
// 多继承同理不做演示

package main

import "fmt"

func main() {

	w := worker{"hcw",35}
	w.day(w)
}

type originbusiness interface {
	eat(int)

}

type modernbusiness interface {
	originbusiness
	do()
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
func (a worker)day(inter modernbusiness)  {
	inter.do()
	inter.eat(12)
	fmt.Printf("get %d\n",inter.get())
}






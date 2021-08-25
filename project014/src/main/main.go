// 利用类型断言根据类型执行

// 混合数组中，student执行study，worker执行work

package main

import "fmt"

func main() {

	li := make([]business,4)
	li[0] = student{"ca",18}
	li[1] = worker{"zlc",60}
	li[2] = student{"zlh",18}
	li[3] = worker{"hcw",40}

	// 为何传地址？
	// 指针默认类型位nil，便于和非默认实例区分
	for _,obj := range li{
		func(x business){
			w := workertran(x)
			s := studenttran(x)
			if w != nil {
				w.work()
			}
			if s != nil{
				s.study()
			}
		}(obj)
	}



}

func check() {
	err := recover()
	if err!=nil{

	}
}
func workertran(x business)*worker {
	defer check()
	neo := x.(worker)
	return &neo
}
func studenttran(x business)*student {
	defer check()
	neo := x.(student)
	return &neo
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
func (a worker)work() {
	fmt.Println("i'm working")
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





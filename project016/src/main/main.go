// 区分指针和值

// 指针型通过类型断言改变原值

package main

import "fmt"

func main() {

	li := [2]business{}
	li[0] = worker{"ca",18}
	li[1] = &worker{"zlc",60}
	for _,sth := range li{
		switch sth.(type) {
		case worker:
			x := sth.(worker)
			x.name="sth"
		case *worker:
			x := sth.(*worker)
			x.name="parser"
		}
	}
	fmt.Println(li[0])
	fmt.Println(li[1])

}

type business interface {
}

type worker struct {
	name string
	age int
}








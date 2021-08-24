// 对象方法

package main

import "fmt"

func main()  {

	neo := per{"caoan",18}
	neo.say()
}

type per struct {
	name string
	age int
}

func (a per)say()  {
	fmt.Printf("i'm %v, %v years old\n",a.name,a.age)

}
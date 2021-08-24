// 对象方法

package main

import "fmt"

func main()  {

	neo := per{"caoan",18}
	neo.say()
	fmt.Println(neo.iq(9999999))
}

type per struct {
	name string
	age int
}

func (a per)iq(wealth int)int  {
	return a.age*wealth

}

func (a per)say()  {
	fmt.Printf("i'm %v, %v years old\n",a.name,a.age)

}
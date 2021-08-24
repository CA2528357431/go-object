// String方法

package main

import "fmt"

func main()  {

	neo := per{"caoan",18}
	fmt.Println(neo)
	// 有 string 方法时，若输出自动调用

}

type per struct {
	name string
	age int
}

func (a per)String()string  {
	return fmt.Sprintf("%s: %d",a.name,a.age)

}
// 基础类型修饰

package main

import "fmt"

func main()  {
	var a intneo = 2
	b := a.show()
	fmt.Println(b)
}

type intneo int

func (a intneo)show() intneo{

	fmt.Println(a)
	var b intneo = a+1
	return b
}


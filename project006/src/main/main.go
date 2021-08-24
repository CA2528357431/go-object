// 工厂模式

// struct私有怎么办
// 创建共有的构造函数

// 属性私有怎么办
// 创造共有的get/set函数

package main

import (
	"fmt"
	"stu"
)

func main()  {
	neo := stu.Createstu("caoan",18)
	fmt.Println(neo)

	neo.SetAge(16)

	age := neo.GetAge()
	fmt.Println(age)
}

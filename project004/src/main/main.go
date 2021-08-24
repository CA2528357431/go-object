// 传针对象方法

// 可以提高效率节约内存，尤其是struct很大的时候
// 对于struct，调用各种方法，通过语法糖，无需区分指针和值的表达
// 不过输出的时候还是要区分

// 如果函数是传值，可传入指针或者原值；如果函数是传地址，只能传址
// 对象只能调用对象的方法，指针可以调用对象和指针的方法

package main

import (
	"fmt"
)

func main()  {

	neo := per{"caoan",18}
	fmt.Println(neo.iq(9999999))
	fmt.Println(neo)

	fmt.Println()

	fmt.Println(neo.neoiq(9999999))
	fmt.Println(neo)

	fmt.Println()

	fmt.Println((&neo).iq(9999999))
	fmt.Println(neo)
	// (&neo).iq()效果同neo.iq()

	fmt.Println()

	this := createper("ca",99)
	fmt.Println(this.iq(9999999))
	fmt.Println(*this)
}

type per struct {
	name string
	age int
}

func (a per)iq(wealth int)int  {
	a.age+=10
	return a.age*wealth

}
func (a *per)neoiq(wealth int)int  {
	a.age+=10
	return a.age*wealth
}
// 改变原结构体
// 使用时直接把指针名当作对象
// 如(*a).age，可写作a.age

func createper(name1 string,age1 int)*per  {
	neo := per{name: name1,age: age1}
	return &neo
}



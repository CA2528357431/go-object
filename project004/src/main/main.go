// 传针对象方法

// 可以提高效率节约内存，尤其是struct很大的时候，

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
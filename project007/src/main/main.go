// 继承与重写

package main

import "fmt"

func main() {

	p := per{
		name: "xx",
		age:  99,
	}
	p.say1()
	p.say2()

	s := stu{
		per:  per{"caoan",18},
		mark: 99,
	}
	s.say1()
	s.say2()
	s.per.say2()
	// s.per正是s在per的基础
	// per也类似于属性
	// 相当于内置一个per实例


}

type per struct {
	name string
	age int
}

func (a per)say1()  {
	fmt.Printf("per say that he is %s and %d years old. by say11111111\n",a.name,a.age)
}

func (a per)say2()  {
	fmt.Printf("per say that he is %s and %d years old. by say22222222\n",a.name,a.age)
}



type stu struct {
	per
	mark int
}

// 多态
func (a stu)say2()  {
	fmt.Printf("stu say that he is %s and %d years old and %d mark. by say22222222\n",a.name,a.age,a.mark)
}
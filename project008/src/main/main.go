// 多继承

// 继承多个，如果有 func/属性 一样，则必须加上级对象名

// go继承的本质是设置一个匿名属性，直接用类型名当作属性名
// 如果struct还有对应类型的其他属性，则不可以全都匿名


package main

import "fmt"

func main() {
	s := son{
		fa: fa{"faname",44},
		ma: ma{"maname",35},
		name: "ca",
		age: 18,
	}
	s.fa.say()
	s.ma.say()
	fmt.Println(s.name)
	fmt.Println(s.fa.name)
	fmt.Println(s.ma.name)


	l := les{
		a:    ma{"kt",18},
		ma:   ma{"mg",17},
		name: "hk",
		age:  3,
	}
	l.ma.say()

	l0 := les0{
		a:    ma{"kt",18},
		b:   ma{"mg",17},
		name: "hk",
		age:  3,
	}
	l0.a.say()
}

type fa struct {
	name string
	age int
}
func (a fa)say()  {
	fmt.Printf("fa say that he is %s and %d years old. by fa\n",a.name,a.age)
}

type ma struct {
	name string
	age int
}
func (a ma)say()  {
	fmt.Printf("ma say that he is %s and %d years old. by ma\n",a.name,a.age)
}


type son struct {
	fa
	ma
	name string
	age int
}

type les struct {
	a ma
	ma
	name string
	age int
}

type les0 struct {
	a ma
	b ma
	name string
	age int
}


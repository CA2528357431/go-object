package stu

type stu struct {
	name string
	age int
}

func (a *stu) GetAge() int {
	return a.age
}

func (a *stu) SetAge(age int) {
	a.age = age
}

func (a stu) iq (wealth int)int  {
	res := a.age+wealth*2
	return res
}

func Createstu(name1 string,age1 int)stu  {
	neo := stu{name: name1,age: age1}
	return neo
}

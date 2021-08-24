package stu

type stu struct {
	name string
	age int
}

func (a stu) iq (wealth int)int  {
	res := a.age+wealth*2
	return res
}

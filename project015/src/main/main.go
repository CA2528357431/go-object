// 类型断言基础数据类型


package main

import "fmt"

func main() {

	// 推广判断基本类型
	list := make([]interface{},3)
	list[0] = 10
	list[1] = 6.4
	list[2] = "yep"
	for _,value := range list{
		if sth,err := value.(int);err==true{
			fmt.Println(sth+100)
		}
		if sth,err := value.(float64);err==true{
			fmt.Println(sth+9.9)
		}
		if sth,err := value.(string);err==true{
			fmt.Println(sth+"111")
		}
	}

	fmt.Println()

	for _,value := range list{
		switch value.(type) {
		case int:
			x:=value.(int)
			fmt.Println(x+100)
		case float64:
			x:=value.(float64)
			fmt.Println(x+9.9)
		case string:
			x:=value.(string)
			fmt.Println(x+"111")
		}
	}

	fmt.Println()


}







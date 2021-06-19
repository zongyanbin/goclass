/**
接口 interface
定义
接口是一类规范，是某一些方法的集合
*/

package main
import "fmt"
// 定义接口
type Animal interface {
	Run()
	Eat()
}

// 结构体
type Cat struct {
	Name string
	Sex bool
}

// 结构体 使用(实现)
type Dog struct {
	Name string
}

type Dogs struct {
	Name string
}

func (c Cat)Run(){
	fmt.Println(c.Name,"开始炮")
}

func (c Cat)Eat(){
	fmt.Println(c.Name,"开始吃")
}

func (d Dog)Run()  {
	fmt.Println(d.Name,"开始跑")
}

func (d Dog)Eat(){
	fmt.Println(d.Name,"开始吃")
}

// 解耦合  我不知道他是哪个结构体但是我想让他做的事情是固定的？
//所以l必须跟着声明好的接口
var L Animal
func main(){

	c:=Cat{
		Name: "TTTT",
		Sex:  false,
	}

	//MyFun(c)
	//L.Run()
	MyFunt(c)
}

// 不想固定参数类型想什么都可以传
func MyFun(a Animal)  {  // 形参,空接口
	//a.Run()
	//L.Eat()
	L=a
}

func MyFunt(l Animal) {
	l.Run()
	l.Eat()
}
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

func main()  {
	// 声明一个接口
	var a Animal

	// 声明一个c结构体 本人符合机构规范
	c:=Cat{
		Name: "Tom",
		Sex:false,
	}

	a = c
	// 通过接口调用方法，会定义到原始实例上的放
	a.Eat()
	a.Run()

}
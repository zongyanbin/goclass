/**
golang 结构体struct声明与使用

使用
	声明
	var qm QM
	qm:=QM{}
	qm:=new(QM)

	访问参数
	qm.xxx
	作为方法参数
	func A(QM){}

	结构体指针
	var qp *QM

声明
	type 结构体名 struct{其他数据类型 数据类型标签}
 */
package main

import "fmt"
//结构体像json json套用json结构体也一样
// 声明一个Qimiao的结构体
type Qimiao struct {
	Name string
	Age int
	Sex bool
	Hobby []string
	//MyHome Home
	Homes Home
}
type Home struct {
	p string
}

func (h *Home)Open(){
	fmt.Println("Open",h.p)
}

// 结构体声明方法
//结构体放在前面，方法名，入参出参，返回
func (q *Qimiao)Song(name string)(rester string){
	rester = "111"
	fmt.Printf("%v唱了一首%v,观记得%v",q.Name,name,rester)
	return  rester
//	return rester
}

func main(){
	// 声明使用3中方式
	// 第一种
	//var qm Qimiao
	//qm.Name="QIMIAO"
	//qm.Age=18
	//qm.Sex=true
	//qm.Hobby=[]string{
	//	"音乐",
	//	"看电影",
	//	"网游戏",
	//}

	//掩饰声明 第二种
	//qm :=Qimiao{
	//	"奇妙",13,true,[]string{"听歌","看电影"},
	//}

	//掩饰声明 第三种
	//qm :=Qimiao{
	//	Name:  "aaa",
	//	Age:   12,
	//	Sex:   false,
	//	Hobby:[]string{"唱歌"},
	//}


	//声明空值 掩饰
	//var qm Qimiao
	//qm:=Qimiao{
	//	Name:  "",
	//	Age:   0,
	//	Sex:   false,
	//	Hobby: nil,
	//}

	//声明空值 Qimiao结构体 那到结构体可以直接使用 例如 qm.name new出来的是地址直接用
	//qm:=new(Qimiao)


	// 把结构体当作一个参数类型传进去,入参和出参
	//qmFunc(qm)
	//fmt.Println(qm.Name)

	//结构体的指针怎么声明
	//结构体可以直接通过内存地址来操作

	//* & 适用于复杂的数据类型 结构体 数组 map 直接取地址就可以用
//	var qmp *Qimiao
//	qmp = &qm
////	qmp.Name="指针修改"
//	(*qmp).Name="奇妙修改过" //用到时候不多
//	fmt.Println(qm, *qmp)

	// 结构体有自己的方法
	// 声明Qimao实例
	qm:=Qimiao{
		Name:  "qimiao",
		Age:   18,
		Sex:   true,
		Hobby: []string{"唱歌","跳舞"},
	}

	//qm.MyHome.p = "北京"
	//qm.Home.p = "北京AA"
	//qm.Open()

	//别名调用
	qm.Homes.p = "北京别名"
	qm.Homes.Open()
	fmt.Println(qm)

	//re :=qm.Song("惊雷")
	//fmt.Println("\n",re)
}

// 把结构体当作一个参数类型传进去
func qmFunc(qm Qimiao)  {
	fmt.Println(qm)
}

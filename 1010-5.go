/**
数组
 */
package main

import "fmt"

func main(){
	//[元素长度]元素类型{元素1，元素2，元素3...}
	a:=[3]int{0,1,2}
	b:=[...]int{1,2,3,4,5,6,7,8,9,0}

	var c = new([10]int)
	c[5]=3
	fmt.Println(a,b,c)

	//为什么要使用数组
	zoom :=[...]string{"狗子","猫跑","猴子跑"}

	//数组的利用和数组的循环
	for i:=0;i<len(zoom);i++ {
		fmt.Println(zoom[i]+"正在跑")
	}

	// for range
	for i,v := range zoom{
		fmt.Println(i,v+"正在跑")
	}

	//len cap 数组长度和容量一样的
	fmt.Println(len(zoom))
	fmt.Println(cap(zoom))

	//多维数组
	err := [3][3]int{
		{0,1,2},
		{1,2,3},
		{2,3,4},
	}

	fmt.Println(err)
	//[
	//	[0,1,2],
	//	[3]int{0,1,2}
	//	[1,2,3],
	//	[3]int{1,2,3}
	//	[2,3,4],
	//	[3]int{2,3,4}
	//]

	//切片是数组的一部分
	cl := a[2]
	c2 := a[2:]
	fmt.Println(cl,c2)

	c2[0] = 5
	fmt.Println(c2)
	fmt.Println(a)

	//怎么拿切片，如何拿切片
	d := [3]int{0,1,2}
	d1 :=d[1:]
	fmt.Println(d1)
	d1 = append(d1,5)
	d1[0] =4

	fmt.Println(d1)

	//切片前闭后开原则 声明
	var d2 []int

	 d22 := make([]int,5)
	 fmt.Println(d2,d22)

	 //第六章 map 使用和声明 字典 hash表  key value

	 //声明方式三种
	 //var m  map[string]string
	 //m = map[string]string{} //注意第一种需要赋值
	 //m["name"]="奇妙"
	 //m["sex"] = "男"
	//fmt.Println(m)

	 //m1 :=map[string]string{}
	 //m1["name"]="奇妙"
	 //m1["sex"] = "男"
	//fmt.Println(m1)

	 //m2 :=make(map[string]string)
	 //m2["name"]="奇妙"
	 //m2["sex"] = "男"
	 //fmt.Println(m2)

	//m3:=map[int]bool{}
	//m3[1]=true
	//m3[2]=false
	//fmt.Println(m3)

	m4 :=map[string]interface{}{} //空接口可以接受任何数据类型  数据类型要统一 map
	m4["a"]=1
	m4["b"]=false
	m4["c"]="str"
	m4["d"]=[]int{1,2,3,4}

	//使用 删除 delete(变量名,key) delete(m,"header")
	//delete(m4,4)
	fmt.Println(m4)
	for k,v := range m4{
		fmt.Println(k,v)
	}

	// 普通 函数方法 func
	func_a(0,"现在小于1")

	// 匿名函数
	n := func(data1 string) {
		fmt.Println(data1)
	}
	n("我是第一个匿名函数")


	ar:=[]string{"1","2","3","4","5"}
	// 不定项参数
	func_b(9527,ar...)


	//自执行函数
	(func(){
		fmt.Println("我在这里执行,别人都没管我")
	})()

	//闭包函数  一个函数返回一个函数是闭包函数
    func_c()(4)

	//defer 延迟调用
	defer  func_d()
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")

}
// main里不能直接写函数定义方法需要在函数外写
// 普通函数的声明
func func_a(data1 int,data2 string)(ret1 int,ret2 string){
	// 没有在函数内部定义 ret1 ret2
	ret1 = data1
	ret2 = data2
	return   //什么都没有写 省略
}

// 不定项参数,不管入参还是出参一定要放到最后一位
func func_b(data1 int,data2 ...string)  {
	fmt.Println(data1,data2)
	for k,v :=range data2{
		fmt.Println(k,v)
	}
}

//闭包函数外面怎么声明里
//里面就需要怎么声明
//返回一个函数
func func_c()func(int){
	return  func(num int){
		fmt.Println("我是闭包函数",num)
	}
}

//defer 延迟调用
func func_d()  {
	fmt.Println("我想最先执行")
}

// 指针和地址 可以隐式也可以显示


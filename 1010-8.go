/**
指针 不能指向具体的数据和值
指针指向和他类型一样的地址

指针指向一个变量的原始内存地址，当用*就表示修改原始地址上的值，&表示取地址， *声明指针

我们想在任何地方改变原始位置，那我就需要用指针去取到原始位置的地址并且对他进行修改

数组指针和指针数组
 */
package main

import "fmt"

func main()  {
	var a int
	a = 123
	fmt.Println(a)
	var b *int
	b=&a  // b指针指向a的地址  b内存地址
	*b = 999 // 哪内存地址 *b
	//b=456
	fmt.Println(a,b,*b)
	fmt.Println(a == *b,&a==b)

	// 数组指针
	var arr [5]string
	arr = [5]string{"1","2","3","4"}
	var arrP *[5]string
	arrP = &arr
	fmt.Println(arr,arrP)

	//指针数组
	var arrpp [5]*string
	var str1 string = "str1"
	var str2 string = "str2"
	var str3 string = "str3"
	var str4 string = "str4"
	var str5 string = "str5"
	arrpp = [5]*string{&str1,&str2,&str3,&str4,&str5}
	*arrpp[2] = "555"
	fmt.Println(str3)
	

	//指针传参  我想改变func外面的东西那么就需要用指针
	// 不想改变外面东西，外面值进来用一圈丢回去，那就要传一个普通的类型
	var str1s = "我定义了"
	pointFun(&str1s)
	fmt.Println(str1s)


	var str6 string = "我是来测试地址的"
	 sp := &str6
	*sp = "我是来测试地址的123123"
	fmt.Println(sp,*sp)
}
func pointFun(p1 *string)  {
	*p1 = "我们变了"
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

 // 全局数组初始化
 var arr0 [5]int = [5]int{1,2,3,4,5}
 var arr1 = [5]int{1,1}
 var arr2 = [...]int{1,2,3,4,5,6,7,8,8}
 var str = [...]string{1:"adaf",3:"adefds"}
var str1 = [5]string{3:"adaf",4:"adefds"}
var arr3 = [...]int{1:1,0:2}
func main () {
	// 局部初始化
	a := [5]int{1, 2, 3}
	b := [4]int{1, 1}
	c := [...]int{1, 2, 3}
	d := [6]string{2: "aaaa", 4: "aaaa"}
	fmt.Println(arr0, arr1, arr2, str, str1, arr3)
	fmt.Println("\n\r")
	fmt.Println(a, b, c, d)

	//声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("空切片")
	} else {
		fmt.Println("不是空切片")
	}
	fmt.Println("----ss4-----")
	var ss4 []int = make([]int, 4, 5)
	fmt.Println(ss4)
	fmt.Println("-----ss4----")
	s2 := []int{}
	var s3 []int = make([]int, 0)
	var s4 []int = make([]int, 1, 1)

	test := [5]int{1,2,3,4,5}
	var res []int
	res=test[1:2]
	fmt.Println(res)

	fmt.Println(s2, s3, s4)
	fmt.Printf("全局变量：s4 %v\n", s2, s3, s4)

	println("-------slic-----")
	slice := []int{0,1,2,3,4,5,6,7,8,9}
	d1 := slice[6:8]
	fmt.Println(d1)

	sliced := [][]int{
		[]int{2,3,4},
		[]int{4,4,4},
	}

	fmt.Println(sliced)

	strz :="hellow world"
	sa :=strz[0:5]
	fmt.Println(sa)

	data_slice := [...]int{0,1,2,3,4,5,6,7,9}
	for index,value := range data_slice{
		fmt.Println("index=",index,"vaule=",value)
	}

	var a001 = []int{1,3,4,5}
	 bs:=a001[2:3]
	 fmt.Println(bs)

	// var a002 = []int{2,3}

	arrayA := [2]int{100, 200}
	//testArrayPoint(&arrayA)   // 1.传数组指针
	arrayB := arrayA[:]
	testArrayPoint(&arrayB)   // 2.传切片
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)

	aas:=10
	bb:=&aas
	*bb=11
	cc :=*bb
	//*bb =1000000
 //  cc :=*bb
   fmt.Printf("%V\n",cc)
	fmt.Printf("%F\n",cc)




	scoreMap := make(map[string]int)
	scoreMap["zhangsan"] =100
	scoreMap["李四"]= 50
	scoreMap["wuwang"] =30
	delete(scoreMap, "李四")

	for k,v := range scoreMap{
		fmt.Println(k,v)
	}
fmt.Println("----------------",scoreMap)
rand.Seed(time.Now().UnixNano()) //初始化随机数种子
var scoreMap = make(map[string]int,200)
}

func testArrayPoint(x *[]int) {
	fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
}











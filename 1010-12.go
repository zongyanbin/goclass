/**
golang的断言 Assertion 和 反射 reflect
断言
			把一节接口类型指定为它原始的类型
什么是反射
			通俗 指导数据的原始数据类型和数据内容，方法等，并且可以进行一定得操作
为什么要用反射
			通过接口或者其他的方式接收到了类型不固定的数据的时候需要写太多的swatch case 断言代码
			此时代码不灵活且通用性差
			反射这时候就可以无视类型
			改变元数据结构中的数据

反射的用法和主要函数

reflect包
	reflect.ValueOf() -> 获取输入参数接口中的数据的值
	reflect.TypeOf() ->	动态获取输入参数接口中的值得类型
	reflect.TypeOf().Kind()->来判断类型
	reflect.ValueOf().Fidle(int)->用来获取值
	reflect.FieldByIndex([]int{0,1})->层级取值
	reflect.ValueOf().Elem()->获取原始数据并操作

举例
	普通反射
	struct反射
	匿名或嵌入字段的反射
	判断传入的类型是否是我们想要的类型
	通过反射修改内容
	通过反射调用方法



*/
package main


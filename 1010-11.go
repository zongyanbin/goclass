/**
goroutine
channel

goroutine
		在调用一个方法的前面加上go 就是goroutine 它会让方法异步执行相当于协程
		示例
		func Run(){
		fmt.Print("123")
		}
		goRun()

		协程管理
		var wg sync.WaitGroup
		wg.Add()
		wg.Done()
		wg.Wait()

channel
	定义chan分为五部
	可读可取 c:=make(chan int)
	可读 var readChan <-chan int =c
	可取 var setChan chan<-int = c
	有缓冲 c:=make(chan int,5)
	无缓冲 c:=make(chan int)
channel 开启后以后是可以close的 当你觉得不再需要并且已经set完成的时候 你就需要去close它
此时需要注意
如果用到了range则必须在range之前就给它关闭

协程和线程
协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，这些用户级线程的调度也是自己实现的。
线程：一个线程上可以跑多个协程，协程是轻量级的线程

 */
package main

import (
	"fmt"
	"sync"
)

func main()  {
	//var wg sync.WaitGroup //声明: 协成管理器
	//wg.Add(1) //我有几个协程
	//go Run(&wg) //协程程序没关系，跑到一边运行去了
	//wg.Wait()

	c1 := make(chan int,10) //有缓冲先存满 chan 可以关闭

	//go func(){
	//	for i:=0; i<10; i++{
	//		c1<-i
	//	}
	//}()
	//
	//for i:=0;i<10; i++{
	//	fmt.Println(<-c1) //取得时候去差容量
	//}

	//c1 := make(chan int)	//无缓冲
	//可读可写
	var readc <-chan int = c1
	var writec chan <- int = c1
	close(c1)
	writec<-1
	fmt.Println(<-readc)

}

// 1 声明一个函数
func Run(wg *sync.WaitGroup){
	fmt.Println("我跑起来了")
	wg.Done()
}

// main包是一个象征性的入口，它不一定真实存在，这里叫做goroutine目录
package main

import (
	"fmt"
	"time"
)

//快速打印hello goroutine:0~hello goroutine:4
//快速其实就是说要我们开多个协程去打印

func hello(i int) {
	//这是内置的输出。用+连接字符串
	println("hello goroutine : " + fmt.Sprintln(i))
}

func HelloGoRoutine() {
	for i := 0; i < 5; i++ {
		//在同一个地址空间中，“go”语句作为独立的并发控制线程(goroutine)开始执行函数调用。
		//go 调用协程只需要在函数前面加上一个go关键字
		go func(j int) {
			hello(j)
		}(i)
	}
	//这里主要用了Sleep来做了阻塞，保证协程在执行完之前主线程不退出
	time.Sleep(time.Second)
}
func main() {
	HelloGoRoutine()
}

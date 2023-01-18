// Go也保留了通过共享内存实现通信的机制
// 但这种方法会存在多个Goroutine同时操作内存的情况

package main

import (
	"sync"
	"time"
)

var (
	x    int64
	lock sync.Mutex
	//Go语言提供了sync包和channel机制来解决并发机制中不同goroutine之间的同步和通信
	//sync.Mutex是Golang标准库提供的一个互斥锁，
	//当一个goroutine获得互斥锁权限后，其他请求锁的goroutine会阻塞在Lock()方法的调用上，直到调用Unlock()方法被释放。
)

// 加锁
func addWithLock() { //通过临界区控制通信的实现
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

// 不加锁
func addWithoutLock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

func Add() {
	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	println("WithoutLock:", x)
	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	println("WithLock:", x)
}

func main() {
	Add()
	//可以看到不加锁可能出现未知的结果
}

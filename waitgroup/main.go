//前面用Sleep()阻塞，但我们并不知道每个协程执行任务的时间
//因此就很难给出精确的Sleep()时间
//所以要用到更优雅的方法--WaiteGroup来实现并发任务的同步

package main

import (
	"fmt"
	"sync"
)

func ManyGoWait() {
	//创建一个sync包的WaitGroup的结构体变量
	var wg sync.WaitGroup
	//告诉计数器有五个子协程
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			//defer延迟，等这个子协程做完后再执行Done()
			defer wg.Done()
			fmt.Println("hello goroutine:", j)
		}(i) //这个括号里的参数是要传给线程的实参，那个j int就是对应接收的形参
	}
	//示例中使用 Add(5) 表示我们有 5个 子任务，然后起了 5个 协程去完成任务，
	//主协程使用 Wait() 方法等待 子协程执行完毕，输出一共等待的时间。
	wg.Wait()
}

func main() {
	ManyGoWait()
}

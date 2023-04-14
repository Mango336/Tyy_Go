package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var (
	x       int64
	wg      sync.WaitGroup // 声明全局等待组变量
	mutex   sync.Mutex
	rwMutex sync.RWMutex

	smp sync.Map
)

func GolangSync() {
	// funcSingleGoroutine()
	// funcMultipleGoroutine()
	// funcChannel()
	// funcSelect()

	// wg.Wait() // 阻塞--直到所有goroutine都Done

	// 读写互斥锁在读多写少的场景下能极大提高程序的性能
	// do(writeWithLock, readWithLock, 10, 10000)     // x: 10 cost: 10.876222265s
	// do(writeWithRWLock, readWithRWLock, 10, 10000) // x: 20 cost: 114.239961ms

	// funcSyncMap()
	// funcAtomic()
	syncExam()
}

func funcSingleGoroutine() {
	wg.Add(1) // 添加等待组中goroutine个数
	go func() {
		fmt.Println("Hello golang...")
		wg.Done() // 当前goroutine执行完毕 直接Done掉
	}()
	fmt.Println("Hello SingleGoroutine...")
}

func funcMultipleGoroutine() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go funcHello(i)
	}
}

func funcHello(i int) {
	defer wg.Done()
	fmt.Printf("Hello %d... \n", i)
}

func funcChannel() {
	var ch0 chan struct{}
	fmt.Println(ch0) // nil

	// 无缓冲区的通道
	ch1 := make(chan int) // 无缓冲区的channel
	// ch2 := make(chan int, 2)
	// 开启一个goroutine去接收ch1中的值
	// 无缓冲的channel需要先开启接收让goroutine出去接收阻塞中，如果后开启的话会导致死锁
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 10 // 发送 10发送到ch1中

	// 有缓冲区的通道
	ch2 := make(chan int, 1) // 有缓冲区的channel
	ch2 <- 2
	value, ok := <-ch2 // 因为有缓冲区 所以不会阻塞
	if ok {            // ch2关闭时 ok为false
		fmt.Println(value)
	} else {
		fmt.Println("ch2 is closed...")
	}

	ch3 := make(chan int, 2)
	ch3 <- 1
	ch3 <- 2
	close(ch3)
	// f0(ch3)

	// for range 接收值
	f1(ch3)

	// 单向通道
	ch4 := producer()
	sum := consumer(ch4)
	fmt.Println(sum)
}
func f0(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("channel is closed...")
			break

		}
		fmt.Printf("value: %v, ok: %v\n", v, ok)
	}
}
func f1(ch chan int) {
	for v := range ch {
		fmt.Printf("value: %v\n", v)
	}
}

// producer 返回一个channel
// 持续将符合条件的数发送到channel中（这里条件是奇数）
// 数据发送完成后将返回的通道关闭
// 返回的是 接收通道
func producer() <-chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

// consumer 用于从通道中接收数据并进行操作
// 参数为 接收通道（<- chan int）所以函数中只能接收数据
// 代码层面限制了consumer发送数据给channel
func consumer(ch <-chan int) (sum int) {
	for v := range ch {
		sum += v
	}
	return sum
}

func funcSelect() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

// 互斥锁
// 使用互斥锁的写操作
func writeWithLock() {
	mutex.Lock()
	x += 1
	time.Sleep(time.Millisecond * 10) // 假设写操作10ms
	mutex.Unlock()
	wg.Done()
}

// 使用互斥锁的读操作
func readWithLock() {
	mutex.Lock()
	time.Sleep(time.Millisecond) // 假设读操作为1ms
	mutex.Unlock()
	wg.Done()
}

// 读写锁
// 写操作
func writeWithRWLock() {
	rwMutex.Lock() // 加写锁
	x += 1
	time.Sleep(time.Millisecond * 10)
	rwMutex.Unlock() // 释放写锁
	wg.Done()
}

// 读操作
func readWithRWLock() {
	rwMutex.RLock()              // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作为1ms
	rwMutex.RUnlock()            // 释放读锁
	wg.Done()
}
func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf() // 写操作函数
	}
	// rc个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf() // 读操作函数
	}
	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x: %v cost: %v\n", x, cost)
}

// 并发安全的单例模式
type singleton struct{} // 单例类

var (
	instance *singleton
	once     sync.Once
)

// 获取单例
func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

// sync.Map包练习
func funcSyncMap() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(val int) {
			key := strconv.Itoa(val)
			smp.Store(key, val)        // sync.Map存储key-value
			value, ok := smp.Load(key) // 读取数据
			if ok {                    // ok表示是否有该key
				fmt.Printf("Sync.Map=>key:%v, value: %v\n", key, value)
			} else {
				fmt.Println("No this key in sync.Map...")
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 原子操作 atomic包练习
type Counter interface { // 接口
	Inc()
	Load() int64
}

// 普通版
type CommonCounter struct {
	counter int64
}

// 方法中接收者为指针时 会修改原来结构体中的对象
// 方法中接收者为值时 创建的是结构体副本，不会修改原来结构体中的对象
func (c *CommonCounter) Inc() {
	c.counter++
}
func (c CommonCounter) Load() int64 {
	return c.counter
}

// 互斥锁版
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (mc *MutexCounter) Inc() {
	mc.lock.Lock()
	mc.counter++
	mc.lock.Unlock()
}
func (mc *MutexCounter) Load() int64 {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	return mc.counter // 先return 再defer
}

// 原子操作版
type AtomicCounter struct {
	counter int64
}

func (ac *AtomicCounter) Inc() {
	atomic.AddInt64(&ac.counter, 1)
}
func (ac AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&ac.counter)
}

func funcAtomic() {
	ctr1 := CommonCounter{} // 非并发安全
	testCounter(&ctr1)      //Counter is 998, cost time is 535.57µs
	ctr2 := MutexCounter{}  // 使用互斥锁实现并发安全
	testCounter(&ctr2)      // Counter is 1000, cost time is 1.193061ms
	ctr3 := AtomicCounter{} // 并发安全 且比互斥锁效率更高
	testCounter(&ctr3)      // Counter is 1000, cost time is 1.161648ms
}

func testCounter(c Counter) {
	var wgTC sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wgTC.Add(1)
		go func() {
			c.Inc()
			wgTC.Done()
		}()
	}
	wgTC.Wait()
	end := time.Now()
	fmt.Printf("Counter is %v, cost time is %v\n", c.Load(), end.Sub(start))
}

var jobChan = make(chan int64, 10)
var resultChan = make(chan int64, 10)

func getNums(ch chan<- int64) {
	defer wg.Done()
	for {
		rand.Seed(int64(time.Now().Nanosecond()))
		x := rand.Int63()
		ch <- x
		time.Sleep(time.Millisecond * 1000)
	}
	// rand.Seed(int64(time.Now().Nanosecond()))
	// x := rand.Int63()
	// ch <- x
	// time.Sleep(time.Millisecond * 1000)
	// wg.Done()

}
func calculateNum(ch1 <-chan int64, ch2 chan<- int64) {
	defer wg.Done()
	for {
		job := <-ch1
		var res int64 = 0
		for job > 0 {
			res += job % 10
			job /= 10
		}
		resultChan <- res
	}
	// job := <-ch1
	// var res int64 = 0
	// for job > 0 {
	// 	res += job % 10
	// 	job /= 10
	// }
	// resultChan <- res
	// wg.Done()
}
func syncExam() {
	wg.Add(1)
	go getNums(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go calculateNum(jobChan, resultChan)
	}
	for res := range resultChan {
		fmt.Println(res)
	}
	wg.Wait()

}

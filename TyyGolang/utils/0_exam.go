// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// var (
// 	wg         sync.WaitGroup
// 	jobChan    = make(chan int64, 10)
// 	resultChan = make(chan int64, 10)
// )

// func getNums(ch chan<- int64) {
// 	defer wg.Done()
// 	for {
// 		rand.Seed(int64(time.Now().Nanosecond()))
// 		x := rand.Int63()
// 		ch <- x
// 		time.Sleep(time.Millisecond * 1000)
// 	}
// 	// rand.Seed(int64(time.Now().Nanosecond()))
// 	// x := rand.Int63()
// 	// ch <- x
// 	// time.Sleep(time.Millisecond * 1000)
// 	// wg.Done()

// }
// func calculateNum(ch1 <-chan int64, ch2 chan<- int64) {
// 	defer wg.Done()
// 	for {
// 		job := <-ch1
// 		var res int64 = 0
// 		for job > 0 {
// 			res += job % 10
// 			job /= 10
// 		}
// 		resultChan <- res
// 	}
// 	// job := <-ch1
// 	// var res int64 = 0
// 	// for job > 0 {
// 	// 	res += job % 10
// 	// 	job /= 10
// 	// }
// 	// resultChan <- res
// 	// wg.Done()
// }
// func main() {
// 	wg.Add(1)
// 	go getNums(jobChan)
// 	wg.Add(24)
// 	for i := 0; i < 24; i++ {
// 		wg.Add(1)
// 		go calculateNum(jobChan, resultChan)
// 	}
// 	for res := range resultChan {
// 		fmt.Println(res)
// 	}
// 	wg.Wait()

// }

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg         sync.WaitGroup
	jobChan    = make(chan int64, 10)
	resultChan = make(chan int64, 10)
)

func getNums(ch chan<- int64) {
	defer close(ch)
	defer wg.Done()
	for i := 0; i < 10; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		x := rand.Int63()
		ch <- x
		time.Sleep(time.Millisecond * 1000)
	}
}
func calculateNum(ch1 <-chan int64, ch2 chan<- int64) {
	defer wg.Done()
	job := <-ch1
	var res int64 = 0
	for job > 0 {
		res += job % 10
		job /= 10
	}
	resultChan <- res
}
func main() {
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

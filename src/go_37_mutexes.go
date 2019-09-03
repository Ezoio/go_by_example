package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var m = make(map[int]int)
	var mutexes = sync.Mutex{}
	var ops uint64 = 0

	for i := 0; i < 100; i++ {
		var total int = 0
		go func() {
			for {
				mutexes.Lock()
				total += m[rand.Intn(5)]
				mutexes.Unlock()
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	for j := 0; j < 10; j++ {
		go func() {
			for {
				key := rand.Intn(5)
				mutexes.Lock()
				m[key] = rand.Intn(100)
				mutexes.Unlock()
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	finalOps := atomic.LoadUint64(&ops)
	mutexes.Lock()
	fmt.Println("ops=", finalOps)
	mutexes.Unlock()
	fmt.Println(m)
}

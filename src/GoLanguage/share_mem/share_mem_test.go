package share_mem_test

import (
	"sync"
	"testing"
	"time"
)

func TestShareMem(t *testing.T) {
	counter := 1
	for i := 0; i < 5; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Millisecond * 40)
	t.Logf("counter:%d", counter)
}

func TestShareMemMutex(t *testing.T) {
	var mut sync.Mutex
	counter := 1
	for i := 0; i < 5; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Millisecond * 40)
	t.Logf("counter:%d", counter)
}

func TestMutexWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 1
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter:%d", counter)
}

package share_mem

import (
    "testing"
    "fmt"
    "time"
    "sync"
)

func TestCounter(t *testing.T) {
    counter := 0
    for i := 0; i < 5000; i++ {
        go func() {
            fmt.Println(counter)
            counter++
        }()
    }
    time.Sleep(1 * time.Second)
    t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
    counter := 0
    var mut sync.Mutex
    for i := 0; i < 5000; i++ {
        go func() {
            defer func() {
                mut.Unlock()
            }()
            mut.Lock()
            counter++;
        }()
    }
    time.Sleep(1 * time.Second)
    t.Logf("counter = %d", counter)
}

func TestCounterGroupWait(t *testing.T) {
    counter := 0
    var mut sync.Mutex
    var wg sync.WaitGroup
    for i := 0; i < 5000; i++ {
        wg.Add(1)
        go func() {
            defer func() {
                mut.Unlock()
            }()
            mut.Lock()
            counter++;
            wg.Done()
        }()
    }
    wg.Wait()
    time.Sleep(1 * time.Second)
    t.Logf("counter = %d", counter)

}

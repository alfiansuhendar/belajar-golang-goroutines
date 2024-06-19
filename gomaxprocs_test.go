package belajargolanggoroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("totalCpu", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("totalThread", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("totalGoRoutine", totalGoRoutine)

	group.Done()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("totalCpu", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("totalThread", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("totalGoRoutine", totalGoRoutine)

	group.Done()
}

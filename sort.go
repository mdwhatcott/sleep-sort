package sleepsort

import (
	"sync"
	"time"
)

type NonNegativeIntegers interface {
	uint | uint8 | uint16 | uint32 | uint64
}

func Sort[T NonNegativeIntegers](input ...T) (result []T) {
	waiter := new(sync.WaitGroup)
	waiter.Add(len(input))
	sorted := make(chan T)
	for _, i := range input {
		go func(i T) {
			time.Sleep(time.Millisecond * time.Duration(i))
			sorted <- i
			waiter.Done()
		}(i)
	}
	go func() { waiter.Wait(); close(sorted) }()
	for s := range sorted {
		result = append(result, s)
	}
	return result
}

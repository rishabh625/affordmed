package fibb

import (
	"sync"
)

var m *sync.Mutex

func init() {
	m = &sync.Mutex{}
}

func Fibbo(n int, memory *map[int]int) int {
	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else if val, ok := (*memory)[n]; ok {
		return val
	}
	val := Fibbo(n-1, memory) + Fibbo(n-2, memory)
	m.Lock()
	(*memory)[n] = val
	m.Unlock()
	return val
}

func Fibb(n int, c chan int, memory *map[int]int) (bool, int) {
	if val, ok := (*memory)[n]; ok {
		return true, val
	}
	c <- n
	return false, 0
}
func Processor(c chan int, memory *map[int]int) int {
	for {
		select {
		case op := <-c:
			n := Fibbo(op, memory)
			return n
		}
	}
}

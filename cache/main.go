package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f     Funcion
	cache map[int]FuncionResult
	lock  sync.Mutex
}

type Funcion func(key int) (interface{}, error)
type FuncionResult struct {
	value interface{}
	err   error
}

func NewCache(f Funcion) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FuncionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	result, exist := m.cache[key]
	m.lock.Unlock()
	if !exist {
		m.lock.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.lock.Unlock()
	}
	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38}
	var wg sync.WaitGroup
	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%d, %s, %d \n", index, time.Since(start), value)
		}(n)

	}
	wg.Wait()
}

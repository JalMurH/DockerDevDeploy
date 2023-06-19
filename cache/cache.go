package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int) int {
	fmt.Printf("Calculando fibonacci costosa para: %d\n", n)
	time.Sleep(5 * time.Second)
	return n
}

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Lock       sync.RWMutex
}

func (s *Service) Work(job int) {
	s.Lock.RLock()
	exist := s.InProgress[job]
	if exist {
		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()
		fmt.Printf("Esperando respuesta para el trabajo: %d", job)
		res := <-response
		fmt.Printf("La respuesta Lista, se recibio: %d\n", res)
	}
	s.Lock.RUnlock()
	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Se eta calculando la respuesta para: %d\n", job)
	resu := ExpensiveFibonacci(job)

	s.Lock.RLock()
	pendingworkers, exists := s.IsPending[job]
	s.Lock.RUnlock()

	if exists {
		for _, pendinpendingworker := range pendingworkers {
			pendinpendingworker <- resu
		}
		fmt.Printf("Resultado enviado para todos los workers %d \n", job)
	}
	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func main() {
	service := NewService()
	jobs := []int{3, 4, 5, 5, 8, 8, 8}
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for _, n := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}
	wg.Wait()
}

package main

import "sync"

func main() {
	sw := sync.WaitGroup{}
	sw.Add(1)
	go func() {
		Producer()
		sw.Done()
	}()
	//go Consumer()

	sw.Wait()
}

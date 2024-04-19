package main

// import "sync"

func main() {
	pipe := make(chan int, 1000)

	// wg := sync.WaitGroup{}
	// wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			setUp(pipe)
			// wg.Done()
		}()
	}

	// go func() {
	// 	wg.Wait()
	// 	close(pipe)
	// }()

	sum := 0
	for c := range pipe {
		sum += c
	}

	print(sum)
}

func setUp(pipe chan int) {
	for i := 1; i < 3000; i++ {
		pipe <- i
	}
}

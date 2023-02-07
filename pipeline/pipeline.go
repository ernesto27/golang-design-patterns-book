package pipeline

import "fmt"

func LaunchPipeline(amount int) int {
	return <-sum(power(generator(amount)))

}

func generator(max int) <-chan int {
	outChInt := make(chan int, 100)

	go func() {
		for i := 1; i <= max; i++ {
			fmt.Println("GENERATOR", i)
			outChInt <- i
		}
		close(outChInt)
	}()

	return outChInt
}

func power(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		var sum int
		for v := range in {
			fmt.Println("SUM ", v)
			sum += v
		}
		out <- sum
		close(out)
	}()

	return out
}

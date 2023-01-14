package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Devd")
	chev := make(chan int)
	chod := make(chan int)
	ctx, _ := context.WithCancel(context.Background())

	go print(ctx, chod, "odd")
	go print(ctx, chev, "even")
	func() {
		for i := 0; i < 15; {
			time.Sleep(1)
			i = i + 2
			chev <- i
		}
		close(chev)
	}()
	func() {
		for i := 1; i < 15; {
			time.Sleep(2)
			i = i + 2
			chod <- i
		}
		close(chod)
	}()
	fmt.Println("Done!!!!!")
}

func print(ctx context.Context, ch <-chan int, name string) {
	for v := range ch {
		fmt.Println(fmt.Sprintf("Printing...%d", v))
	}
	fmt.Println(fmt.Sprintf("Closing....%s", name))
}

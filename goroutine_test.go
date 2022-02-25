package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	for i := 0; i < 100; i++ {
		fmt.Println("Helloo GoWorld", i)
	}
}

func RunHelloGoroutine() {
	for i := 0; i < 100; i++ {
		fmt.Println("Helloo Goroutine", i)
	}
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	go RunHelloGoroutine()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(16 * time.Second)
}

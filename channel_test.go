package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- "Andri Yabu Ndapa Rehing"
		fmt.Println("Selesai mengirim data ke channel")
		time.Sleep(2 * time.Second)
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

package learn_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func PrintHellow() {
	fmt.Println("Hellow")
}

func TestGoroutine(t *testing.T) {
	go PrintHellow()
	fmt.Println("After")

	time.Sleep(1 * time.Second)
}

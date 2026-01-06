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
	go PrintHellow() // kalau ada return value, dia tidak bisa ditangkap
	fmt.Println("After")

	time.Sleep(1 * time.Second)
}

func DisplayFor(number int) {
	fmt.Println("Display number: ", number)
}

func TestBanyakGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayFor(i)
	}

	time.Sleep(10 * time.Second)
}

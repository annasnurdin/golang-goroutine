package learn_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestMembuatChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) //selalu tutup channel, pakai defer karena akan selalu dijalankan setelah fungsi selesai dijalankan

	// channel <- "Annas" // mengirim data ke channel
	// data := <-channel  // ambil data dari channel, masukkan ke var data

	// fmt.Println(<-channel) // mengirim channel ke parameter

	go func() {
		time.Sleep(3 * time.Second)
		channel <- "Annas" // (2) harus ada yang dikirim ke channel, karena ada yang menggunakannya
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel // (1) channel harus dipakai, kalau tidak, error
	fmt.Print(data)

	time.Sleep(2 * time.Second)
}

func Response(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Annas"
}

func TestChannelParam(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go Response(channel)
	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Annas N"
}

func OnlyOut(channel <-chan string) {
	time.Sleep(2 * time.Second)
	data := <-channel
	fmt.Println(data)
}

func TestInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

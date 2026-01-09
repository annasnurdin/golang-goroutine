package learn_golang_goroutine

import (
	"fmt"
	"strconv"
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

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	channel <- "Annas"
	channel <- "N"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("datanya: ", data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go Response(channel1)
	go Response(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2: ", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go Response(channel1)
	go Response(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2: ", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}

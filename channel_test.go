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

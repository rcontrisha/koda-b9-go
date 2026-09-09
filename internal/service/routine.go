package service

import (
	"fmt"
	"sync"
	"time"
)

func Rutinitas() {
	var wg sync.WaitGroup

	fmt.Println("Bangun Tidur.")
	wg.Go(mandi)
	wg.Go(ngopi)
	wg.Go(sarapan)
	wg.Go(beberes)
	wg.Wait()
	fmt.Println("Berangkat Kerja.")
}

func mandi() {
	fmt.Println("Mulai mandi.")
	time.Sleep(20 * time.Millisecond)
	fmt.Println("Selesai mandi.")
}

func ngopi() {
	fmt.Println("Bikin kopi.")
	time.Sleep(45 * time.Millisecond)
	fmt.Println("Selesai ngopi.")
}

func sarapan() {
	fmt.Println("Bikin sarapan")
	time.Sleep(15 * time.Millisecond)
	fmt.Println("Selesai sarapan.")
}

func beberes() {
	fmt.Println("Merapikan kasur.")
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Selesai merapikan kasur.")
}
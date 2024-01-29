// Saya memilih soal ini:
// https://tlx.toki.id/problems/osn-2008/2D

package main

import "fmt"

func main() {
	// Meminta input dari pengguna
	fmt.Print("Masukkan nilai n (1 ≤ n ≤ 50.000): ")
	var n int
	fmt.Scanln(&n)

	// Inisialisasi barisan
	barisan := make([]int, n)
	barisan[0] = 1

	// Menghitung suku-suku berikutnya
	for i := 1; i < n; i++ {
		barisan[i] = barisan[i-1] + i + 1
	}

	// Menampilkan hasil keluaran
	for i := 0; i < n; i++ {
		fmt.Println(barisan[i])
	}
}

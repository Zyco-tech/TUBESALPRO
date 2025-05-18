package main

import (
	"fmt"
	"time"
)

type dataLangganan struct {
	nama       string
	metode     string
	biaya      float64
	tanggal    int
	status     string
	jatuhTempo int
}

const NMAX int = 10

type TabLangganan [NMAX]dataLangganan

func main() {
	var A TabLangganan
	var jumlah, pilih int

	for {
		menu()
		fmt.Println("Pilih (1/2/3/4/5/6/7/8) ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			MenambahLangganan(&A, &jumlah)
		case 2:
			ubahData(&A, jumlah)
		case 3:
			hapusData(&A, &jumlah)
		case 4:
			TampilkanSemuaLangganan(A, jumlah)
		case 5:
			menuSearch()
			fmt.Println("Pilih (1/2/3/4/5/6) ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				seqSearch(A, jumlah, pilih)
			case 2:
				seqSearch(A, jumlah, pilih)
			case 3:
				// biaya binary
			case 4:
				binSearch(A, jumlah, pilih)
			case 5:
				// status seq
			case 6:
				break
			}

		case 6:
			menuSort()
			fmt.Println("Pilih (1/2/3) ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				SelectionSortBiaya(&A, jumlah)
			case 2:
				InsertionSortJatuhTempo(&A, jumlah)
			case 3:
				break
			}
		case 7:
			fmt.Printf("Total Pengeluaran Perbulan Adalah Rp.%.2f\n", hitungPengeluaran(A, jumlah))
		case 8:
			return
		}
	}
}
func menu() {
	fmt.Println("==============MENU=============")
	fmt.Println("1. Tambah Data Langganan")
	fmt.Println("2. Ubah Data Langganan")
	fmt.Println("3. Hapus Data Langganan")
	fmt.Println("4.	Tampilkan Semua Langganan")
	fmt.Println("5.	Search Data")
	fmt.Println("6.	Sort Data")
	fmt.Println("7.	Hitung Pengeluaran Perbulan") //mungkin ganti lagi kalimatnya
	fmt.Println("8.	Exit")
	fmt.Println("===============================")
}
func menuSearch() {
	fmt.Println("Search Berdasarkan Apa?")
	fmt.Println("1.	Nama")
	fmt.Println("2.	Metode Pembayaran")
	fmt.Println("3.	Biaya")
	fmt.Println("4.	Tanggal Pembayaran")
	fmt.Println("5.	Status")
	fmt.Println("6.	Kembali")
}
func menuSort() {
	fmt.Println("Sort Berdasarkan Apa?")
	fmt.Println("1.	Biaya")
	fmt.Println("2.	Tanggal Jatuh Tempo")
	fmt.Println("3.	Kembali")
}
func MenambahLangganan(A *TabLangganan, jumlah *int) {
	if *jumlah >= NMAX {
		fmt.Println("Data langganan penuh!")
		return
	}

	fmt.Println("=== Tambah Langganan ===")
	fmt.Print("Nama Layanan: ")
	fmt.Scan(&A[*jumlah].nama)
	fmt.Print("Metode Pembayaran: ")
	fmt.Scan(&A[*jumlah].metode)
	fmt.Print("Biaya Bulanan: ")
	fmt.Scan(&A[*jumlah].biaya)

	fmt.Print("Tanggal Pembayaran (1-31): ")
	fmt.Scan(&A[*jumlah].tanggal)

	fmt.Print("Status (aktif/nonaktif): ")
	fmt.Scan(&A[*jumlah].status)
	*jumlah++
	fmt.Println("Langganan berhasil ditambahkan.")
}
func ubahData(A *TabLangganan, jumlah int) {
	var j int
	//prosedur print tabel
	TampilkanSemuaLangganan(*A, jumlah)
	fmt.Println("Pilih nomor yang akan diubah (0 untuk kembali)")
	fmt.Scan(&j)
	if j == 0 {
		return
	}
	fmt.Println("=== Ubah Langganan ===")
	fmt.Print("Nama Layanan: ")
	fmt.Scan(&A[j-1].nama)
	fmt.Print("Metode Pembayaran: ")
	fmt.Scan(&A[j-1].metode)
	fmt.Print("Biaya Bulanan: ")
	fmt.Scan(&A[j-1].biaya)

	fmt.Print("Tanggal Pembayaran (1-31): ")
	fmt.Scan(&A[j-1].tanggal)

	fmt.Print("Status (aktif/nonaktif): ")
	fmt.Scan(&A[j-1].status)
	fmt.Println("Langganan berhasil diubah.")
}
func hapusData(A *TabLangganan, n *int) {
	var j, i int
	TampilkanSemuaLangganan(*A, *n)
	fmt.Println("pilih nomor yg akan dihapus (0 untuk kembali)")
	fmt.Scan(&j)
	if j == 0 {
		return
	}
	for i = j - 1; i < *n; i++ {
		A[i] = A[i+1]
	}
	*n--
}
func seqSearch(A TabLangganan, n, pilihan int) {
	var i int
	var ketemu bool
	var dicari string

	fmt.Println("Masukkan yang akan dicari")
	fmt.Scan(&dicari)
	ketemu = false
	i = 0
	if pilihan == 1 {
		for i < n && !ketemu {
			ketemu = A[i].nama == dicari
			i++
		}
	} else if pilihan == 2 {
		for i < n && !ketemu {
			ketemu = A[i].metode == dicari
			i++
		}
	}
	if ketemu {
		TampilkanLangganan(A, i-1)
	} else {
		fmt.Println("Langganan tidak ditemukan")
	}
}
func binSearch(A TabLangganan, n, pilihan int) {
	var left, right, mid int
	var idx int
	var dicari int
	fmt.Println("Masukkan yang akan dicari")
	fmt.Scan(&dicari)

	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if dicari < A[mid].tanggal {
			right = mid - 1
		} else if dicari > A[mid].tanggal {
			left = right + 1
		} else {
			idx = mid
		}
		mid = (left + right) / 2
	}
	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	} else {
		TampilkanLangganan(A, idx)
	}
}

func TampilkanLangganan(A TabLangganan, i int) {

	fmt.Printf("%d. Nama: %s\n", i+1, A[i].nama)
	fmt.Printf("   Metode Pembayaran: %s\n", A[i].metode)
	fmt.Printf("   Biaya Bulanan: %.2f\n", A[i].biaya)
	fmt.Printf("   Tanggal Pembayaran: %d\n", A[i].tanggal)
	hitungTempo(&A, i)
	if A[i].jatuhTempo < 0 {
		fmt.Printf("   Jatuh tempo telah lewat %d hari\n", (time.Now().Day() - A[i].tanggal))
	} else {
		fmt.Printf("   %d hari sebelum jatuh tempo\n", A[i].jatuhTempo)
	}
	fmt.Printf("   Status: %s\n", A[i].status)
	fmt.Println("---------------------------")
}
func SelectionSortBiaya(A *TabLangganan, N int) {
	var i, idx, pass int
	var temp dataLangganan
	pass = 1

	for pass < N {
		idx = pass
		i = pass - 1

		for i < N {
			if A[i].biaya > A[idx].biaya {
				idx = i
			}
			i++
		}

		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++

	}
}
func InsertionSortJatuhTempo(A *TabLangganan, N int) {
	var pass, i int
	var temp dataLangganan

	pass = 1
	for pass <= N-1 {
		i = pass
		temp = A[pass]

		for i > 0 && temp.tanggal > A[i-1].tanggal {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}
func hitungPengeluaran(A TabLangganan, jumlah int) float64 {
	var total float64
	var i int
	for i = 0; i < jumlah; i++ {
		total += A[i].biaya
	}
	return total
}
func TampilkanSemuaLangganan(A TabLangganan, jumlah int) {
	var i int

	if jumlah == 0 {
		fmt.Println("Belum ada data langganan.")
		return
	}

	fmt.Println("=== Daftar Langganan ===")
	for i = 0; i < jumlah; i++ {
		TampilkanLangganan(A, i)
	}
	fmt.Printf("Total langganan: %d\n", jumlah)
}
func hitungTempo(A *TabLangganan, i int) {
	A[i].jatuhTempo = A[i].tanggal - time.Now().Day()
}

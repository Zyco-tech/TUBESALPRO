package main

import (
	"fmt"
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
			ubahData(&A)
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
				//nama seqSearch(A, jumlah, pilih)
			case 2:
				// metode seq Search
			case 3:
				// biaya binary
			case 4:
				// tgl binary
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
				//insertion sort blm jadi
			case 3:
				break
			}
		case 7:
			//func hitung pengeluaran per bulan blm
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
func TampilkanLangganan(A TabLangganan, i int){

	fmt.Printf("%d. Nama: %s\n", i+1, A[i].nama)
	fmt.Printf("   Metode Pembayaran: %s\n", A[i].metode)
	fmt.Printf("   Biaya Bulanan: %.2f\n", A[i].biaya)
	fmt.Printf("   Tanggal Pembayaran: %d\n", A[i].tanggal)
	hitungTempo(&A,i)
	if A[i].jatuhTempo < 0{
		fmt.Printf("   Jatuh tempo telah lewat %d hari\n",(time.Now().Day()-A[i].tanggal))
	} else {
		fmt.Printf("   %d hari sebelum jatuh tempo\n",A[i].jatuhTempo)
	}
	fmt.Printf("   Status: %s\n", A[i].status)
	fmt.Println("-------------------------")
}
func SelectionSortBiaya(A *TabLangganan, N int) {
    var i, idx, pass int
    var temp dataLangganan
    pass = 1

    for pass < N {
        idx = pass 
        i = pass-1

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

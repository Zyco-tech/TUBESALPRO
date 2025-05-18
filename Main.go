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
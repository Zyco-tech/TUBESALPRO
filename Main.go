package main

import (
	"fmt"
	"time"
)

type dataLangganan struct {
	id         string
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
		fmt.Print("Pilih (1/2/3/4/5/6/7/8) : ")
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
			search(A, jumlah)
		case 6:
			sort(A, jumlah)
		case 7:
			hitungPengeluaran(A, jumlah)
		}
		if pilih == 8 {
			break
		}
	}
}
func menu() {
	fmt.Printf("%-30s\n", "============= MENU =============")
	fmt.Printf("%-30s\n", "1. Tambah data langganan")
	fmt.Printf("%-30s\n", "2. Ubah data langganan")
	fmt.Printf("%-30s\n", "3. Hapus data langganan")
	fmt.Printf("%-30s\n", "4. Tampilkan tabel data")
	fmt.Printf("%-30s\n", "5. Search Data")
	fmt.Printf("%-30s\n", "6. Sort Data")
	fmt.Printf("%-30s\n", "7. Hitung Pengeluaran Perbulan")
	fmt.Printf("%-30s\n", "8. Exit")
	fmt.Printf("%-30s\n", "================================")
}
func menuSearch() {
	fmt.Printf("%-30s\n", "Search Berdasarkan Apa?")
	fmt.Printf("%-30s\n", "1. ID")
	fmt.Printf("%-30s\n", "2. Nama")
	fmt.Printf("%-30s\n", "3. Metode Pembayaran")
	fmt.Printf("%-30s\n", "4. Biaya")
	fmt.Printf("%-30s\n", "5. Tanggal Pembayaran")
	fmt.Printf("%-30s\n", "6. Status")
	fmt.Printf("%-30s\n", "7. Nilai Terbesar")
	fmt.Printf("%-30s\n", "8. Nilai Terkecil")
	fmt.Printf("%-30s\n", "9. Kembali")
}
func menuSort() {
	fmt.Printf("%-30s\n", "Urutkan Berdasarkan?")
	fmt.Printf("%-30s\n", "1. ID")
	fmt.Printf("%-30s\n", "2. Nama")
	fmt.Printf("%-30s\n", "3. Metode Pembayaran ")
	fmt.Printf("%-30s\n", "4. Biaya")
	fmt.Printf("%-30s\n", "5. Tanggal")
	fmt.Printf("%-30s\n", "6. Kembali")
}
func menuFindMaxMin() {
	fmt.Printf("%-30s\n", "Search Berdasarkan Apa?")
	fmt.Printf("%-30s\n", "1. Biaya")
	fmt.Printf("%-30s\n", "2. Tanggal")
	fmt.Printf("%-30s\n", "3. Kembali")
}
func MenambahLangganan(A *TabLangganan, jumlah *int) {
	var pilih int
	var i int

	if *jumlah >= NMAX {
		fmt.Println("Data langganan penuh!")
		return
	}

	fmt.Println("======= Tambah Langganan =======")
	// Revisi penambahan ID di langganan
	fmt.Print("ID Layanan (Contoh 001): ")
	fmt.Scan(&A[*jumlah].id)
	for i = 0; i < *jumlah; i++ {
		if A[i].id == A[*jumlah].id {
			A[*jumlah] = A[*jumlah+1]
			fmt.Println("ID ini sudah ada!")
			return
		}
	}

	fmt.Print("Nama Layanan: ")
	fmt.Scan(&A[*jumlah].nama)
	// Revisi pengecekan jika nama sama yg dimasukkan
	for i = 0; i < *jumlah; i++ {
		if A[i].nama == A[*jumlah].nama {
			A[*jumlah] = A[*jumlah+1]
			fmt.Println("Langganan ini sudah ada!")
			return
		}
	}

	fmt.Print("Metode Pembayaran: ")
	fmt.Scan(&A[*jumlah].metode)
	fmt.Print("Biaya Bulanan: ")
	fmt.Scan(&A[*jumlah].biaya)

	fmt.Print("Tanggal Pembayaran (1-31): ")
	fmt.Scan(&A[*jumlah].tanggal)
	if A[*jumlah].tanggal < 1 || A[*jumlah].tanggal > 31 {
		fmt.Println("Tanggal Tidak Valid!")
		return
	}

	fmt.Print("Status aktif/nonaktif (1/2): ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		A[*jumlah].status = "Aktif"
	} else if pilih == 2 {
		A[*jumlah].status = "Nonaktif"
	} else {
		fmt.Println("Pilihan Tidak Valid!")
		return
	}

	*jumlah += 1
	fmt.Println("Langganan berhasil ditambahkan.")
	fmt.Println("================================")
}
func ubahData(A *TabLangganan, jumlah int) {
	var j, i, pilih int
	var ada bool = false
	var tempID, tempNama, idAwal string

	if jumlah <= 0 {
		fmt.Println("Data Langganan Kosong!")
		return
	}

	TampilkanSemuaLangganan(*A, jumlah)
	fmt.Print("Pilih nomor yang akan diubah (0 untuk kembali): ")
	fmt.Scan(&j)
	if j == 0 {
		return
	}
	for i < jumlah && !ada {
		ada = i == j-1
		i++
	}
	if ada {
		fmt.Println("====== Ubah Data Langganan ======")
		// Revisi pengecekan jika ID sama yg dimasukkan
		idAwal = A[j-1].id
		fmt.Print("ID Layanan (Contoh 001): ")
		fmt.Scan(&tempID)
		for i = 0; i < jumlah; i++ {
			if A[i].id == tempID {
				fmt.Println("ID ini sudah ada!")
				return
			}
		}
		A[j-1].id = tempID

		fmt.Print("Nama Layanan: ")
		fmt.Scan(&tempNama)
		// Revisi pengecekan jika nama sama yg dimasukkan
		for i = 0; i < jumlah; i++ {
			if A[i].nama == tempNama {
				A[j-1].id = idAwal
				fmt.Println("Langganan ini sudah ada!")
				return
			}
		}
		A[j-1].nama = tempNama

		fmt.Print("Metode Pembayaran: ")
		fmt.Scan(&A[j-1].metode)

		fmt.Print("Biaya Bulanan: ")
		fmt.Scan(&A[j-1].biaya)

		fmt.Print("Tanggal Pembayaran (1-31): ")
		fmt.Scan(&A[j-1].tanggal)
		if A[j-1].tanggal < 1 && A[j-1].tanggal > 31 {
			fmt.Println("Tanggal Tidak Valid")
			return
		}

		fmt.Print("Status aktif/nonaktif (1/2): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			A[j-1].status = "Aktif"
		} else if pilih == 2 {
			A[j-1].status = "Nonaktif"
		} else {
			fmt.Println("Status Tidak Valid")
			return
		}

		fmt.Println("Langganan berhasil diubah.")
		fmt.Println("================================")
	} else {
		fmt.Println("Data tidak ditemukan!")
		return
	}
}
func hapusData(A *TabLangganan, n *int) {
	var j, i int
	var ada bool = false

	if *n <= 0 {
		fmt.Println("Data Langganan Kosong!")
		return
	}
	TampilkanSemuaLangganan(*A, *n)
	fmt.Print("Pilih nomor yang akan dihapus (0 untuk kembali): ")
	fmt.Scan(&j)
	if j == 0 {
		return
	}
	for i < *n && !ada {
		ada = i == j-1
		i++
	}

	if ada {
		for i = j - 1; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n -= 1
	} else {
		fmt.Println("Data tidak ditemukan!")
		return
	}
}
func search(A TabLangganan, jumlah int) {
	var pilih int
	var iMax, iMin int

	if jumlah <= 0 {
		fmt.Println("Data Langganan Kosong!")
		return
	}

	menuSearch()
	fmt.Print("Pilih (1/2/3/4/5/6/7) : ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		seqSearch(A, jumlah, pilih) // id
	case 2:
		seqSearch(A, jumlah, pilih) // nama
	case 3:
		seqSearch(A, jumlah, pilih) // metode
	case 4:
		binSearch(A, jumlah, pilih) // biaya
	case 5:
		binSearch(A, jumlah, pilih) // tgl
	case 6:
		seqSearch(A, jumlah, pilih) // status
	case 7:
		menuFindMaxMin()
		fmt.Print("Pilih (1/2/3) : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			findMax(A, jumlah, pilih, &iMax)
			fmt.Println("================================")
			fmt.Println("Biaya terbesar ada di:")
			TampilkanLangganan(A, iMax)
		case 2:
			findMax(A, jumlah, pilih, &iMax)
			fmt.Println("================================")
			fmt.Println("Tanggal terbesar ada di:")
			TampilkanLangganan(A, iMax)
		case 3:
			return
		}
	case 8:
		menuFindMaxMin()
		fmt.Print("Pilih (1/2/3) : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			findMin(A, jumlah, pilih, &iMin)
			fmt.Println("================================")
			fmt.Println("Biaya terkecil ada di:")
			TampilkanLangganan(A, iMin)
		case 2:
			findMin(A, jumlah, pilih, &iMin)
			fmt.Println("================================")
			fmt.Println("Tanggal terkecil ada di:")
			TampilkanLangganan(A, iMin)
		case 3:
			return
		}
	case 9:
		return
	}
}
func seqSearch(A TabLangganan, n, pilihan int) {
	var i, jumAktif, jumNonaktif int
	var ketemu bool
	var dicari string

	i = 0
	if pilihan == 1 || pilihan == 2 || pilihan == 3 {
		fmt.Print("Masukkan yang akan dicari: ")
		fmt.Scan(&dicari)
		ketemu = false
		if pilihan == 1 {
			for i < n && !ketemu {
				ketemu = A[i].id == dicari
				i++
			}
		} else if pilihan == 2 {
			for i < n && !ketemu {
				ketemu = A[i].nama == dicari
				i++
			}
		} else if pilihan == 3 {
			for i < n {
				if i == 0 {
					fmt.Printf("Metode pembayaran %s berada pada data:\n", dicari)
				}
				if A[i].metode == dicari {
					TampilkanLangganan(A, i)
				}
				i++
			}
		}
	}

	if pilihan == 6 {
		fmt.Println("Search pilihan aktif/nonaktif (1/2)")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			for i = 0; i < n; i++ {
				if A[i].status == "Aktif" {
					jumAktif++
				}
			}
			if jumAktif != 0 {
				fmt.Println("Langganan aktif berada pada data:")
				for i = 0; i < n; i++ {
					if A[i].status == "Aktif" {
						TampilkanLangganan(A, i)
					}
				}
				return
			} else {
				fmt.Println("Tidak ada langganan aktif")
				return
			}
		} else if pilihan == 2 {
			for i = 0; i < n; i++ {
				if A[i].status == "Nonaktif" {
					jumNonaktif++
				}
			}
			if jumNonaktif != 0 {
				fmt.Println("Langganan nonaktif berada pada data:")
				for i = 0; i < n; i++ {
					if A[i].status == "Nonaktif" {
						TampilkanLangganan(A, i)
					}
				}
				return
			} else {
				fmt.Println("Tidak ada langganan nonaktif")
				return
			}
		} else {
			fmt.Println("Pilihan Tidak Valid")
			return
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
	var tanggal int
	var biaya float64
	var pass, i int
	var temp dataLangganan

	if pilihan == 4 {
		fmt.Print("Masukkan biaya yang akan dicari: ")
		fmt.Scan(&biaya)

		pass = 1
		for pass <= n-1 {
			i = pass
			temp = A[pass]
			for i > 0 && temp.biaya < A[i-1].biaya {
				A[i] = A[i-1]
				i--
			}
			A[i] = temp
			pass++
		}

		left = 0
		right = n - 1
		idx = -1
		mid = (left + right) / 2
		for left <= right && idx == -1 {

			if biaya < A[mid].biaya {
				right = mid - 1
			} else if biaya > A[mid].biaya {
				left = mid + 1
			} else {
				idx = mid
			}
			mid = (left + right) / 2
		}
	} else if pilihan == 5 {
		fmt.Print("Masukkan tanggal yang akan dicari: ")
		fmt.Scan(&tanggal)

		pass = 1
		for pass <= n-1 {
			i = pass
			temp = A[pass]
			for i > 0 && temp.tanggal < A[i-1].tanggal {
				A[i] = A[i-1]
				i--
			}
			A[i] = temp
			pass++
		}

		left = 0
		right = n - 1
		idx = -1
		mid = (left + right) / 2
		for left <= right && idx == -1 {

			if tanggal < A[mid].tanggal {
				right = mid - 1
			} else if tanggal > A[mid].tanggal {
				left = mid + 1
			} else {
				idx = mid
			}
			mid = (left + right) / 2
		}
	}

	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	} else {
		fmt.Printf("Data yang dicari berada di posisi ke-%d setelah diurutkan ascending\n", idx+1)
		TampilkanLangganan(A, idx)
	}
}

func findMin(A TabLangganan, jumlah int, pilihan int, min *int) {
	if jumlah == 0 {
		*min = 0
		return
	}
	*min = 0
	if pilihan == 1 { //biaya
		for i := 1; i < jumlah; i++ {
			if A[i].biaya < A[*min].biaya {
				*min = i
			}
		}
	} else if pilihan == 2 { //tanggal
		for i := 1; i < jumlah; i++ {
			if A[i].tanggal < A[*min].tanggal {
				*min = i
			}
		}
	}
}
func findMax(A TabLangganan, jumlah int, pilihan int, max *int) {
	if jumlah == 0 {
		*max = 0
		return
	}
	*max = 0
	if pilihan == 1 {
		for i := 1; i < jumlah; i++ {
			if A[i].biaya > A[*max].biaya && A[i].status == "Aktif" {
				*max = i
			}
		}
	} else if pilihan == 2 {
		for i := 1; i < jumlah; i++ {
			if A[i].tanggal > A[*max].tanggal && A[i].status == "Aktif" {
				*max = i
			}
		}
	}
}

func TampilkanLangganan(A TabLangganan, i int) {
	fmt.Printf("%d. ID: %s\n", i+1, A[i].id)
	fmt.Printf("   Nama: %s\n", A[i].nama)
	fmt.Printf("   Metode Pembayaran: %s\n", A[i].metode)
	fmt.Printf("   Biaya Bulanan: %.2f\n", A[i].biaya)
	fmt.Printf("   Tanggal Pembayaran: %d\n", A[i].tanggal)
	if A[i].status == "Aktif" {
		hitungTempo(&A, i)
		if A[i].jatuhTempo < 0 {
			fmt.Printf("   Jatuh tempo telah lewat : %d hari\n", (time.Now().Day() - A[i].tanggal))
		} else if A[i].jatuhTempo == 0 {
			fmt.Println("   Jatuh tempo hari ini")
		} else {
			fmt.Printf("   %d hari sebelum jatuh tempo\n", A[i].jatuhTempo)
		}
	}
	fmt.Printf("   Status: %s\n", A[i].status)
	fmt.Println("================================")
}
func sort(A TabLangganan, jumlah int) {
	var pilih int
	var cara string

	if jumlah <= 0 {
		fmt.Println("Data Langganan Kosong!")
		return
	}

	fmt.Println("Urutkan Secara Ascending atau Descending (1/2) : ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		cara = "Ascending"
		menuSort()
		fmt.Print("Pilih (1/2/3/4/5/6) : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			InsertionSort(A, jumlah, pilih, cara) //id
		case 2:
			InsertionSort(A, jumlah, pilih, cara) // nama
		case 3:
			InsertionSort(A, jumlah, pilih, cara) // metode
		case 4:
			SelectionSortBiaya(A, jumlah, cara) // biaya
		case 5:
			InsertionSort(A, jumlah, pilih, cara) // tgl
		}
	} else if pilih == 2 {
		cara = "Descending"
		menuSort()
		fmt.Print("Pilih (1/2/3/4/5/6) : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			InsertionSort(A, jumlah, pilih, cara) //id
		case 2:
			InsertionSort(A, jumlah, pilih, cara) // nama
		case 3:
			InsertionSort(A, jumlah, pilih, cara) // metode
		case 4:
			SelectionSortBiaya(A, jumlah, cara) // biaya
		case 5:
			InsertionSort(A, jumlah, pilih, cara) // tgl
		}
	} else {
		fmt.Println("Pilihan Tidak Valid!")
		return
	}
}
func SelectionSortBiaya(A TabLangganan, N int, cara string) {
	var i, idx, pass int
	var temp dataLangganan
	pass = 1

	if cara == "Ascending" {
		for pass < N {
			idx = pass
			i = pass - 1

			for i < N {
				if A[i].biaya < A[idx].biaya {
					idx = i
				}
				i++
			}

			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		}
	} else if cara == "Descending" {
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
	TampilkanSemuaLangganan(A, N)
}
func InsertionSort(A TabLangganan, N, pilihan int, cara string) {
	var pass, i int
	var temp dataLangganan

	pass = 1
	if cara == "Ascending" {
		switch pilihan {
		case 1:
			for pass <= N-1 {
				i = pass
				temp = A[pass]
				for i > 0 && temp.id < A[i-1].id {
					A[i] = A[i-1]
					i--
				}
				A[i] = temp
				pass++
			}
		case 2:
			for pass <= N-1 {
				i = pass
				temp = A[pass]
				for i > 0 && temp.nama < A[i-1].nama {
					A[i] = A[i-1]
					i--
				}
				A[i] = temp
				pass++
			}
		case 3:
			for pass <= N-1 {
				i = pass
				temp = A[pass]
				for i > 0 && temp.metode < A[i-1].metode {
					A[i] = A[i-1]
					i--
				}
				A[i] = temp
				pass++
			}
		case 5:
			for pass <= N-1 {
				i = pass
				temp = A[pass]
				for i > 0 && temp.tanggal < A[i-1].tanggal {
					A[i] = A[i-1]
					i--
				}
				A[i] = temp
				pass++
			}
		}
	} else if cara == "Descending" {
		switch pilihan {
		case 1:
			for pass <= N-1 {
				i = pass
				temp = A[pass]
				for i > 0 && temp.id > A[i-1].id {
					A[i] = A[i-1]
					i--
				}
				A[i] = temp
				pass++
			}
		case 2:
			for pass <= N-1 {
				i = pass
				temp = A[pass]
				for i > 0 && temp.nama > A[i-1].nama {
					A[i] = A[i-1]
					i--
				}
				A[i] = temp
				pass++
			}
		case 3:
			for pass <= N-1 {
				i = pass
				temp = A[pass]
				for i > 0 && temp.metode > A[i-1].metode {
					A[i] = A[i-1]
					i--
				}
				A[i] = temp
				pass++
			}
		case 5:
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
	}
	TampilkanSemuaLangganan(A, N)
}
func hitungPengeluaran(A TabLangganan, jumlah int) {
	var total float64
	var i, pilih, max int

	if jumlah <= 0 {
		fmt.Println("Data Langganan Kosong!")
		return
	}

	for i = 0; i < jumlah; i++ {
		if A[i].status == "Aktif" {
			total += A[i].biaya
		}
	}

	fmt.Printf("Total Pengeluaran Perbulan Adalah Rp.%.2f\n", total)
	pilih = 1
	findMax(A, jumlah, pilih, &max)

	fmt.Printf("Data ke-%d disarankan untuk dihapus agar menghemat biaya\n", max+1)
	TampilkanLangganan(A, max)
}
func TampilkanSemuaLangganan(A TabLangganan, jumlah int) {
	var i int

	if jumlah == 0 {
		fmt.Println("Belum ada data langganan.")
		return
	}

	fmt.Println("======= Daftar Langganan =======")
	for i = 0; i < jumlah; i++ {
		TampilkanLangganan(A, i)
	}
	fmt.Printf("Total langganan: %d\n", jumlah)
}
func hitungTempo(A *TabLangganan, i int) {
	A[i].jatuhTempo = A[i].tanggal - time.Now().Day()
}

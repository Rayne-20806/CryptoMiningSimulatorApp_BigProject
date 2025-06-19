package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// digunakan untuk input dan variabel global
var scanner = bufio.NewScanner(os.Stdin)

// func membaca string, interger, float
func bacaString(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
func bacaInt(prompt string, min, max int) int {
	for {
		inputString := bacaString(prompt)
		val, err := strconv.Atoi(inputString)
		if err != nil {
			fmt.Println("⚠️Input tidak valid")
			continue
		}
		if val < min || val > max {
			fmt.Printf(" ⚠️Input harus diantara %d dan %d\n", min, max)
			continue
		}
		return val
	}

}
func bacaFloat(prompt string, min float64) float64 {
	for {
		inputString := bacaString(prompt)
		val, err := strconv.ParseFloat(inputString, 64)
		if err != nil {
			fmt.Println("⚠️Input tidak valid, masukkan angka desimal (3.14 contohnya).")
			continue
		}
		if val < min {
			fmt.Printf(" ⚠️Input harus diantara lebih besar atau == %.2f\n", min)
			continue
		}
		return val
	}
}

// fungsi tekan enter untuk kembali ke menu
func tekanEnterUntukLanjut() {
	fmt.Print("\nTekan enter untuk melanjutkan...")
	scanner.Scan()
}

// struct aset kripto yang digunakan untuk menyimpan data aset cripto
type asetKrypto struct {
	nama             string
	tingkatKesulitan int
	estimasiReward   float64
	algoritma        string
}

// Struct untuk menampilkan informasi aset kripto dasar sistem
type infoKriptoDasarSistem struct {
	namaLengkap string
	namaAlias   []string
	algoritma   string
}

type riwayatMining struct {
	namaAset              string
	jumlahReward          float64
	waktuSelesai          time.Time
	durasiSimulasi        time.Duration
	dayaKomputasiPengguna float64
	blokDitemukan         float64
	periodeSimulasi       time.Time
}

// variabel global  yang digunakan menyimpan data aset kripto dan riwayat mining
var daftarAsetKripto [1000]asetKrypto
var jumlahAset int = 0
var daftarRiwayatMining [1000]riwayatMining
var jumlahRiwayatMining int = 0

var daftarInfoKriptoSistem = []infoKriptoDasarSistem{
	{"BITCOIN", []string{"BTC", "BITCOIN"}, "SHA-256"},
	{"ETHEREUM", []string{"ETH", "ETHERIUM"}, "ETHASH"},
	{"LITECOIN", []string{"LTC", "LITECOIN"}, "SCRYPT"},
	{"DOGECOIN", []string{"DOGE", "DG"}, "SCRYPT"},
	{"MONERO", []string{"XMR", "MONERO"}, "RANDOMX"},
	{"CARDANO", []string{"ADA", "CARDANO"}, "OUROBOROS"},
	{"SOLANA", []string{"SOL", "SOLANA"}, "PROOF OF HISTORY"},
	{"BITCOIN CASH", []string{"BCH", "BITCOIN CASH"}, "SHA-256"},
	{"SIBA INU", []string{"SIB", "SIBA INU"}, "ERC-20"},
}

// fungsi untuk algoritma crypto
func algoritmaCrypto(namaAset string) string {
	namaAsetUpper := strings.ToUpper(namaAset)
	switch namaAsetUpper {
	case "BTC", "BITCOIN":
		return "SHA-256"
	case "ETH", "ETHEREUM":
		return "ETHASH"
	case "LTC", "LITECOIN":
		return "SCRYPT"
	case "DG", "DOGECOIN":
		return "SCRYPT"
	case "XMR", "MONERO":
		return "RANDOMX"
	case "ADA", "CARDANO":
		return "OUROBOROS"
	case "SOL", "SOLANA":
		return "PROOF OF HISTORY"
	case "BCH", "BITCOIN CASH":
		return "SHA-256"
	case "SHIB", "SHIBA INU":
		return "ERC-20"
	default:
		var inputAlgoritma string
		for {
			inputAlgoritma = bacaString(fmt.Sprintf("Algoritma crypto untuk %s tidak dikenali, masukkan algoritma: ", namaAset))
			if inputAlgoritma != "" {
				return strings.ToUpper(inputAlgoritma)
			}
			fmt.Println("Tidak boleh kosong")
		}

	}
}

func inisialisasiDataAwalAset(aset *[1000]asetKrypto, jumlah *int) {
	if *jumlah == 0 {
		dataAwal := []struct {
			namaAset         string
			TingkatKesulitan int
			estimasiReward   float64
		}{
			{namaAset: "BITCOIN", TingkatKesulitan: 8, estimasiReward: 6.20},
			{namaAset: "ETHEREUM", TingkatKesulitan: 7, estimasiReward: 2.50},
			{namaAset: "LITECOIN", TingkatKesulitan: 5, estimasiReward: 12.50},
			{namaAset: "DOGECOIN", TingkatKesulitan: 3, estimasiReward: 100.00},
			{namaAset: "MONERO", TingkatKesulitan: 6, estimasiReward: 2.00},
		}
		for _, data := range dataAwal {
			if *jumlah < len(aset) {
				namaKapital := strings.ToUpper(data.namaAset)
				duplicate := false
				for i := 0; i < *jumlah; i++ {
					if aset[i].nama == namaKapital {
						duplicate = true
						break
					}
				}
				if !duplicate {
					aset[*jumlah].nama = namaKapital
					aset[*jumlah].tingkatKesulitan = data.TingkatKesulitan
					aset[*jumlah].estimasiReward = data.estimasiReward
					aset[*jumlah].algoritma = algoritmaCrypto(namaKapital)
					*jumlah++

				}
			} else {
				fmt.Println("⚠️kapasitas penuh")
				break
			}
		}
	}
}

func tampilkanInfoKriptoSistem() {
	fmt.Println("==Info Aset Kripto dan Algoritma Crypto==")
	if len(daftarInfoKriptoSistem) == 0 {
		fmt.Println("ℹ️ Daftar aset kripto dan algoritma crypto kosong")
		tekanEnterUntukLanjut()
		return
	}
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Printf("| %-20s | %-25s | %-20s | \n", " NAMA", "NAMA ALIAS", "ALGORITMA")
	fmt.Println(" -----------------------------------------------------------------------------")
	for _, info := range daftarInfoKriptoSistem {
		var aliasKapital []string
		for _, a := range info.namaAlias {
			aliasKapital = append(aliasKapital, strings.ToUpper(a))
		}
		aliasString := strings.Join(aliasKapital, ", ")
		fmt.Printf("| %-20s | %-25s | %-20s |\n", strings.ToUpper(info.namaLengkap), aliasString, info.algoritma)
	}
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("ℹ️ Jika nama aset tidak ada di daftar, input algorima secara manual")
	tekanEnterUntukLanjut()

}

// fungsi untuk menambahkan aset kripto ke dalam daftar
func tambahAset(aset *[1000]asetKrypto, jumlah *int) {
	fmt.Println("==Menu Tambah Aset Crypto==")
	if *jumlah >= len(aset) {
		fmt.Println("Jumlah aset sudah mencapai batas")
		tekanEnterUntukLanjut()
		return
	}
	var banyakAset int
	banyakAset = bacaInt("Tambahkan berapa jumlah aset yang anda mau (1-10 adalah jumlah): ", 1, 10)
	// perulangan sesuai banyak aset kripto
	for i := 0; i < banyakAset; i++ {
		if *jumlah >= len(aset) {
			fmt.Println("⚠️Jumlah aset sudah mencapai bata, tidak dapat menambah lagi")
			break
		}
		fmt.Printf("\n==Menambahkan aset ke %d dari %d ==\n", i+1, banyakAset)
		var nama, namaKapital string
		var tingkatKesulitan int
		var estimasiReward float64
		var algoritma string
		for {
			nama = bacaString("Masukkan nama aset kripto: ")
			if nama == "" {
				fmt.Println("Nama tidak boleh kosong")
				continue
			}
			namaKapital = strings.ToUpper(nama)
			duplicate := false
			for j := 0; j < *jumlah; j++ {
				if aset[j].nama == namaKapital {
					fmt.Println("nama sudah ada!!!! Masukkan nama yang lain!!!")
					duplicate = true
					break
				}
			}
			if !duplicate {
				break

			}
		}
		// input tingkat kesulitan dan estimasi reward (int dan float)
		tingkatKesulitan = bacaInt("Masukkan tingkat kesulitan (1-9): ", 1, 9)
		estimasiReward = bacaFloat("Masukkan estimasi reward per-block: ", 0.000001) // minimal input

		algoritma = algoritmaCrypto(namaKapital)
		// menyimpat aset keseluruhan yg di input sebelumnya
		aset[*jumlah].nama = namaKapital
		aset[*jumlah].tingkatKesulitan = tingkatKesulitan
		aset[*jumlah].estimasiReward = estimasiReward
		aset[*jumlah].algoritma = algoritma
		*jumlah++
		fmt.Println("✅ Aset berhasil ditambahkan")
	}
	tekanEnterUntukLanjut()

}

// fungsi untuk mengubah aset
func ubahAset(aset *[1000]asetKrypto, jumlah *int) {
	fmt.Println("==Menu Ubah Aset Crypto==")
	if *jumlah == 0 {
		fmt.Println("⚠️Aset tidak ada yang bisa diubah karena kosong")
		tekanEnterUntukLanjut()
		return
	}
	// menampilkan semua aset crypto agar mudah mengubah
	lihatAset(aset, jumlah)

	var indexUbah int
	indexUbah = bacaInt(fmt.Sprintf("Masukkan index aset yang ingin diubah (1-%d, 0 untuk batal): ", *jumlah), 0, *jumlah)
	for indexUbah == 0 {
		fmt.Println(" ⚠️Aset batal diibah karena index 0")
		tekanEnterUntukLanjut()
		return
	}
	indexUbah -= 1 // konversi ke index 0, array dimulai dari 0

	// konfirmasi ubah
	var konfirmasi string
	konfirmasi = bacaString("\nApakah anda yakin ingin mengubah aset ini? (ya/tidak): ")
	if strings.ToLower(konfirmasi) != "ya" {
		fmt.Println(" ℹ️Aset batal diubah karena pengguna tidak setuju")
		tekanEnterUntukLanjut()
		return
	}

	// menu pilihan ubah aset
	var pilihUbah int
	fmt.Println("==Menu Ubah Aset Crypto")
	fmt.Println("1. Ubah Semua Data Aset")
	fmt.Println("2. Ubah Nama Aset Saja")
	fmt.Println("3. Ubah Tingkat Kesulitan Saja")
	fmt.Println("4. Ubah Estimasi Reward Saja")
	fmt.Println("0. Batal pengubahan")
	pilihUbah = bacaInt("Pilih aset yang ingin diubah dari menu (0-4)", 0, 4)

	switch pilihUbah {
	//ubah semua
	case 1:
		var namaBaru, namaKapitalBaru string
		var tingkatKesulitanBaru int
		var estimasiRewardBaru float64
		var algoritmaBaru string

		for {
			namaBaru = bacaString(fmt.Sprintf("masukkan nama aset baru (sebelumnya: %s): ", aset[indexUbah].nama))
			if namaBaru == "" {
				fmt.Println("nama tidak boleh kosong")
				continue
			}
			namaKapitalBaru = strings.ToUpper(namaBaru)
			var cekNama = (namaKapitalBaru == aset[indexUbah].nama)
			duplicate := false
			if !cekNama {
				for k := 0; k < *jumlah; k++ {
					if k != indexUbah && aset[k].nama == namaKapitalBaru {
						duplicate = true
						break
					}
				}
			}
			if duplicate {
				fmt.Printf("⚠️ Nama aset %s sudah digunakan oleh aset lain\n", namaBaru)
			} else {
				break
			}
		}
		// ubah tingkat kesulitan, estimasi reward, algoritma pada aset kripto
		tingkatKesulitanBaru = bacaInt(fmt.Sprintf("masukkan tingkat kesulitan baru (0-9, sebelumnya %d): ", aset[indexUbah].tingkatKesulitan), 1, 9)
		estimasiRewardBaru = bacaFloat(fmt.Sprintf("Estimasi reward baru(>0, sebelumnya : %.2f: ", aset[indexUbah].estimasiReward), 0.0000001)
		algoritmaBaru = algoritmaCrypto(namaKapitalBaru)

		// konfirmasi ubah aset
		konfirmasi := bacaString("Simpan semua perubahan set ini? (ya/tidak): ")
		if strings.ToLower(konfirmasi) == "ya" {
			aset[indexUbah].nama = namaKapitalBaru
			aset[indexUbah].tingkatKesulitan = tingkatKesulitanBaru
			aset[indexUbah].estimasiReward = estimasiRewardBaru
			aset[indexUbah].algoritma = algoritmaBaru
			fmt.Println(" ✅Aset berhasil diubah")
		} else {
			fmt.Println("❌ Perubahan dibatalkan")
		}
	case 2:
		var namaBaru string
		var namaKapitalBaru string
		for {
			namaBaru = bacaString(fmt.Sprintf("Masukkan nama aset baru (sebelumnya: %s): ", aset[indexUbah].nama))
			if namaBaru == "" {
				fmt.Println("nama tidak boleh kosong")
				continue
			}
			namaKapitalBaru = strings.ToUpper(namaBaru)
			if namaKapitalBaru == aset[indexUbah].nama {
				fmt.Println("Nama baru tidak boleh sama dengan nama aset sebelumnya")
				return
			}
			duplicate := false
			for k := 0; k < *jumlah; k++ {
				if k != indexUbah && aset[k].nama == namaKapitalBaru {
					duplicate = true
					break
				}
			}
			if duplicate {
				fmt.Printf("⚠️ Nama aset %s sudah digunakan oleh aset lain\n", namaBaru)
			} else {
				break
			}
		}
		algoritmaBerubah := algoritmaCrypto(namaKapitalBaru)
		konfirmasiPerubahan := bacaString("Simpan perubahan nama aset ini? (ya/tidak): ")
		if strings.ToLower(konfirmasiPerubahan) == "ya" {
			aset[indexUbah].nama = namaKapitalBaru
			aset[indexUbah].algoritma = algoritmaBerubah
			fmt.Println("Nama aset berhasil diubah")
		} else {
			fmt.Println("Perubahan dibatalkan")
		}
	case 3:
		var tingkatKesulitanBaru int
		tingkatKesulitanBaru = bacaInt("Masukkan tingkat kesulitan baru (1-9): ", 1, 9)
		if tingkatKesulitanBaru == aset[indexUbah].tingkatKesulitan {
			fmt.Println("Tingkat kesulitan baru tidak boleh sama dengan tingkat kesulitan sebelumnya")
		} else {
			konfirmasiPerubahan := bacaString("Simpan perubahan tingkat kesulitan ini? (ya/tidak)")
			if strings.ToLower(konfirmasiPerubahan) == "ya" {
				aset[indexUbah].tingkatKesulitan = tingkatKesulitanBaru
				fmt.Println("✅Tingkat kesulitan aset berhasil diubah")
			} else {
				fmt.Println("Perubahan dibatalkan")
			}
		}
	case 4:
		var estimasiRewardBaru float64
		estimasiRewardBaru = bacaFloat("Estimasi reward baru (usd > 0): ", 0.1000000)
		if estimasiRewardBaru == aset[indexUbah].estimasiReward {
			fmt.Println("ℹ️ Estimasi reward baru tidak boleh sama dengan estimasi reward sebelumnya")
		} else {
			konfirmasiPerubahan := bacaString("Simpan perubahan estimasi reward ini? (ya/tidak)")
			if strings.ToLower(konfirmasiPerubahan) == "ya" {
				aset[indexUbah].estimasiReward = estimasiRewardBaru
				fmt.Println("✅Estimasi reward aset berhasil diubah")
			} else {
				fmt.Println("Perubahan dibatalkan")
			}
		}
	case 0:
		fmt.Println("ℹ️ Pengubahan dibatalkan.")
		break
	default:
		fmt.Println("Pilihan tidak tersedia. Silakan pilih opsi lainnya.")
	}
	tekanEnterUntukLanjut()
}

// Fungsi yang digunakan untuk menghapus aset crypto
func hapusAset(aset *[1000]asetKrypto, jumlah *int) {
	var pilihMenuHapus int
	fmt.Println("==Menu Hapus Aset Crypto==")
	if *jumlah == 0 {
		fmt.Println("⚠️Aset tidak ada yang bisa dihapus karena kosong")
		tekanEnterUntukLanjut()
		return
	}
	fmt.Println("1. Hapus aset crypto berdasarkan Nomor Urut dari tabel daftar aset")
	fmt.Println("2. Hapus aset crypto berdasarkan kategori Tingkat Kesulitan")
	fmt.Println("3. Hapus aset crypto berdasarkan kategori algoritma")
	fmt.Println("0. Kembali ke menu utama")
	pilihMenuHapus = bacaInt("Pilih menu hapus aset crypto: ", 0, 3)
	switch pilihMenuHapus {
	case 1:
		fmt.Println("-==Hapus berdasarkan No urut==")
		lihatAset(aset, jumlah)
		var indexHapus int
		indexHapus = bacaInt(fmt.Sprintf("Masukkan nomor yang ingin dihapus (1-%d, 0 untuk batal): ", *jumlah), 0, *jumlah)
		if indexHapus == 0 {
			fmt.Println("Pemnghapusan aset dibatalkan")
			tekanEnterUntukLanjut()
			return
		}
		indexHapus -= 1
		asetYangDihapus := aset[indexHapus].nama
		konfirmasiHapus := bacaString("Apakah ingin menghapus aset ini? (ya/tidak)")

		if strings.ToLower(konfirmasiHapus) == "ya" {
			// hapus aset dari array
			for i := indexHapus; i < *jumlah-1; i++ {
				aset[i] = aset[i+1]
			}
			aset[*jumlah-1] = asetKrypto{}
			*jumlah -= 1
			fmt.Printf(" ✅ Aset %s berhasil dihapus.\n", asetYangDihapus)
		} else {
			fmt.Println("ℹ️ Penghapusan dibatalkan.")
		}
		tekanEnterUntukLanjut()

	case 2:
		var piliTingkatSulit int
		var minKesulitan, maxKesulitan int
		fmt.Println("==Hapus berdasarkan tingkat kesulitan==")
		fmt.Println("1. Easy (Kesulitan 1-3)")
		fmt.Println("2. Medium (Kesulitan 4-6)")
		fmt.Println("3. Hard (Kesulitan 7-10)")
		fmt.Println("0. Batal")
		piliTingkatSulit = bacaInt("Pilih kategori tingkat kesulitan: ", 0, 3)
		if piliTingkatSulit == 0 {
			fmt.Println("Penghapusan dibatalkan")
			tekanEnterUntukLanjut()
			return
		}
		switch piliTingkatSulit {
		case 1:
			minKesulitan, maxKesulitan = 1, 3
		case 2:
			minKesulitan, maxKesulitan = 4, 6
		case 3:
			minKesulitan, maxKesulitan = 7, 10
		default:
			fmt.Println("Pilihan tidak valid")
			tekanEnterUntukLanjut()
			return
		}
		jumlahYangDihapus := 0
		i := *jumlah - 1
		for i >= 0 {
			if aset[i].tingkatKesulitan >= minKesulitan && aset[i].tingkatKesulitan <= maxKesulitan {
				fmt.Printf("Menghapus '%s' (Kesulitan: %d)...\n", aset[i].nama, aset[i].tingkatKesulitan)
				for k := i; k < *jumlah-1; k++ {
					aset[k] = aset[k+1]
				}
				aset[*jumlah-1] = asetKrypto{}
				*jumlah--
				jumlahYangDihapus++
			}
			i--
		}
		if jumlahYangDihapus > 0 {
			fmt.Printf(" ✅ %d aset berhasil dihapus.\n", jumlahYangDihapus)
		} else {
			fmt.Println("ℹ️ Penghapusan tidak ada.")
		}
		tekanEnterUntukLanjut()

	case 3:
		fmt.Println("==Hapus berdasarkan algoritma==")
		fmt.Println("Daftar algoritma yang tersedia:")
		mapAlgoritma := make(map[string]int)
		adaAlgoritma := false
		for i := 0; i < *jumlah; i++ {
			mapAlgoritma[aset[i].algoritma]++
		}
		if len(mapAlgoritma) > 0 {
			for algoritma, hitung := range mapAlgoritma {
				fmt.Printf("%s (%d)\n", algoritma, hitung)
			}
			adaAlgoritma = true
		}
		if !adaAlgoritma {
			fmt.Println("Tidak ada algoritma yang tersedia untuk dihapus.")
		}
		namaAlgoInput := bacaString(" Masukkan nama algoritma yang ingin dihapus 0 untuk batal: ")
		if namaAlgoInput == "0" || namaAlgoInput == "" {
			fmt.Println("Penghapusan batal")
			tekanEnterUntukLanjut()
			return
		}
		jumlahDihapus := 0
		i := *jumlah - 1
		for i >= 0 {
			if strings.ToUpper(aset[i].algoritma) == strings.ToUpper(namaAlgoInput) {
				fmt.Printf("Menghapus '%s' (Kesulitan: %d)...\n", aset[i].nama, aset[i].tingkatKesulitan)
				for k := i; k < *jumlah-1; k++ {
					aset[k] = aset[k+1]
				}
				aset[*jumlah-1] = asetKrypto{}
				*jumlah--
				jumlahDihapus++
			}
			i--
		}
		if jumlahDihapus > 0 {
			fmt.Printf(" ✅ %d aset berhasil dihapus.\n", jumlahDihapus)
		} else {
			fmt.Println("Penghapusan tidak ada.")
		}
		tekanEnterUntukLanjut()
	case 0:
		fmt.Println("Kembali ke menu utama")
		return

	default:
		fmt.Println("Pilihan tidak valid")
		tekanEnterUntukLanjut()
	}
}

// Fungsi untuk melihat aset crypto pengguna
func lihatAset(aset *[1000]asetKrypto, jumlah *int) {
	fmt.Println("==Menu Lihat Aset Crypto==")
	if *jumlah == 0 {
		fmt.Println("⚠️Aset tidak ada yang bisa dilihat karena kosong")
	}
	fmt.Println("---------------------------------------------------------------------------------------------------")
	fmt.Printf("| %-5s | %-20s | %-18s | %-18s | %-18s |\n", "Index", "Nama Aset", "Tingkat Kesulitan", "Reward (USD)", "AlgoritmaCrypto")
	fmt.Println("---------------------------------------------------------------------------------------------------")

	// Looping semua data
	for i := 0; i < *jumlah; i++ {
		fmt.Printf("| %-5d | %-20s | %-18d | %-18.2f | %-18s |\n",
			i+1,
			aset[i].nama,
			aset[i].tingkatKesulitan,
			aset[i].estimasiReward,
			aset[i].algoritma)
	}
	fmt.Println("---------------------------------------------------------------------------------------------------")
}

// Fungsi yang digunakan untuk mencari aset crypto menggunakan sequential dan binary, note: Binary lebih efektif dibandingkan dengan sequential

// Fungsi pencarian sequential berdasarkan nama aset kripto
func sequentialSearchNama(aset *[1000]asetKrypto, jumlah *int) {
	fmt.Println("==Menu Pencarian Sequential Berdasarkan Nama==")
	if *jumlah == 0 {
		fmt.Println("⚠️ Daftar aset kripto masih kosong!")
		return
	}

	var namaCari string
	namaCari = bacaString("Masukkan nama aset kripto yang ingin dicari: ")
	if namaCari == "" {
		fmt.Println("⚠️ Nama aset kripto tidak boleh kosong")
		return
	}
	namaCari = strings.ToUpper(namaCari)

	ditemukan := false
	fmt.Println("Hasil cari sequential: ")
	fmt.Println("---------------------------------------------------------------------------------------------------")
	fmt.Printf("| %-5s | %-20s | %-18s | %-18s | %-18s |\n", "Index", "Nama Aset", "Tingkat Kesulitan", "Reward (USD)", "AlgoritmaCrypto")
	fmt.Println("---------------------------------------------------------------------------------------------------")

	for i := 0; i < *jumlah; i++ {
		if strings.Contains(strings.ToUpper(aset[i].nama), namaCari) {
			fmt.Printf("| %-5d | %-20s | %-18d | %-18.2f | %-18s |\n",
				i+1,
				aset[i].nama,
				aset[i].tingkatKesulitan,
				aset[i].estimasiReward,
				aset[i].algoritma)
			ditemukan = true
		}
	}

	fmt.Println("---------------------------------------------------------------------------------------------------")

	if !ditemukan {
		fmt.Printf("⚠️ Aset kripto dengan nama \"%s\" tidak ditemukan!\n", namaCari)
	} else {
		fmt.Println("✅ Pencarian selesai")
	}
}

// Fungsi untuk mengurutkan array berdasarkan nama untuk binary search
func urutkanNamaAset(aset *[1000]asetKrypto, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		for j := 0; j < jumlah-i-1; j++ {
			if strings.ToUpper(aset[j].nama) > strings.ToUpper(aset[j+1].nama) {
				// Tukar posisi
				aset[j], aset[j+1] = aset[j+1], aset[j]
			}
		}
	}
}

// Fungsi binary search berdasarkan nama aset kripto
func binarySearchNama(aset *[1000]asetKrypto, jumlah *int) {
	fmt.Println("==Menu Pencarian Binary Search Berdasarkan Nama==")
	if *jumlah == 0 {
		fmt.Println("⚠️ Daftar aset kripto masih kosong!")
		return
	}

	// Membuat salinan array untuk diurutkan
	var asetUrut [1000]asetKrypto
	for i := 0; i < *jumlah; i++ {
		asetUrut[i] = aset[i]
	}

	// Urutkan array berdasarkan nama untuk binary search
	urutkanNamaAset(&asetUrut, *jumlah)

	var namaCari string
	namaCari = bacaString("Masukkan nama aset kripto yang ingin dicari: ")
	if namaCari == "" {
		fmt.Println("⚠️ Nama aset kripto tidak boleh kosong")
		return
	}
	namaCari = strings.ToUpper(namaCari)

	// Binary search
	kiri := 0
	kanan := *jumlah - 1
	ditemukan := false
	var indeksAsli int = -1
	var namaDiTengah string

	fmt.Println("Hasil pencarian binary search: ")
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		namaDiTengah = strings.ToUpper(asetUrut[tengah].nama)
		// Bandingkan string dalam bentuk uppercase
		if namaDiTengah == namaCari {
			// Temukan di array asli untuk mendapatkan indeks yang benar
			for i := 0; i < *jumlah; i++ {
				if aset[i].nama == asetUrut[tengah].nama {
					indeksAsli = i
					break
				}
			}
			fmt.Println("---------------------------------------------------------------------------------------------------")
			fmt.Printf("| %-5s | %-20s | %-18s | %-18s | %-18s |\n", "Index", "Nama Aset", "Tingkat Kesulitan", "Reward (USD)", "AlgoritmaCrypto")
			fmt.Println("---------------------------------------------------------------------------------------------------")
			fmt.Printf("| %-5d | %-20s | %-18d | %-18.2f | %-18s |\n",
				indeksAsli+1,
				aset[indeksAsli].nama,
				aset[indeksAsli].tingkatKesulitan,
				aset[indeksAsli].estimasiReward,
				aset[indeksAsli].algoritma)
			ditemukan = true
			break
		} else if strings.ToUpper(asetUrut[tengah].nama) < namaCari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	fmt.Println("---------------------------------------------------------------------------------------------------")

	if !ditemukan {
		fmt.Printf("⚠️ Aset kripto dengan nama \"%s\" tidak ditemukan!\n", namaCari)
	} else {
		fmt.Println("✅ Pencarian selesai")
	}
}

// Fungsi utama pencarian yang memanggil kedua fungsi di atas
func cariAset(aset *[1000]asetKrypto, jumlah *int) {
	if *jumlah == 0 {
		fmt.Println("⚠️ Daftar aset kripto masih kosong!!!, silahkan kembali ke menu utama")
		return
	}
	fmt.Println("==Menu Cari Aset Crypto==")
	fmt.Println("1. Cari aset Sequential")
	fmt.Println("2. Cari aset Binary")
	fmt.Println("0. Kembali ke menu utama")
	var pilihMenuCari int
	pilihMenuCari = bacaInt("Masukkan pilihan menu cari(0-2): ", 0, 2)
	switch pilihMenuCari {
	case 1:
		sequentialSearchNama(aset, jumlah)
	case 2:
		binarySearchNama(aset, jumlah)
	case 0:
		fmt.Println(" Kemabli ke menu utama")
		return
	default:
		fmt.Println("Pilihan tidak valid!")
	}
	tekanEnterUntukLanjut()
}

// Fungsi yang digunakan untuk mengurutkan aset crypto dengan selection dan isertion sort
func selectionTingkatKesulitan(aset *[1000]asetKrypto, jumlah *int) {
	fmt.Println("==Menu Urutkan Aset Secara Selection Berdasarkan Tingkat Kesulitan==")
	if *jumlah < 2 {
		fmt.Println("⚠️ Daftar aset kripto Tidak memenuhi syarat untuk diurutkan!!!")
		return
	}

	fmt.Println("Pilih jenis pengurutan:")
	fmt.Println("1. Ascending (kecil ke besar)")
	fmt.Println("2. Descending (besar ke kecil)")
	var pilihJenisUrut int
	pilihJenisUrut = bacaInt("Masukkan pilihan jenis urutan(1-2): ", 1, 2)

	// Selection sort algorithm
	n := *jumlah
	for i := 0; i < n-1; i++ {
		idxExtreme := i
		for j := i + 1; j < n; j++ {
			if pilihJenisUrut == 1 { // ascending
				if aset[j].tingkatKesulitan < aset[idxExtreme].tingkatKesulitan {
					idxExtreme = j
				}
			} else { // descending
				if aset[j].tingkatKesulitan > aset[idxExtreme].tingkatKesulitan {
					idxExtreme = j
				}
			}
		}
		// Swap nilai
		if idxExtreme != i {
			aset[i], aset[idxExtreme] = aset[idxExtreme], aset[i]
		}
	}

	// Menampilkan hasil pengurutan
	fmt.Println("✅ Aset berhasil diurutkan berdasarkan tingkat kesulitan")
	lihatAset(aset, jumlah)
}

// Implementasi Selection Sort berdasarkan estimasi reward
func selectionReward(aset *[1000]asetKrypto, jumlah *int) {
	fmt.Println("==Menu Urutkan Aset Secara Selection Berdasarkan Estimasi Reward==")
	if *jumlah < 2 {
		fmt.Println("⚠️ Daftar aset kripto tidak memenuhi syarat untuk diurutkan!!!(!=0 atau > 2)")
		return
	}

	fmt.Println("Pilih jenis pengurutan:")
	fmt.Println("1. Ascending (kecil ke besar)")
	fmt.Println("2. Descending (besar ke kecil)")
	var pilihJenisUrut int
	pilihJenisUrut = bacaInt("Masukkan pilihan jenis urutan(1-2): ", 1, 2)

	// Selection sort algorithm
	for i := 0; i < *jumlah-1; i++ {
		idxExtreme := i
		for j := i + 1; j < *jumlah; j++ {
			if pilihJenisUrut == 1 { // ascending
				if aset[j].estimasiReward < aset[idxExtreme].estimasiReward {
					idxExtreme = j
				}
			} else { // descending
				if aset[j].estimasiReward > aset[idxExtreme].estimasiReward {
					idxExtreme = j
				}
			}
		}
		// Swap nilai
		if idxExtreme != i {
			aset[i], aset[idxExtreme] = aset[idxExtreme], aset[i]
		}
	}

	// Menampilkan hasil pengurutan
	fmt.Println("✅ Aset berhasil diurutkan berdasarkan estimasi reward")
	lihatAset(aset, jumlah)
}

// Fungsi utama Selection Sort
func selectionSort(aset *[1000]asetKrypto, jumlah *int) {
	fmt.Println("==Menu Urutkan Selection Aset Crypto==")
	if *jumlah < 2 {
		fmt.Println("⚠️ Daftar aset kripto tidak memenuhi syarat untuk diurutkan!!!(!=0 atau > 2)")
	}
	fmt.Println("1. Urutkan berdasarkan Tingkat Kesulitan")
	fmt.Println("2. Urutkan berdasarkan Reward (USD)")
	fmt.Println("0. Kembali ke menu Pengurutan utama")
	var pilihMenuSelection int
	pilihMenuSelection = bacaInt("Masukkan pilihan menu selection(1-3): ", 0, 2)

	switch pilihMenuSelection {
	case 1:
		selectionTingkatKesulitan(aset, jumlah)
	case 2:
		selectionReward(aset, jumlah)
	case 0:
		fmt.Println("Kembali ke menu Pengurutan utama")
		return
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

// Implementasi Insertion Sort berdasarkan tingkat kesulitan
func insertionTingkatKesulitan(aset *[1000]asetKrypto, jumlah *int) {
	fmt.Println("==Menu Urutkan Aset Secara Insertion Berdasarkan Tingkat Kesulitan==")
	if *jumlah < 2 {
		fmt.Println("⚠️ Tidak memenuhi syarat pengurutan karena kurang dari 2 aset")
		return
	}

	fmt.Println("Pilih jenis pengurutan:")
	fmt.Println("1. Ascending (kecil ke besar)")
	fmt.Println("2. Descending (besar ke kecil)")
	var pilihJenisUrut int
	pilihJenisUrut = bacaInt("Masukkan pilihan jenis urutan(1-2): ", 1, 2)

	// Insertion sort algorithm
	for i := 1; i < *jumlah; i++ {
		key := aset[i]
		j := i - 1

		if pilihJenisUrut == 1 { // ascending
			for j >= 0 && aset[j].tingkatKesulitan > key.tingkatKesulitan {
				aset[j+1] = aset[j]
				j--
			}
		} else { // descending
			for j >= 0 && aset[j].tingkatKesulitan < key.tingkatKesulitan {
				aset[j+1] = aset[j]
				j--
			}
		}
		aset[j+1] = key
	}

	// Menampilkan hasil pengurutan
	fmt.Println("✅ Aset berhasil diurutkan berdasarkan tingkat kesulitan")
	lihatAset(aset, jumlah)
}

// Implementasi Insertion Sort berdasarkan estimasi reward
func insertionReward(aset *[1000]asetKrypto, jumlah *int) {
	fmt.Println("==Menu Urutkan Aset Secara Insertion Berdasarkan Estimasi Reward==")
	if *jumlah < 2 {
		fmt.Println("⚠️ Daftar aset tidak memenuhi syarat pengurutan karena kurang dari 2 aset")
		return
	}

	fmt.Println("Pilih jenis pengurutan:")
	fmt.Println("1. Ascending (kecil ke besar)")
	fmt.Println("2. Descending (besar ke kecil)")
	var pilihJenisUrut int
	pilihJenisUrut = bacaInt("Masukkan pilihan jenis urutan(1-2): ", 1, 2)

	// Insertion sort algorithm
	for i := 1; i < *jumlah; i++ {
		key := aset[i]
		j := i - 1
		if pilihJenisUrut == 1 { // ascending
			for j >= 0 && aset[j].estimasiReward > key.estimasiReward {
				aset[j+1] = aset[j]
				j--
			}
		} else { // descending
			for j >= 0 && aset[j].estimasiReward < key.estimasiReward {
				aset[j+1] = aset[j]
				j--
			}
		}
		aset[j+1] = key
	}

	// Menampilkan hasil pengurutan
	fmt.Println("✅ Aset berhasil diurutkan berdasarkan estimasi reward")
	lihatAset(aset, jumlah)
}

// Fungsi utama Insertion Sort
func insertionSort(aset *[1000]asetKrypto, jumlah *int) {
	fmt.Println("==Menu Urutkan Insertion Aset Crypto==")
	if *jumlah < 2 {
		fmt.Println("⚠️ Daftar aset tidak memenuhi syarat pengurutan karena kurang dari 2 aset")
	}
	fmt.Println("1. Urutkan berdasarkan Tingkat Kesulitan")
	fmt.Println("2. Urutkan berdasarkan Reward (USD)")
	fmt.Println("0. Kembali ke menu utama")
	var pilihMenuInsertion int
	pilihMenuInsertion = bacaInt("Masukkan pilihan menu (0-2): ", 0, 2)

	switch pilihMenuInsertion {
	case 1:
		insertionTingkatKesulitan(aset, jumlah)
	case 2:
		insertionReward(aset, jumlah)
	case 0:
		return
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

// Fungsi utama pengurutan
func urutkanAset(aset *[1000]asetKrypto, jumlah *int) {
	fmt.Println("==Menu Urutkan Aset Crypto==")
	if *jumlah < 2 {
		fmt.Println("⚠️ Daftar aset tidak memenuhi syarat pengurutan karena kurang dari 2 aset")
		return
	}
	fmt.Println("1. Urutkan aset dengan Selection Sort")
	fmt.Println("2. Urutkan aset dengan Insertion Sort")
	fmt.Println("0. Kembali ke menu utama")
	var pilihMenuUrutAset int
	pilihMenuUrutAset = bacaInt("Masukkan pilihan menu (0-2): ", 0, 2)

	switch pilihMenuUrutAset {
	case 1:
		selectionSort(aset, jumlah)
	case 2:
		insertionSort(aset, jumlah)
	case 0:
		return
	default:
		fmt.Println("Pilihan tidak valid!")
	}
	tekanEnterUntukLanjut()
}

// Fungsi yang digunakan untuk simulasi mining crypto pengguna
func simulasiMining(aset *[1000]asetKrypto, jumlahAset *int, riwayat *[1000]riwayatMining, riwayatJumlah *int) {
	fmt.Println("==Menu Simulasi Mining!!!==")
	if *jumlahAset == 0 {
		fmt.Println("⚠️ Daftar aset kripto masih kosong! Silahkan tambahkan terlebih dahulu")
		tekanEnterUntukLanjut()
		return
	}
	if *riwayatJumlah >= len(*riwayat) {
		fmt.Println("⚠️ Riwayat mining sudah penuh! Tidak dapat menambahkan lagi")
		tekanEnterUntukLanjut()
		return
	}
	lihatAset(aset, jumlahAset)

	var pilihNomorAset int
	var asetPilih asetKrypto
	pilihNomorAset = bacaInt(fmt.Sprintf("Pilih nomor urut yang ingin ditambang (1-%d, 0 untuk batal)", *jumlahAset), 0, *jumlahAset)
	if pilihNomorAset == 0 {
		fmt.Println(" ⚠️ Anda memilih untuk batal")
		tekanEnterUntukLanjut()
		return
	}
	asetPilih = aset[pilihNomorAset-1]

	var dayaKomputasiPengguna float64
	dayaKomputasiPengguna = bacaFloat("Masukkan daya komputasi pengguna (dalam MH/s, angka > 0): ", 0.00000001)

	var durasiSimulasi int
	durasiSimulasi = bacaInt("Masukkan durasi simulasi (dalam detik , 10-3600, > 0): ", 1, 3600)
	fmt.Printf("\nMemulai simulasi untuk %s (%d detik, daya %.2f).....\n", asetPilih.nama, durasiSimulasi, dayaKomputasiPengguna)

	faktorKesulitan := float64(asetPilih.tingkatKesulitan) * 500.0
	bobotAlgoritma := 1.0
	switch strings.ToLower(asetPilih.algoritma) {
	case "sha-256":
		bobotAlgoritma = 1.0
	case "srypt":
		bobotAlgoritma = 1.2
	case "ethash":
		bobotAlgoritma = 1.1
	case "randomx":
		bobotAlgoritma = 1.5
	default:
		bobotAlgoritma = 1.0
	}

	// Estimasi detik yang dibutuhkan untuk menambang 1 block
	estimasiDetikPerBlok := (faktorKesulitan * bobotAlgoritma) / (dayaKomputasiPengguna / 20.0)
	if estimasiDetikPerBlok < 0.5 {
		estimasiDetikPerBlok = 0.5
	}
	fmt.Printf("Estimasi sistem : %.2f detik/block untuk algoritma %s dengan daya %.2f \n", estimasiDetikPerBlok, asetPilih.algoritma, dayaKomputasiPengguna)

	var totalBlokDitemukan = 0
	waktuSimulasi := time.Now()

	fmt.Println("Proses Mining aset..........")
	for sisaWaktu := durasiSimulasi; sisaWaktu > 0; sisaWaktu-- {
		keberuntunganDetikIni := (rand.Float64() * 0.4) + 0.8
		dayaEfektifDetikIni := dayaKomputasiPengguna * keberuntunganDetikIni

		peluangBlokPerDetik := dayaEfektifDetikIni / (float64(asetPilih.tingkatKesulitan*asetPilih.tingkatKesulitan) * bobotAlgoritma * 50.0)

		if rand.Float64() < peluangBlokPerDetik {
			totalBlokDitemukan++
		}
		fmt.Printf("\rSisa waktu: %d detik..... mencari blok....(Perkiraan blok ditemukan: %d) ", sisaWaktu, totalBlokDitemukan)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("\n Simulasi Mining  selesai")

	totalRewardSimulasi := totalBlokDitemukan * int(asetPilih.estimasiReward)

	fmt.Printf("\n==Hasil Simulasi==\n")
	fmt.Printf("Aset Ditambang         : %s\n", asetPilih.nama)
	fmt.Printf("Durasi Simulasi        : %d detik\n", durasiSimulasi)
	fmt.Printf("Daya Komputasi Digunakan: %.2f MH/s\n", dayaKomputasiPengguna)
	fmt.Printf("Perkiraan Blok Ditemukan: %d blok\n", totalBlokDitemukan)
	fmt.Printf("Total Estimasi Reward  : %d %s\n", totalRewardSimulasi, asetPilih.nama)

	//Simpan hasil ke riwayatSimulasi
	if *riwayatJumlah < len(*riwayat) {
		var periodeSimulasiMining = time.Now()
		(*riwayat)[*riwayatJumlah] = riwayatMining{
			namaAset:              asetPilih.nama,
			jumlahReward:          float64(totalRewardSimulasi),
			waktuSelesai:          time.Now(),
			durasiSimulasi:        time.Since(waktuSimulasi).Round(time.Second),
			dayaKomputasiPengguna: dayaKomputasiPengguna,
			blokDitemukan:         float64(totalBlokDitemukan),
			periodeSimulasi:       periodeSimulasiMining,
		}
		*riwayatJumlah++
		fmt.Println("✅ Hasil simulasi berhasil disimpan ke riwayat.")
	} else {
		fmt.Println("❌ Riwayat simulasi sudah penuh.")
	}
	tekanEnterUntukLanjut()
}

// Fungs yang digunakan  untuk menampilkan laporan total mining
func laporanTotalMining(riwayat *[1000]riwayatMining, riwayatJumlah *int) {
	fmt.Println("==Menu Laporan Total Mining==")
	fmt.Println("Berikut adalah laporan total mining anda dalam periode tertentu: ")
	if *riwayatJumlah == 0 {
		fmt.Println("Tidak ada riwayat simulasi mining")
		tekanEnterUntukLanjut()
		return
	}
	fmt.Println("Berikut adalah laporan total mining anda:")
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("| %-4s | %-20s | %-15s | %-20s | %-15s | %-15s | %-20s |\n",
		"No.", "Aset Kripto", "Reward", "Waktu Selesai", "Durasi (s)", "Daya (MH/s)", "Blok Ditemukan",
	)
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------")

	var totalRewardKeseluruhan float64 = 0
	var totalBlokKeseluruhan float64 = 0

	for i := 0; i < *&jumlahRiwayatMining; i++ {
		data := (*riwayat)[i]
		fmt.Printf("| %-4d | %-20s | %-15.4f | %-20s | %-15.0f | %-15.2f | %-20.2f |\n",
			i+1,
			data.namaAset,
			data.jumlahReward,
			data.waktuSelesai.Format("2006-01-02 15:04:05"), // Format waktu selesai
			data.durasiSimulasi.Seconds(),                   // Durasi dalam detik
			data.dayaKomputasiPengguna,
			data.blokDitemukan,
		)
		totalRewardKeseluruhan += data.jumlahReward
		totalBlokKeseluruhan += data.blokDitemukan
	}
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("| %-27s | %-15.4f | %-55s | %-20.2f |\n",
		"TOTAL KESELURUHAN:", totalRewardKeseluruhan, "", totalBlokKeseluruhan)
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------")
	tekanEnterUntukLanjut()
}

func main() {
	inisialisasiDataAwalAset(&daftarAsetKripto, &jumlahAset)

	var pilihMenuUtama int
	for {
		fmt.Println()
		fmt.Println("|||=======================================|||")
		fmt.Printf("||| WELCOME TO CRYPTO MINING SIMULATOR APP|||\n")
		fmt.Println("|||=======================================|||")
		fmt.Println("|||=======Manajemen Aset Crypto===========|||")
		fmt.Println("||| 1.➕Menambah aset kripto             |||")
		fmt.Println("||| 2.✏️ Mengubah aset kripto              |||")
		fmt.Println("||| 3.🗑️ Menghapus aset kripto             |||")
		fmt.Println("||| 4.📋 Melihat aset kripto              |||")
		fmt.Println("|||=======================================|||")
		fmt.Println("|||=======Menu Analisis Data Crpto========||| ")
		fmt.Println("||| 5.🔍 Mencari aset Kripto              |||")
		fmt.Println("||| 6.📊 Pengurutan aset kripto           |||")
		fmt.Println("||| 7.⛏️ Mulai simulasi mining             |||")
		fmt.Println("||| 8.💰 laporan total hasil mining       |||")
		fmt.Println("||| 9.📚 Info Aset dan Algoritma sistem   |||")
		fmt.Println("||| 0.🚪 Exit                             |||")
		fmt.Println("|||=======================================|||")
		pilihMenuUtama = bacaInt("Pilih Menu yangg anda mau(0-9): ", -100000, 1000000)
		switch pilihMenuUtama {
		case 1:
			tambahAset(&daftarAsetKripto, &jumlahAset)
		case 2:
			ubahAset(&daftarAsetKripto, &jumlahAset)
		case 3:
			hapusAset(&daftarAsetKripto, &jumlahAset)
		case 4:
			lihatAset(&daftarAsetKripto, &jumlahAset)
			tekanEnterUntukLanjut()
		case 5:
			cariAset(&daftarAsetKripto, &jumlahAset)
		case 6:
			urutkanAset(&daftarAsetKripto, &jumlahAset)
		case 7:
			simulasiMining(&daftarAsetKripto, &jumlahAset, &daftarRiwayatMining, &jumlahRiwayatMining)
		case 8:
			laporanTotalMining(&daftarRiwayatMining, &jumlahRiwayatMining)
		case 9:
			tampilkanInfoKriptoSistem()
		case 0:
			// kembali ke menu login awal
			fmt.Println("Terima kasih telah menggunakan aplikasi Crypto Mining Simulator App")
			fmt.Println("")
			os.Exit(0)
		default:
			fmt.Println("Maaf, pilihan anda tidak sesuai denga menu yang ada, silahkan coba lagi..........")
			tekanEnterUntukLanjut()
		}
	}

}

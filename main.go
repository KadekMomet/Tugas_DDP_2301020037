package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Karyawan adalah struktur data untuk menyimpan informasi karyawan
type Karyawan struct {
	Nama                string
	PunyaIstri          bool
	PunyaAnak           bool
	JumlahAnak          int
	JumlahPengantaran   int
	PengantaranBerhasil int
	TujuanPengantaran   []string
}

//  menghitung gaji karyawan berdasarkan aturan tertentu
func (k *Karyawan) HitungGaji() int {
	// Konstanta gaji pokok, tunjangan istri, dan upah pengantaran
	const gajiPokok = 4000000
	const tunjanganIstri = 1000000
	const upahPengantaran = 10000

	gaji := gajiPokok

	if k.PunyaIstri {
		gaji += tunjanganIstri
	}
	gaji += k.JumlahAnak * 500000
	gaji += k.JumlahPengantaran * upahPengantaran

	if k.PengantaranBerhasil > 50 {
		gaji += 1000000
	} else if k.PengantaranBerhasil > 10 {
		gaji += 500000
	}

	if k.PengantaranBerhasil < 5 {
		pemotongan := 500000
		gaji -= pemotongan
	}

	return gaji
}

//  menambahkan pengantaran baru ke karyawan
func (k *Karyawan) TambahPengantaran(berhasil bool) {
	k.JumlahPengantaran++
	if berhasil {
		k.PengantaranBerhasil++
	}
}

//  menampilkan informasi karyawan
func (k *Karyawan) TampilkanInfo() string {
	info := fmt.Sprintf("Nama Karyawan: %s\nKeterangan: %s dan memiliki %d anak\n", k.Nama, func() string {
		if k.PunyaIstri {
			return "Sudah Menikah"
		}
		return "Belum Menikah"
	}(), k.JumlahAnak)

	if len(k.TujuanPengantaran) > 0 {
		info += "Tujuan Pengantaran:\n"
		for i, tujuan := range k.TujuanPengantaran {
			info += fmt.Sprintf("%d. %s\n", i+1, tujuan)
		}
	}

	return info
}


func main() {
	// Membuat objek karyawan

	for {
		fmt.Println("Program Penggajian Karyawan")
		fmt.Println("==========================")
		fmt.Println("1. Tambah Karyawan")
		fmt.Println("2. Cari Karyawan")
		fmt.Println("3. Hapus Karyawan")
		fmt.Println("4. Tambah Pengantaran")
		fmt.Println("5. Cari Pengantaran")
		fmt.Println("6. Hapus Pengantaran")
		fmt.Println("7. Hitung Gaji")
		fmt.Println("8. Keluar")

		var pilihan int
		fmt.Print("Pilih Menu [1-8]: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahKaryawan()
		case 2:
			CariKaryawan()
		case 3:
			HapusKaryawan()
		case 4:
			TambahPengantaran()
		case 5:
			CariPengantaran()
		case 6:
			HapusPengantaran()
		case 7:
			HitungGaji()
		case 8:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}

//  menambahkan karyawan baru ke data
func TambahKaryawan() {
	var nama string
	var punyaIstri, punyaAnak string

	fmt.Print("Nama Karyawan: ")
	fmt.Scan(&nama)
	fmt.Print("Punya Istri (y/n): ")
	fmt.Scan(&punyaIstri)

	var jumlahAnak int
	if punyaIstri == "y" {
		fmt.Print("Punya Anak (y/n): ")
		fmt.Scan(&punyaAnak)
		if punyaAnak == "y" {
			fmt.Print("Masukkan jumlah anak: ")
			fmt.Scan(&jumlahAnak)
		}
	}

	karyawanBaru := Karyawan{
		Nama:       nama,
		PunyaIstri: punyaIstri == "y",
		PunyaAnak:  punyaAnak == "y",
		JumlahAnak: jumlahAnak,
	}
	saveData(karyawanBaru)
	fmt.Println("Karyawan berhasil ditambahkan!")
}

//  mencari karyawan berdasarkan nama
func CariKaryawan() {
	saveDatas := loadData()
	var namaCari string
	fmt.Print("Masukkan nama karyawan: ")
	fmt.Scan(&namaCari)

	var karyawan *Karyawan
	for i := range saveDatas {
		if saveDatas[i].Nama == namaCari {
			karyawan = &saveDatas[i]
			break
		}
	}

	if karyawan != nil {
		fmt.Print(karyawan.TampilkanInfo())
	} else {
		fmt.Println("Karyawan tidak ditemukan!")
	}
}

//  menghapus karyawan berdasarkan indeks
func HapusKaryawan() {
	saveDatas := loadData()
	fmt.Println("Daftar Karyawan:")
	for i := range saveDatas {
		fmt.Printf("%d. %s\n", i+1, saveDatas[i].Nama)
	}

	var indexHapus int
	fmt.Print("Pilih Karyawan yang akan dihapus [1-", len(saveDatas), "]: ")
	fmt.Scan(&indexHapus)
	indexHapus--

	if indexHapus >= 0 && indexHapus < len(saveDatas) {
		saveDatas = append(saveDatas[:indexHapus], saveDatas[indexHapus+1:]...)
		saveDataA(saveDatas)
		fmt.Println("Karyawan berhasil dihapus!")
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

//  menambahkan pengantaran baru ke karyawan
func TambahPengantaran() {
	saveDatas := loadData()
	fmt.Println("Daftar Karyawan:")
	for i, karyawan := range saveDatas {
		fmt.Printf("%d. %s\n", i+1, karyawan.Nama)
	}

	var indexPilih int
	fmt.Print("Pilih Karyawan yang akan melakukan pengantaran [1-", len(saveDatas), "]: ")
	fmt.Scan(&indexPilih)
	indexPilih--

	if indexPilih >= 0 && indexPilih < len(saveDatas) {

		var tujuan string
		fmt.Print("Masukkan tujuan pengantaran: ")
		fmt.Scan(&tujuan)

		var berhasil string
		fmt.Print("Apakah pengantaran berhasil (y/n): ")
		fmt.Scan(&berhasil)

		karyawan := saveDatas[indexPilih]
		karyawan.TambahPengantaran(berhasil == "y")

		// Mengosongkan TujuanPengantaran sebelum menambahkan tujuan baru
		karyawan.TujuanPengantaran = nil

		if berhasil == "y" {
			karyawan.TujuanPengantaran = append(karyawan.TujuanPengantaran, tujuan)
		}

		saveData(karyawan)
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

//  mencari pengantaran karyawan berdasarkan nama
func CariPengantaran() {
	saveDatas := loadData()
	fmt.Println("Daftar Karyawan:")
	for i := range saveDatas {
		fmt.Printf("%d. %s\n", i+1, saveDatas[i].Nama)
	}

	var indexCari int
	fmt.Print("Pilih Karyawan yang akan dicari pengantarannya [1-", len(saveDatas), "]: ")
	fmt.Scan(&indexCari)
	indexCari--

	if indexCari >= 0 && indexCari < len(saveDatas) {
		karyawan := &(saveDatas)[indexCari]
		fmt.Printf("Nama Karyawan: %s\n", karyawan.Nama)
		fmt.Printf("Jumlah Pengantaran: %d\n", karyawan.JumlahPengantaran)
		fmt.Printf("Jumlah Pengantaran Berhasil: %d\n", karyawan.PengantaranBerhasil)
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

// menghapus pengantaran karyawan berdasarkan indeks
func HapusPengantaran() {
	saveDatas := loadData()
	fmt.Println("Daftar Karyawan:")
	for i := range saveDatas {
		fmt.Printf("%d. %s\n", i+1, saveDatas[i].Nama)
	}

	var indexHapusPengantaran int
	fmt.Print("Pilih Karyawan yang akan dihapus pengantarannya [1-", len(saveDatas), "]: ")
	fmt.Scan(&indexHapusPengantaran)
	indexHapusPengantaran--

	if indexHapusPengantaran >= 0 && indexHapusPengantaran < len(saveDatas) {
		karyawan := saveDatas[indexHapusPengantaran]
		karyawan.JumlahPengantaran = 0
		karyawan.PengantaranBerhasil = 0
		saveData(karyawan)
		fmt.Println("Pengantaran berhasil dihapus!")
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

// menghitung gaji karyawan berdasarkan indeks
func HitungGaji() {
	saveDatas := loadData()
	fmt.Println("Daftar Karyawan:")
	for i := range saveDatas {
		fmt.Printf("%d. %s\n", i+1, saveDatas[i].Nama)
	}

	var indexHitungGaji int
	fmt.Print("Pilih Karyawan yang akan dihitung gajinya [1-", len(saveDatas), "]: ")
	fmt.Scan(&indexHitungGaji)
	indexHitungGaji--

	if indexHitungGaji >= 0 && indexHitungGaji < len(saveDatas) {
		karyawan := saveDatas[indexHitungGaji]
		gaji := karyawan.HitungGaji()
		fmt.Printf("Gaji %s adalah Rp %d\n", karyawan.Nama, gaji)
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

// saveDataA menyimpan data ke file JSON
func saveDataA(data []Karyawan) {
	wd, _ := os.Getwd()
	jsonData, jsonError := json.Marshal(data)
	if jsonError != nil {
		log.Fatalln("Can't Marshal the Data")
	}
	writeError := os.WriteFile(fmt.Sprintf("%s/%s", wd, "save.json"), jsonData, os.ModePerm)
	if writeError != nil {
		log.Fatalln("Can't write the file")
	}
}

// saveData menyimpan data karyawan ke dalam file JSON
func saveData(data Karyawan) {
	var temps []Karyawan
	wd, _ := os.Getwd()
	oldfile, err := os.ReadFile(fmt.Sprintf("%s/%s", wd, "save.json"))
	if os.IsNotExist(err) {
		file, errs := os.Create(fmt.Sprintf("%s/%s", wd, "save.json"))
		if errs != nil {
			log.Fatalln("Error while Creating file")
		}
		temps = append(temps, data)
		jsonData, jsonError := json.Marshal(temps)
		if jsonError != nil {
			log.Fatalln("Failed to Marshal the Data")
		}
		_, writeError := file.Write(jsonData)
		if writeError != nil {
			log.Fatalln("Failed to Write the Data")
		}
	}
	if err != nil {
		log.Fatalln("Failed to Open the file")
	}
	
	jsonErr := json.Unmarshal(oldfile, &temps)
	if jsonErr != nil {
		log.Fatalln("Failed to Unmarshal the data")
	}
	temps = append(temps, data)
	jsonData, jsonError := json.Marshal(temps)
	if jsonError != nil {
		fmt.Println("Failed to Marshal the data")
	}
	writeError := os.WriteFile(fmt.Sprintf("%s/%s", wd, "save.json"), jsonData, os.ModePerm)
	if writeError != nil {
		log.Fatalln("Can't write the file")
	}
}

// loadData membaca data dari file JSON dan mengembalikan slice Karyawan
func loadData() []Karyawan {
	var temp []Karyawan
	wd, _ := os.Getwd()
	oldfile, err := os.ReadFile(fmt.Sprintf("%s/%s", wd, "save.json"))
	if os.IsNotExist(err) {
		log.Fatalln("Can't read the file!, is the file already exists?")
	}
	jsonError := json.Unmarshal(oldfile, &temp)
	if jsonError != nil {
		log.Fatalln("Can't unmarshal the data")
	}
	return temp
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Soal struct {
	Pertanyaan   string
	Pilihan      []string
	JawabanBenar int
}

type Kuis struct {
	SoalSoal []Soal
	Skor     int
}

func main() {
	// Menginput nama dan nim user
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama Anda: ")
	scanner.Scan()
	nama := scanner.Text()
	fmt.Print("Masukkan nim anda: ")
	scanner.Scan()
	nim := scanner.Text()
	fmt.Println()

	// Daftar soal yang akan ditanyakan
	soal1 := Soal{
		Pertanyaan:   "Siapakah presiden pertama Indonesia?",
		Pilihan:      []string{"0. Soekarno", "1. Bj. Habibie", "2. Megawati", "3. Jokowi"},
		JawabanBenar: 0,
	}

	soal2 := Soal{
		Pertanyaan:   "Pada tanggal berapakah hari lahir pancasila diperingati?",
		Pilihan:      []string{"0. 17 Agustus", "1. 1 Maret", "2. 1 Juni", "3. 28 Oktober"},
		JawabanBenar: 2,
	}

	soal3 := Soal{
		Pertanyaan:   "Apa nama mata uang dari negara Thailand?",
		Pilihan:      []string{"0. Yen", "1. Won", "2. Dollar", "3. Baht"},
		JawabanBenar: 3,
	}

	soal4 := Soal{
		Pertanyaan:   "Flora yang sangat terkenal di negara jepang adalah jenis bunga yang dikenal dengan nama bunga?",
		Pilihan:      []string{"0. Tulip", "1. Sakura", "2. Edelweis", "3. Reflesia arnodi"},
		JawabanBenar: 1,
	}

	soal5 := Soal{
		Pertanyaan:   "Sungai terpanjang di dunia adalah?",
		Pilihan:      []string{"0. Ganga", "1. Nil", "2. kapuas", "3. Amazon"},
		JawabanBenar: 1,
	}

	// Menyimpan soal-soal dalam slice Soal
	soal := []Soal{soal1, soal2, soal3, soal4, soal5}

	// Inisialisasi kuis
	kuis := Kuis{
		SoalSoal: soal,
		Skor:     0,
	}

	// Proses kuis
	for i, s := range kuis.SoalSoal {
		fmt.Println("Soal ke", i+1)
		fmt.Println(s.Pertanyaan)
		for _, p := range s.Pilihan {
			fmt.Println(p)
		}

		// Input jawaban
		fmt.Print("Masukkan pilihan jawaban Anda (0-3): ")
		scanner.Scan()
		jawabanStr := scanner.Text()
		jawaban := 0
		fmt.Sscanf(jawabanStr, "%d", &jawaban)

		// Mengecek jawaban user
		if jawaban == s.JawabanBenar {
			fmt.Println("Jawaban Anda benar!")
			kuis.Skor++
		} else {
			fmt.Println("Jawaban Anda salah.")
		}

		fmt.Println()
	}

	// Menampilkan hasil kuis yang telah dikerjakan oleh user
	fmt.Println("Statistic kuis")
	fmt.Printf("Nama         : %s\n", nama)
	fmt.Printf("Nim          : %s\n", nim)
	fmt.Printf("Skor         : %d\n", 20*kuis.Skor)
	fmt.Printf("Jawaban Benar: %d\n", kuis.Skor)
	fmt.Printf("Jawaban Salah: %d\n", 5-kuis.Skor)

}

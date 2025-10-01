package main

import "fmt"

var (
	minatList    = []string{}
	keahlianList = []string{}
)

type Karir struct {
	Nama      string
	Kecocokan int
	Industri  string
	Gaji      string
}

type Pekerjaan struct {
	Nama string
	Gaji string
}

func main() {
	var keluar bool
	var pilihan int

	for !keluar {
		fmt.Println("\nMenu Pilihan:")
		fmt.Println("1. Tambah minat")
		fmt.Println("2. Tambah keahlian")
		fmt.Println("3. Edit minat")
		fmt.Println("4. Edit keahlian")
		fmt.Println("5. Karir rekomendasi")
		fmt.Println("6. Tampilkan Minat Dan Keahlian ")
		fmt.Println("7. Pekerjaan ")
		fmt.Println("8. Keluar")
		fmt.Print("Masukkan pilihan (1-8): ")
		fmt.Scan(&pilihan)
		if pilihan >= 9 {
			fmt.Println("Input tidak valid, silakan coba lagi")
			continue
		}

		switch pilihan {
		case 1:
			tambahMinat()
		case 2:
			tambahKeahlian()
		case 3:
			editMinat()
		case 4:
			editKeahlian()
		case 5:
			karirRekomendasi()
		case 6:
			tampilMinatDanKeahlian(minatList, keahlianList)
		case 7:
			pekerjaan()
		case 8:
			fmt.Println("Terima kasih, program selesai.")
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid, silakan pilih 1-8")
		}
	}
}

func tambahMinat() {
	var input string

	daftarMinat := []string{
		"Membaca",
		"Menulis",
		"Menggambar",
		"Fotografi",
		"Menghitung",
		"Menganalisis",
		"Melukis",
		"Bermain alat musik",
	}

	fmt.Println("\n=== Tambah Minat ===")
	fmt.Println("Daftar minat yang tersedia:")
	for i, minat := range daftarMinat {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	sisa := 3 - len(minatList)
	if sisa <= 0 {
		fmt.Println("\nYahh kamu sudah menambahkan 3 minat")
		return
	}

	fmt.Printf("\nKamu bisa mengisi %d minat (ketik '-' untuk menyudahi tambah minat)\n", sisa)
	for i := 0; i < sisa; i++ {
		fmt.Printf("Pilihan %d (1-8 atau - untuk berhenti): ", i+1)
		fmt.Scan(&input)

		if input == "-" {
			fmt.Println("Kamu Menghentikan Penambahan.")
			break
		}

		var pilihan int
		_, err := fmt.Sscan(input, &pilihan)
		if err != nil || pilihan < 1 || pilihan  > len(daftarMinat) {
			fmt.Println("Input tidak valid")
			return
		}

		duplikat := false
		for _, m := range minatList {
			if m == daftarMinat[pilihan-1] {
				duplikat = true
				break
			}
		}
		if duplikat {
			fmt.Println("Minat sudah ada, coba yang lain.")
			i--
			continue
		}

		minatList = append(minatList, daftarMinat[pilihan-1])
		fmt.Println("--Minat berhasil ditambahkan--")
	}

	fmt.Println("Minat kamu sekarang:", minatList)
}

func tambahKeahlian() {
	var pilihan, sisa int
	sisa = 1 - len(keahlianList)

	fmt.Println("\n=== Tambah Keahlian ===")
	daftarKeahlian := []string{
		"Coding",
		"Analisis",
		"Design Grafis",
		"Editing Video",
		"Menghafal",
		"Mengamati",
		"Melukis",
	}

	fmt.Println("Daftar Keahlian yang tersedia:")
	for i, ahli := range daftarKeahlian {
		fmt.Printf("%d. %s\n", i+1, ahli)
	}

	if sisa <= 0 {
		fmt.Println("\nYahh kamu sudah menambahkan 1 keahlian")
		return
	}

	fmt.Printf("\nKamu bisa mengisi %d keahlian\n", sisa)
	for i := 0; i < sisa; i++ {
		fmt.Printf("Pilihan %d (1-7): ", i+1)
		fmt.Scan(&pilihan)

		if pilihan < 1 || pilihan > 7 {
			fmt.Println("Input tidak valid")
			return
		}

		keahlianList = append(keahlianList, daftarKeahlian[pilihan-1])
		fmt.Println("--Keahlian berhasil ditambahkan--")
	}
}

func editMinat() {
	var pilihan int

	fmt.Println("\n=== Edit Minat ===")
	if len(minatList) == 0 {
		fmt.Println("Belum ada minat yang tersedia")
		return
	}
	fmt.Println("Minat Anda saat ini:")
	for i, minat := range minatList {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	fmt.Println("\n pilih menu dibawah ini:")
	fmt.Println("1. Tambah Minat")
	fmt.Println("2. Hapus Minat (pilih 1)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilihan kamu: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahMinat()
	case 2:
		hapusMinat()
	case 3:
		return
	default:
		fmt.Println("Pilihan Tidak Valid")
	}
}

func editKeahlian() {
	var pilihan int

	fmt.Println("\n=== Edit Keahlian ===")
	if len(keahlianList) == 0 {
		fmt.Println("Belum ada Keahlian yang tersedia")
		return
	}
	fmt.Println("Keahlian Anda saat ini:")
	for i, keahlian := range keahlianList {
		fmt.Printf("%d. %s\n", i+1, keahlian)
	}
	fmt.Println("\n pilih menu dibawah ini:")
	fmt.Println("1. Tambah keahlian")
	fmt.Println("2. Hapus keahlian (pilih 1)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilihan kamu: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahKeahlian()
	case 2:
		hapusKeahlian()
	case 3:
		return
	default:
		fmt.Println("Pilihan Tidak Valid")
	}
}

func hapusMinat() {
	fmt.Println("\n=== Hapus Minat ===")
	var pilihan int

	fmt.Print("\nPilih nomor minat yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(minatList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	minatTerhapus := minatList[pilihan-1]
	minatList = append(minatList[:pilihan-1], minatList[pilihan:]...)
	fmt.Println("Minat berhasil dihapus\n", minatTerhapus, 3-len(minatList))
}

func hapusKeahlian() {
	var pilihan int

	fmt.Print("\nPilih nomor keahlian yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(keahlianList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	keahlianTerhapus := keahlianList[pilihan-1]
	keahlianList = append(keahlianList[:pilihan-1], keahlianList[pilihan:]...)
	fmt.Printf("Keahlian '%s' berhasil dihapus\n", keahlianTerhapus)
}
func karirRekomendasi() {
	var rekomendasi []Karir
	var presentase int
	fmt.Println("\n=== Rekomendasi Karir ===")

	if len(minatList) == 0 && len(keahlianList) == 0 {
		fmt.Println("Belum ada data minat atau keahlian")
		return
	}

	fmt.Println("Berdasarkan minat dan keahlian Kamu:")
	fmt.Println("- Minat:", minatList)
	fmt.Println("- Keahlian:", keahlianList)

	for _, minat := range minatList {
		for _, keahlian := range keahlianList {
			if (minat == "Menggambar" || minat == "Melukis") && keahlian == "Design Grafis" {
				rekomendasi = append(rekomendasi, Karir{"UI/UX Designer", 90, "Teknologi", "6-10 juta"})
			} else if minat == "Membaca" && keahlian == "Coding" {
				rekomendasi = append(rekomendasi, Karir{"Software Engineer", 85, "Teknologi", "7-15 juta"})
			} else if minat == "Mengamati" || keahlian == "Menghitung" || keahlian == "Analisis" {
				rekomendasi = append(rekomendasi, Karir{"Data Analyst", 80, "Data", "6-12 juta"})
			} else if minat == "Fotografi" && keahlian == "Editing Video" {
				rekomendasi = append(rekomendasi, Karir{"Video Editor", 75, "Kreatif", "4-8 juta"})
			} else if (minat == "Menggambar") && (keahlian == "Coding" || keahlian == "Design Grafis") {
				rekomendasi = append(rekomendasi, Karir{"Front-End Developer", 85, "Teknologi", "6-10 juta"})
			} else if minat == "Menulis" && (keahlian == "Menghafal" || keahlian == "Mengamati") {
				rekomendasi = append(rekomendasi, Karir{"Technical Writer", 70, "Penulisan", "6-12 juta"})
			} else if minat == "Bermain alat musik" && keahlian == "Menghafal" {
				rekomendasi = append(rekomendasi, Karir{"Music Content Editor", 65, "Musik", "5-9 juta"})
			} else if minat == "Menulis" && (keahlian == "Mengamati" || keahlian == "Coding") {
				rekomendasi = append(rekomendasi, Karir{"SEO Specialist", 75, "Digital Marketing", "5-8 juta"})
			} else if minat == "Melukis" && keahlian == "Design Grafis" {
				rekomendasi = append(rekomendasi, Karir{"Ilustrator Digital", 80, "Kreatif", "5-9 juta"})
			} else if minat == "Membaca" && (keahlian == "Menulis" || keahlian == "Menggambar") {
				rekomendasi = append(rekomendasi, Karir{"Content Strategist", 70, "Marketing", "6-11 juta"})
			}
		}
	}
	if len(rekomendasi) > 0 {
		fmt.Println("\nIni adalah rekomendasi karir yang cocok:")
		fmt.Println("+-------------------------+--------------+---------------+--------------+")
		fmt.Println("|      PEKERJAAN          | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
		fmt.Println("+-------------------------+--------------+---------------+--------------+")

		for _, karir := range rekomendasi {
			fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
				karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
		}

		fmt.Println("+-------------------------+--------------+---------------+--------------+")
		fmt.Println("Pilih yuk")
		fmt.Println("1. Urut Presentasi kecocokan paling tinggi")
		fmt.Println("2. Urut Presentasi kecocokan paling Rendah")
		fmt.Println("3. Cari Industri")
		fmt.Println("4. Kembali ke menu utama")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&presentase)

		switch presentase {
		case 1:
			for i := 0; i < len(rekomendasi)-1; i++ {
				minIdx := i
				for j := i + 1; j < len(rekomendasi); j++ {
					if rekomendasi[j].Kecocokan > rekomendasi[minIdx].Kecocokan {
						minIdx = j
					}
				}
				rekomendasi[i], rekomendasi[minIdx] = rekomendasi[minIdx], rekomendasi[i]
			}

			fmt.Println("\nUrutan rekomendasi Tertinggi:")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			for i := 0; i < len(rekomendasi); i++ {
				karir := rekomendasi[i]
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
			}
			fmt.Println("+-------------------------+--------------+---------------+--------------+")

		case 2:
			for i := 1; i < len(rekomendasi); i++ {
				key := rekomendasi[i]
				j := i - 1
				for j >= 0 && rekomendasi[j].Kecocokan > key.Kecocokan {
					rekomendasi[j+1] = rekomendasi[j]
					j--
				}
				rekomendasi[j+1] = key
			}

			fmt.Println("\nUrutan rekomendasi Terendah:")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			for i := 0; i < len(rekomendasi); i++ {
				karir := rekomendasi[i]
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
			}
			fmt.Println("+-------------------------+--------------+---------------+--------------+")

		case 3:
			var industriDicari string
			fmt.Print("Masukkan industri yang dicari: ")
			fmt.Scan(&industriDicari)
			binarySearch(rekomendasi, industriDicari)

		case 4:
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	} else {
		fmt.Println("Maaf, belum ada karir yang cocok ditemukan berdasarkan minat dan keahlian kamu.")
	}
}

func tampilMinatDanKeahlian(minatList []string, keahlianList []string) {
	var minat, keahlian string
	fmt.Println("\n=== Daftar Minat & Keahlian ===")
	fmt.Printf("%-4s| %-20s| %-20s\n", "No", "Minat", "Keahlian")
	fmt.Println("----+--------------------+----------------------")

	max := len(minatList)
	if len(keahlianList) > max {
		max = len(keahlianList)
	}

	for i := 0; i < max; i++ {
		if i < len(minatList) {
			minat = minatList[i]
		}
		if i < len(keahlianList) {
			keahlian = keahlianList[i]
		}
		fmt.Printf("%-4d| %-20s| %-20s\n", i+1, minat, keahlian)
	}
}

func pekerjaan() {
	var pilihan, inputAngka int
	pekerjaanData := []Pekerjaan{
		{"UI/UX Designer", "6-10 juta"},
		{"Software Engineer", "7-15 juta"},
		{"Data Analyst", "6-12 juta"},
		{"Video Editor", "4-8 juta"},
		{"Front-End Developer", "6-10 juta"},
		{"Technical Writer", "6-12 juta"},
		{"Music Content Editor", "5-9 juta"},
		{"SEO Specialist", "5-8 juta"},
		{"Illustrator Digital", "5-9 juta"},
		{"Content Strategist", "6-11 juta"},
	}

	fmt.Println("\n+-------------------------+--------------+")
	fmt.Println("|      PEKERJAAN         |     GAJI     |")
	fmt.Println("+-------------------------+--------------+")
	for i := 0; i < len(pekerjaanData); i++ {
		fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
	}
	fmt.Println("+-------------------------+--------------+")
	fmt.Println("\nPilih urutan:")
	fmt.Println("1. Gaji Terbesar ke Terkecil")
	fmt.Println("2. Gaji Terkecil ke Besar")
	fmt.Println("3. Cari pekerjaan berdasarkan gaji")
	fmt.Println("4. Kembali ke Menu Utama")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		for i := 0; i < len(pekerjaanData)-1; i++ {
			maxIdx := i
			for j := i + 1; j < len(pekerjaanData); j++ {
				if pekerjaanData[j].Gaji > pekerjaanData[maxIdx].Gaji {
					maxIdx = j
				}
			}
			pekerjaanData[i], pekerjaanData[maxIdx] = pekerjaanData[maxIdx], pekerjaanData[i]
		}

		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")
		for i := 0; i < len(pekerjaanData); i++ {
			fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
		}
		fmt.Println("+-------------------------+--------------+")

	case 2:
		for i := 1; i < len(pekerjaanData); i++ {
			key := pekerjaanData[i]
			j := i - 1
			for j >= 0 && pekerjaanData[j].Gaji > key.Gaji {
				pekerjaanData[j+1] = pekerjaanData[j]
				j--
			}
			pekerjaanData[j+1] = key
		}

		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")
		for i := 0; i < len(pekerjaanData); i++ {
			fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
		}
		fmt.Println("+-------------------------+--------------+")

	case 3:

		fmt.Print("\nMasukkan angka awal gaji (contoh: 6 untuk '6-10 juta'): ````````````````````````````````                                                                                                `")
		fmt.Scan(&inputAngka)

		found := false
		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")

		for i := 0; i < len(pekerjaanData); i++ {
			gaji := pekerjaanData[i].Gaji
			angkaAwal := int(gaji[0] - '0')

			if angkaAwal == inputAngka {
				fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, gaji)
				found = true
			}
		}

		fmt.Println("+-------------------------+--------------+")
		if !found {
			fmt.Printf("Tidak ada pekerjaan dengan gaji mulai dari %d\n", inputAngka)
		}
	case 4:
		return
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func binarySearch(rekomendasi []Karir, industriDicari string) {
	for i := 0; i < len(rekomendasi)-1; i++ {
		for j := 0; j < len(rekomendasi)-i-1; j++ {
			if rekomendasi[j].Industri > rekomendasi[j+1].Industri {
				rekomendasi[j], rekomendasi[j+1] = rekomendasi[j+1], rekomendasi[j]
			}
		}
	}
	rendah := 0
	tinggi := len(rekomendasi) - 1
	found := false

	fmt.Println("\nHasil pencarian industri", industriDicari, ":")
	fmt.Println("+-------------------------+--------------+---------------+--------------+")
	fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
	fmt.Println("+-------------------------+--------------+---------------+--------------+")

	for rendah <= tinggi {
		tengah := (rendah + tinggi) / 2

		if rekomendasi[tengah].Industri == industriDicari {
			found = true
			fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
				rekomendasi[tengah].Nama, rekomendasi[tengah].Kecocokan,
				rekomendasi[tengah].Industri, rekomendasi[tengah].Gaji)
			kiri := tengah - 1
			for kiri >= 0 && rekomendasi[kiri].Industri == industriDicari {
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					rekomendasi[kiri].Nama, rekomendasi[kiri].Kecocokan,
					rekomendasi[kiri].Industri, rekomendasi[kiri].Gaji)
				kiri--
			}
			kanan := tengah + 1
			for kanan < len(rekomendasi) && rekomendasi[kanan].Industri == industriDicari {
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					rekomendasi[kanan].Nama, rekomendasi[kanan].Kecocokan,
					rekomendasi[kanan].Industri, rekomendasi[kanan].Gaji)
				kanan++
			}
			break
		} else if rekomendasi[tengah].Industri < industriDicari {
			rendah = tengah + 1
		} else {
			tinggi = tengah - 1
		}
	}

	if !found {
		fmt.Println("| Tidak ditemukan karir dengan industri tersebut |")
	}
	fmt.Println("+-------------------------+--------------+---------------+--------------+")
}
package main

import "fmt"

var (
	minatList    = []string{}
	keahlianList = []string{}
)

type Karir struct {
	Nama      string
	Kecocokan int
	Industri  string
	Gaji      string
}

type Pekerjaan struct {
	Nama string
	Gaji string
}

func main() {
	var keluar bool
	var pilihan int

	for !keluar {
		fmt.Println("\nMenu Pilihan:")
		fmt.Println("1. Tambah minat")
		fmt.Println("2. Tambah keahlian")
		fmt.Println("3. Edit minat")
		fmt.Println("4. Edit keahlian")
		fmt.Println("5. Karir rekomendasi")
		fmt.Println("6. Tampilkan Minat Dan Keahlian ")
		fmt.Println("7. Pekerjaan ")
		fmt.Println("8. Keluar")
		fmt.Print("Masukkan pilihan (1-8): ")
		fmt.Scan(&pilihan)
		if pilihan >= 9 {
			fmt.Println("Input tidak valid, silakan coba lagi")
			continue
		}

		switch pilihan {
		case 1:
			tambahMinat()
		case 2:
			tambahKeahlian()
		case 3:
			editMinat()
		case 4:
			editKeahlian()
		case 5:
			karirRekomendasi()
		case 6:
			tampilMinatDanKeahlian(minatList, keahlianList)
		case 7:
			pekerjaan()
		case 8:
			fmt.Println("Terima kasih, program selesai.")
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid, silakan pilih 1-8")
		}
	}
}

func tambahMinat() {
	var input string

	daftarMinat := []string{
		"Membaca",
		"Menulis",
		"Menggambar",
		"Fotografi",
		"Menghitung",
		"Menganalisis",
		"Melukis",
		"Bermain alat musik",
	}

	fmt.Println("\n=== Tambah Minat ===")
	fmt.Println("Daftar minat yang tersedia:")
	for i, minat := range daftarMinat {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	sisa := 3 - len(minatList)
	if sisa <= 0 {
		fmt.Println("\nYahh kamu sudah menambahkan 3 minat")
		return
	}

	fmt.Printf("\nKamu bisa mengisi %d minat (ketik '-' untuk menyudahi tambah minat)\n", sisa)
	for i := 0; i < sisa; i++ {
		fmt.Printf("Pilihan %d (1-8 atau - untuk berhenti): ", i+1)
		fmt.Scan(&input)

		if input == "-" {
			fmt.Println("Kamu Menghentikan Penambahan.")
			break
		}

		var pilihan int
		_, err := fmt.Sscan(input, &pilihan)
		if err != nil || pilihan < 1 || pilihan  > len(daftarMinat) {
			fmt.Println("Input tidak valid")
			return
		}

		duplikat := false
		for _, m := range minatList {
			if m == daftarMinat[pilihan-1] {
				duplikat = true
				break
			}
		}
		if duplikat {
			fmt.Println("Minat sudah ada, coba yang lain.")
			i--
			continue
		}

		minatList = append(minatList, daftarMinat[pilihan-1])
		fmt.Println("--Minat berhasil ditambahkan--")
	}

	fmt.Println("Minat kamu sekarang:", minatList)
}

func tambahKeahlian() {
	var pilihan, sisa int
	sisa = 1 - len(keahlianList)

	fmt.Println("\n=== Tambah Keahlian ===")
	daftarKeahlian := []string{
		"Coding",
		"Analisis",
		"Design Grafis",
		"Editing Video",
		"Menghafal",
		"Mengamati",
		"Melukis",
	}

	fmt.Println("Daftar Keahlian yang tersedia:")
	for i, ahli := range daftarKeahlian {
		fmt.Printf("%d. %s\n", i+1, ahli)
	}

	if sisa <= 0 {
		fmt.Println("\nYahh kamu sudah menambahkan 1 keahlian")
		return
	}

	fmt.Printf("\nKamu bisa mengisi %d keahlian\n", sisa)
	for i := 0; i < sisa; i++ {
		fmt.Printf("Pilihan %d (1-7): ", i+1)
		fmt.Scan(&pilihan)

		if pilihan < 1 || pilihan > 7 {
			fmt.Println("Input tidak valid")
			return
		}

		keahlianList = append(keahlianList, daftarKeahlian[pilihan-1])
		fmt.Println("--Keahlian berhasil ditambahkan--")
	}
}

func editMinat() {
	var pilihan int

	fmt.Println("\n=== Edit Minat ===")
	if len(minatList) == 0 {
		fmt.Println("Belum ada minat yang tersedia")
		return
	}
	fmt.Println("Minat Anda saat ini:")
	for i, minat := range minatList {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	fmt.Println("\n pilih menu dibawah ini:")
	fmt.Println("1. Tambah Minat")
	fmt.Println("2. Hapus Minat (pilih 1)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilihan kamu: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahMinat()
	case 2:
		hapusMinat()
	case 3:
		return
	default:
		fmt.Println("Pilihan Tidak Valid")
	}
}

func editKeahlian() {
	var pilihan int

	fmt.Println("\n=== Edit Keahlian ===")
	if len(keahlianList) == 0 {
		fmt.Println("Belum ada Keahlian yang tersedia")
		return
	}
	fmt.Println("Keahlian Anda saat ini:")
	for i, keahlian := range keahlianList {
		fmt.Printf("%d. %s\n", i+1, keahlian)
	}
	fmt.Println("\n pilih menu dibawah ini:")
	fmt.Println("1. Tambah keahlian")
	fmt.Println("2. Hapus keahlian (pilih 1)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilihan kamu: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahKeahlian()
	case 2:
		hapusKeahlian()
	case 3:
		return
	default:
		fmt.Println("Pilihan Tidak Valid")
	}
}

func hapusMinat() {
	fmt.Println("\n=== Hapus Minat ===")
	var pilihan int

	fmt.Print("\nPilih nomor minat yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(minatList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	minatTerhapus := minatList[pilihan-1]
	minatList = append(minatList[:pilihan-1], minatList[pilihan:]...)
	fmt.Println("Minat berhasil dihapus\n", minatTerhapus, 3-len(minatList))
}

func hapusKeahlian() {
	var pilihan int

	fmt.Print("\nPilih nomor keahlian yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(keahlianList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	keahlianTerhapus := keahlianList[pilihan-1]
	keahlianList = append(keahlianList[:pilihan-1], keahlianList[pilihan:]...)
	fmt.Printf("Keahlian '%s' berhasil dihapus\n", keahlianTerhapus)
}
func karirRekomendasi() {
	var rekomendasi []Karir
	var presentase int
	fmt.Println("\n=== Rekomendasi Karir ===")

	if len(minatList) == 0 && len(keahlianList) == 0 {
		fmt.Println("Belum ada data minat atau keahlian")
		return
	}

	fmt.Println("Berdasarkan minat dan keahlian Kamu:")
	fmt.Println("- Minat:", minatList)
	fmt.Println("- Keahlian:", keahlianList)

	for _, minat := range minatList {
		for _, keahlian := range keahlianList {
			if (minat == "Menggambar" || minat == "Melukis") && keahlian == "Design Grafis" {
				rekomendasi = append(rekomendasi, Karir{"UI/UX Designer", 90, "Teknologi", "6-10 juta"})
			} else if minat == "Membaca" && keahlian == "Coding" {
				rekomendasi = append(rekomendasi, Karir{"Software Engineer", 85, "Teknologi", "7-15 juta"})
			} else if minat == "Mengamati" || keahlian == "Menghitung" || keahlian == "Analisis" {
				rekomendasi = append(rekomendasi, Karir{"Data Analyst", 80, "Data", "6-12 juta"})
			} else if minat == "Fotografi" && keahlian == "Editing Video" {
				rekomendasi = append(rekomendasi, Karir{"Video Editor", 75, "Kreatif", "4-8 juta"})
			} else if (minat == "Menggambar") && (keahlian == "Coding" || keahlian == "Design Grafis") {
				rekomendasi = append(rekomendasi, Karir{"Front-End Developer", 85, "Teknologi", "6-10 juta"})
			} else if minat == "Menulis" && (keahlian == "Menghafal" || keahlian == "Mengamati") {
				rekomendasi = append(rekomendasi, Karir{"Technical Writer", 70, "Penulisan", "6-12 juta"})
			} else if minat == "Bermain alat musik" && keahlian == "Menghafal" {
				rekomendasi = append(rekomendasi, Karir{"Music Content Editor", 65, "Musik", "5-9 juta"})
			} else if minat == "Menulis" && (keahlian == "Mengamati" || keahlian == "Coding") {
				rekomendasi = append(rekomendasi, Karir{"SEO Specialist", 75, "Digital Marketing", "5-8 juta"})
			} else if minat == "Melukis" && keahlian == "Design Grafis" {
				rekomendasi = append(rekomendasi, Karir{"Ilustrator Digital", 80, "Kreatif", "5-9 juta"})
			} else if minat == "Membaca" && (keahlian == "Menulis" || keahlian == "Menggambar") {
				rekomendasi = append(rekomendasi, Karir{"Content Strategist", 70, "Marketing", "6-11 juta"})
			}
		}
	}
	if len(rekomendasi) > 0 {
		fmt.Println("\nIni adalah rekomendasi karir yang cocok:")
		fmt.Println("+-------------------------+--------------+---------------+--------------+")
		fmt.Println("|      PEKERJAAN          | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
		fmt.Println("+-------------------------+--------------+---------------+--------------+")

		for _, karir := range rekomendasi {
			fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
				karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
		}

		fmt.Println("+-------------------------+--------------+---------------+--------------+")
		fmt.Println("Pilih yuk")
		fmt.Println("1. Urut Presentasi kecocokan paling tinggi")
		fmt.Println("2. Urut Presentasi kecocokan paling Rendah")
		fmt.Println("3. Cari Industri")
		fmt.Println("4. Kembali ke menu utama")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&presentase)

		switch presentase {
		case 1:
			for i := 0; i < len(rekomendasi)-1; i++ {
				minIdx := i
				for j := i + 1; j < len(rekomendasi); j++ {
					if rekomendasi[j].Kecocokan > rekomendasi[minIdx].Kecocokan {
						minIdx = j
					}
				}
				rekomendasi[i], rekomendasi[minIdx] = rekomendasi[minIdx], rekomendasi[i]
			}

			fmt.Println("\nUrutan rekomendasi Tertinggi:")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			for i := 0; i < len(rekomendasi); i++ {
				karir := rekomendasi[i]
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
			}
			fmt.Println("+-------------------------+--------------+---------------+--------------+")

		case 2:
			for i := 1; i < len(rekomendasi); i++ {
				key := rekomendasi[i]
				j := i - 1
				for j >= 0 && rekomendasi[j].Kecocokan > key.Kecocokan {
					rekomendasi[j+1] = rekomendasi[j]
					j--
				}
				rekomendasi[j+1] = key
			}

			fmt.Println("\nUrutan rekomendasi Terendah:")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			for i := 0; i < len(rekomendasi); i++ {
				karir := rekomendasi[i]
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
			}
			fmt.Println("+-------------------------+--------------+---------------+--------------+")

		case 3:
			var industriDicari string
			fmt.Print("Masukkan industri yang dicari: ")
			fmt.Scan(&industriDicari)
			binarySearch(rekomendasi, industriDicari)

		case 4:
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	} else {
		fmt.Println("Maaf, belum ada karir yang cocok ditemukan berdasarkan minat dan keahlian kamu.")
	}
}

func tampilMinatDanKeahlian(minatList []string, keahlianList []string) {
	var minat, keahlian string
	fmt.Println("\n=== Daftar Minat & Keahlian ===")
	fmt.Printf("%-4s| %-20s| %-20s\n", "No", "Minat", "Keahlian")
	fmt.Println("----+--------------------+----------------------")

	max := len(minatList)
	if len(keahlianList) > max {
		max = len(keahlianList)
	}

	for i := 0; i < max; i++ {
		if i < len(minatList) {
			minat = minatList[i]
		}
		if i < len(keahlianList) {
			keahlian = keahlianList[i]
		}
		fmt.Printf("%-4d| %-20s| %-20s\n", i+1, minat, keahlian)
	}
}

func pekerjaan() {
	var pilihan, inputAngka int
	pekerjaanData := []Pekerjaan{
		{"UI/UX Designer", "6-10 juta"},
		{"Software Engineer", "7-15 juta"},
		{"Data Analyst", "6-12 juta"},
		{"Video Editor", "4-8 juta"},
		{"Front-End Developer", "6-10 juta"},
		{"Technical Writer", "6-12 juta"},
		{"Music Content Editor", "5-9 juta"},
		{"SEO Specialist", "5-8 juta"},
		{"Illustrator Digital", "5-9 juta"},
		{"Content Strategist", "6-11 juta"},
	}

	fmt.Println("\n+-------------------------+--------------+")
	fmt.Println("|      PEKERJAAN         |     GAJI     |")
	fmt.Println("+-------------------------+--------------+")
	for i := 0; i < len(pekerjaanData); i++ {
		fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
	}
	fmt.Println("+-------------------------+--------------+")
	fmt.Println("\nPilih urutan:")
	fmt.Println("1. Gaji Terbesar ke Terkecil")
	fmt.Println("2. Gaji Terkecil ke Besar")
	fmt.Println("3. Cari pekerjaan berdasarkan gaji")
	fmt.Println("4. Kembali ke Menu Utama")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		for i := 0; i < len(pekerjaanData)-1; i++ {
			maxIdx := i
			for j := i + 1; j < len(pekerjaanData); j++ {
				if pekerjaanData[j].Gaji > pekerjaanData[maxIdx].Gaji {
					maxIdx = j
				}
			}
			pekerjaanData[i], pekerjaanData[maxIdx] = pekerjaanData[maxIdx], pekerjaanData[i]
		}

		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")
		for i := 0; i < len(pekerjaanData); i++ {
			fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
		}
		fmt.Println("+-------------------------+--------------+")

	case 2:
		for i := 1; i < len(pekerjaanData); i++ {
			key := pekerjaanData[i]
			j := i - 1
			for j >= 0 && pekerjaanData[j].Gaji > key.Gaji {
				pekerjaanData[j+1] = pekerjaanData[j]
				j--
			}
			pekerjaanData[j+1] = key
		}

		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")
		for i := 0; i < len(pekerjaanData); i++ {
			fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
		}
		fmt.Println("+-------------------------+--------------+")

	case 3:

		fmt.Print("\nMasukkan angka awal gaji (contoh: 6 untuk '6-10 juta'): ````````````````````````````````                                                                                                `")
		fmt.Scan(&inputAngka)

		found := false
		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")

		for i := 0; i < len(pekerjaanData); i++ {
			gaji := pekerjaanData[i].Gaji
			angkaAwal := int(gaji[0] - '0')

			if angkaAwal == inputAngka {
				fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, gaji)
				found = true
			}
		}

		fmt.Println("+-------------------------+--------------+")
		if !found {
			fmt.Printf("Tidak ada pekerjaan dengan gaji mulai dari %d\n", inputAngka)
		}
	case 4:
		return
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func binarySearch(rekomendasi []Karir, industriDicari string) {
	for i := 0; i < len(rekomendasi)-1; i++ {
		for j := 0; j < len(rekomendasi)-i-1; j++ {
			if rekomendasi[j].Industri > rekomendasi[j+1].Industri {
				rekomendasi[j], rekomendasi[j+1] = rekomendasi[j+1], rekomendasi[j]
			}
		}
	}
	rendah := 0
	tinggi := len(rekomendasi) - 1
	found := false

	fmt.Println("\nHasil pencarian industri", industriDicari, ":")
	fmt.Println("+-------------------------+--------------+---------------+--------------+")
	fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
	fmt.Println("+-------------------------+--------------+---------------+--------------+")

	for rendah <= tinggi {
		tengah := (rendah + tinggi) / 2

		if rekomendasi[tengah].Industri == industriDicari {
			found = true
			fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
				rekomendasi[tengah].Nama, rekomendasi[tengah].Kecocokan,
				rekomendasi[tengah].Industri, rekomendasi[tengah].Gaji)
			kiri := tengah - 1
			for kiri >= 0 && rekomendasi[kiri].Industri == industriDicari {
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					rekomendasi[kiri].Nama, rekomendasi[kiri].Kecocokan,
					rekomendasi[kiri].Industri, rekomendasi[kiri].Gaji)
				kiri--
			}
			kanan := tengah + 1
			for kanan < len(rekomendasi) && rekomendasi[kanan].Industri == industriDicari {
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					rekomendasi[kanan].Nama, rekomendasi[kanan].Kecocokan,
					rekomendasi[kanan].Industri, rekomendasi[kanan].Gaji)
				kanan++
			}
			break
		} else if rekomendasi[tengah].Industri < industriDicari {
			rendah = tengah + 1
		} else {
			tinggi = tengah - 1
		}
	}

	if !found {
		fmt.Println("| Tidak ditemukan karir dengan industri tersebut |")
	}
	fmt.Println("+-------------------------+--------------+---------------+--------------+")
}
package main

import "fmt"

var (
	minatList    = []string{}
	keahlianList = []string{}
)

type Karir struct {
	Nama      string
	Kecocokan int
	Industri  string
	Gaji      string
}

type Pekerjaan struct {
	Nama string
	Gaji string
}

func main() {
	var keluar bool
	var pilihan int

	for !keluar {
		fmt.Println("\nMenu Pilihan:")
		fmt.Println("1. Tambah minat")
		fmt.Println("2. Tambah keahlian")
		fmt.Println("3. Edit minat")
		fmt.Println("4. Edit keahlian")
		fmt.Println("5. Karir rekomendasi")
		fmt.Println("6. Tampilkan Minat Dan Keahlian ")
		fmt.Println("7. Pekerjaan ")
		fmt.Println("8. Keluar")
		fmt.Print("Masukkan pilihan (1-8): ")
		fmt.Scan(&pilihan)
		if pilihan >= 9 {
			fmt.Println("Input tidak valid, silakan coba lagi")
			continue
		}

		switch pilihan {
		case 1:
			tambahMinat()
		case 2:
			tambahKeahlian()
		case 3:
			editMinat()
		case 4:
			editKeahlian()
		case 5:
			karirRekomendasi()
		case 6:
			tampilMinatDanKeahlian(minatList, keahlianList)
		case 7:
			pekerjaan()
		case 8:
			fmt.Println("Terima kasih, program selesai.")
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid, silakan pilih 1-8")
		}
	}
}

func tambahMinat() {
	var input string

	daftarMinat := []string{
		"Membaca",
		"Menulis",
		"Menggambar",
		"Fotografi",
		"Menghitung",
		"Menganalisis",
		"Melukis",
		"Bermain alat musik",
	}

	fmt.Println("\n=== Tambah Minat ===")
	fmt.Println("Daftar minat yang tersedia:")
	for i, minat := range daftarMinat {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	sisa := 3 - len(minatList)
	if sisa <= 0 {
		fmt.Println("\nYahh kamu sudah menambahkan 3 minat")
		return
	}

	fmt.Printf("\nKamu bisa mengisi %d minat (ketik '-' untuk menyudahi tambah minat)\n", sisa)
	for i := 0; i < sisa; i++ {
		fmt.Printf("Pilihan %d (1-8 atau - untuk berhenti): ", i+1)
		fmt.Scan(&input)

		if input == "-" {
			fmt.Println("Kamu Menghentikan Penambahan.")
			break
		}

		var pilihan int
		_, err := fmt.Sscan(input, &pilihan)
		if err != nil || pilihan < 1 || pilihan  > len(daftarMinat) {
			fmt.Println("Input tidak valid")
			return
		}

		duplikat := false
		for _, m := range minatList {
			if m == daftarMinat[pilihan-1] {
				duplikat = true
				break
			}
		}
		if duplikat {
			fmt.Println("Minat sudah ada, coba yang lain.")
			i--
			continue
		}

		minatList = append(minatList, daftarMinat[pilihan-1])
		fmt.Println("--Minat berhasil ditambahkan--")
	}

	fmt.Println("Minat kamu sekarang:", minatList)
}

func tambahKeahlian() {
	var pilihan, sisa int
	sisa = 1 - len(keahlianList)

	fmt.Println("\n=== Tambah Keahlian ===")
	daftarKeahlian := []string{
		"Coding",
		"Analisis",
		"Design Grafis",
		"Editing Video",
		"Menghafal",
		"Mengamati",
		"Melukis",
	}

	fmt.Println("Daftar Keahlian yang tersedia:")
	for i, ahli := range daftarKeahlian {
		fmt.Printf("%d. %s\n", i+1, ahli)
	}

	if sisa <= 0 {
		fmt.Println("\nYahh kamu sudah menambahkan 1 keahlian")
		return
	}

	fmt.Printf("\nKamu bisa mengisi %d keahlian\n", sisa)
	for i := 0; i < sisa; i++ {
		fmt.Printf("Pilihan %d (1-7): ", i+1)
		fmt.Scan(&pilihan)

		if pilihan < 1 || pilihan > 7 {
			fmt.Println("Input tidak valid")
			return
		}

		keahlianList = append(keahlianList, daftarKeahlian[pilihan-1])
		fmt.Println("--Keahlian berhasil ditambahkan--")
	}
}

func editMinat() {
	var pilihan int

	fmt.Println("\n=== Edit Minat ===")
	if len(minatList) == 0 {
		fmt.Println("Belum ada minat yang tersedia")
		return
	}
	fmt.Println("Minat Anda saat ini:")
	for i, minat := range minatList {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	fmt.Println("\n pilih menu dibawah ini:")
	fmt.Println("1. Tambah Minat")
	fmt.Println("2. Hapus Minat (pilih 1)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilihan kamu: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahMinat()
	case 2:
		hapusMinat()
	case 3:
		return
	default:
		fmt.Println("Pilihan Tidak Valid")
	}
}

func editKeahlian() {
	var pilihan int

	fmt.Println("\n=== Edit Keahlian ===")
	if len(keahlianList) == 0 {
		fmt.Println("Belum ada Keahlian yang tersedia")
		return
	}
	fmt.Println("Keahlian Anda saat ini:")
	for i, keahlian := range keahlianList {
		fmt.Printf("%d. %s\n", i+1, keahlian)
	}
	fmt.Println("\n pilih menu dibawah ini:")
	fmt.Println("1. Tambah keahlian")
	fmt.Println("2. Hapus keahlian (pilih 1)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilihan kamu: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahKeahlian()
	case 2:
		hapusKeahlian()
	case 3:
		return
	default:
		fmt.Println("Pilihan Tidak Valid")
	}
}

func hapusMinat() {
	fmt.Println("\n=== Hapus Minat ===")
	var pilihan int

	fmt.Print("\nPilih nomor minat yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(minatList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	minatTerhapus := minatList[pilihan-1]
	minatList = append(minatList[:pilihan-1], minatList[pilihan:]...)
	fmt.Println("Minat berhasil dihapus\n", minatTerhapus, 3-len(minatList))
}

func hapusKeahlian() {
	var pilihan int

	fmt.Print("\nPilih nomor keahlian yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(keahlianList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	keahlianTerhapus := keahlianList[pilihan-1]
	keahlianList = append(keahlianList[:pilihan-1], keahlianList[pilihan:]...)
	fmt.Printf("Keahlian '%s' berhasil dihapus\n", keahlianTerhapus)
}
func karirRekomendasi() {
	var rekomendasi []Karir
	var presentase int
	fmt.Println("\n=== Rekomendasi Karir ===")

	if len(minatList) == 0 && len(keahlianList) == 0 {
		fmt.Println("Belum ada data minat atau keahlian")
		return
	}

	fmt.Println("Berdasarkan minat dan keahlian Kamu:")
	fmt.Println("- Minat:", minatList)
	fmt.Println("- Keahlian:", keahlianList)

	for _, minat := range minatList {
		for _, keahlian := range keahlianList {
			if (minat == "Menggambar" || minat == "Melukis") && keahlian == "Design Grafis" {
				rekomendasi = append(rekomendasi, Karir{"UI/UX Designer", 90, "Teknologi", "6-10 juta"})
			} else if minat == "Membaca" && keahlian == "Coding" {
				rekomendasi = append(rekomendasi, Karir{"Software Engineer", 85, "Teknologi", "7-15 juta"})
			} else if minat == "Mengamati" || keahlian == "Menghitung" || keahlian == "Analisis" {
				rekomendasi = append(rekomendasi, Karir{"Data Analyst", 80, "Data", "6-12 juta"})
			} else if minat == "Fotografi" && keahlian == "Editing Video" {
				rekomendasi = append(rekomendasi, Karir{"Video Editor", 75, "Kreatif", "4-8 juta"})
			} else if (minat == "Menggambar") && (keahlian == "Coding" || keahlian == "Design Grafis") {
				rekomendasi = append(rekomendasi, Karir{"Front-End Developer", 85, "Teknologi", "6-10 juta"})
			} else if minat == "Menulis" && (keahlian == "Menghafal" || keahlian == "Mengamati") {
				rekomendasi = append(rekomendasi, Karir{"Technical Writer", 70, "Penulisan", "6-12 juta"})
			} else if minat == "Bermain alat musik" && keahlian == "Menghafal" {
				rekomendasi = append(rekomendasi, Karir{"Music Content Editor", 65, "Musik", "5-9 juta"})
			} else if minat == "Menulis" && (keahlian == "Mengamati" || keahlian == "Coding") {
				rekomendasi = append(rekomendasi, Karir{"SEO Specialist", 75, "Digital Marketing", "5-8 juta"})
			} else if minat == "Melukis" && keahlian == "Design Grafis" {
				rekomendasi = append(rekomendasi, Karir{"Ilustrator Digital", 80, "Kreatif", "5-9 juta"})
			} else if minat == "Membaca" && (keahlian == "Menulis" || keahlian == "Menggambar") {
				rekomendasi = append(rekomendasi, Karir{"Content Strategist", 70, "Marketing", "6-11 juta"})
			}
		}
	}
	if len(rekomendasi) > 0 {
		fmt.Println("\nIni adalah rekomendasi karir yang cocok:")
		fmt.Println("+-------------------------+--------------+---------------+--------------+")
		fmt.Println("|      PEKERJAAN          | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
		fmt.Println("+-------------------------+--------------+---------------+--------------+")

		for _, karir := range rekomendasi {
			fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
				karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
		}

		fmt.Println("+-------------------------+--------------+---------------+--------------+")
		fmt.Println("Pilih yuk")
		fmt.Println("1. Urut Presentasi kecocokan paling tinggi")
		fmt.Println("2. Urut Presentasi kecocokan paling Rendah")
		fmt.Println("3. Cari Industri")
		fmt.Println("4. Kembali ke menu utama")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&presentase)

		switch presentase {
		case 1:
			for i := 0; i < len(rekomendasi)-1; i++ {
				minIdx := i
				for j := i + 1; j < len(rekomendasi); j++ {
					if rekomendasi[j].Kecocokan > rekomendasi[minIdx].Kecocokan {
						minIdx = j
					}
				}
				rekomendasi[i], rekomendasi[minIdx] = rekomendasi[minIdx], rekomendasi[i]
			}

			fmt.Println("\nUrutan rekomendasi Tertinggi:")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			for i := 0; i < len(rekomendasi); i++ {
				karir := rekomendasi[i]
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
			}
			fmt.Println("+-------------------------+--------------+---------------+--------------+")

		case 2:
			for i := 1; i < len(rekomendasi); i++ {
				key := rekomendasi[i]
				j := i - 1
				for j >= 0 && rekomendasi[j].Kecocokan > key.Kecocokan {
					rekomendasi[j+1] = rekomendasi[j]
					j--
				}
				rekomendasi[j+1] = key
			}

			fmt.Println("\nUrutan rekomendasi Terendah:")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			for i := 0; i < len(rekomendasi); i++ {
				karir := rekomendasi[i]
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
			}
			fmt.Println("+-------------------------+--------------+---------------+--------------+")

		case 3:
			var industriDicari string
			fmt.Print("Masukkan industri yang dicari: ")
			fmt.Scan(&industriDicari)
			binarySearch(rekomendasi, industriDicari)

		case 4:
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	} else {
		fmt.Println("Maaf, belum ada karir yang cocok ditemukan berdasarkan minat dan keahlian kamu.")
	}
}

func tampilMinatDanKeahlian(minatList []string, keahlianList []string) {
	var minat, keahlian string
	fmt.Println("\n=== Daftar Minat & Keahlian ===")
	fmt.Printf("%-4s| %-20s| %-20s\n", "No", "Minat", "Keahlian")
	fmt.Println("----+--------------------+----------------------")

	max := len(minatList)
	if len(keahlianList) > max {
		max = len(keahlianList)
	}

	for i := 0; i < max; i++ {
		if i < len(minatList) {
			minat = minatList[i]
		}
		if i < len(keahlianList) {
			keahlian = keahlianList[i]
		}
		fmt.Printf("%-4d| %-20s| %-20s\n", i+1, minat, keahlian)
	}
}

func pekerjaan() {
	var pilihan, inputAngka int
	pekerjaanData := []Pekerjaan{
		{"UI/UX Designer", "6-10 juta"},
		{"Software Engineer", "7-15 juta"},
		{"Data Analyst", "6-12 juta"},
		{"Video Editor", "4-8 juta"},
		{"Front-End Developer", "6-10 juta"},
		{"Technical Writer", "6-12 juta"},
		{"Music Content Editor", "5-9 juta"},
		{"SEO Specialist", "5-8 juta"},
		{"Illustrator Digital", "5-9 juta"},
		{"Content Strategist", "6-11 juta"},
	}

	fmt.Println("\n+-------------------------+--------------+")
	fmt.Println("|      PEKERJAAN         |     GAJI     |")
	fmt.Println("+-------------------------+--------------+")
	for i := 0; i < len(pekerjaanData); i++ {
		fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
	}
	fmt.Println("+-------------------------+--------------+")
	fmt.Println("\nPilih urutan:")
	fmt.Println("1. Gaji Terbesar ke Terkecil")
	fmt.Println("2. Gaji Terkecil ke Besar")
	fmt.Println("3. Cari pekerjaan berdasarkan gaji")
	fmt.Println("4. Kembali ke Menu Utama")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		for i := 0; i < len(pekerjaanData)-1; i++ {
			maxIdx := i
			for j := i + 1; j < len(pekerjaanData); j++ {
				if pekerjaanData[j].Gaji > pekerjaanData[maxIdx].Gaji {
					maxIdx = j
				}
			}
			pekerjaanData[i], pekerjaanData[maxIdx] = pekerjaanData[maxIdx], pekerjaanData[i]
		}

		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")
		for i := 0; i < len(pekerjaanData); i++ {
			fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
		}
		fmt.Println("+-------------------------+--------------+")

	case 2:
		for i := 1; i < len(pekerjaanData); i++ {
			key := pekerjaanData[i]
			j := i - 1
			for j >= 0 && pekerjaanData[j].Gaji > key.Gaji {
				pekerjaanData[j+1] = pekerjaanData[j]
				j--
			}
			pekerjaanData[j+1] = key
		}

		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")
		for i := 0; i < len(pekerjaanData); i++ {
			fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
		}
		fmt.Println("+-------------------------+--------------+")

	case 3:

		fmt.Print("\nMasukkan angka awal gaji (contoh: 6 untuk '6-10 juta'): ````````````````````````````````                                                                                                `")
		fmt.Scan(&inputAngka)

		found := false
		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")

		for i := 0; i < len(pekerjaanData); i++ {
			gaji := pekerjaanData[i].Gaji
			angkaAwal := int(gaji[0] - '0')

			if angkaAwal == inputAngka {
				fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, gaji)
				found = true
			}
		}

		fmt.Println("+-------------------------+--------------+")
		if !found {
			fmt.Printf("Tidak ada pekerjaan dengan gaji mulai dari %d\n", inputAngka)
		}
	case 4:
		return
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func binarySearch(rekomendasi []Karir, industriDicari string) {
	for i := 0; i < len(rekomendasi)-1; i++ {
		for j := 0; j < len(rekomendasi)-i-1; j++ {
			if rekomendasi[j].Industri > rekomendasi[j+1].Industri {
				rekomendasi[j], rekomendasi[j+1] = rekomendasi[j+1], rekomendasi[j]
			}
		}
	}
	rendah := 0
	tinggi := len(rekomendasi) - 1
	found := false

	fmt.Println("\nHasil pencarian industri", industriDicari, ":")
	fmt.Println("+-------------------------+--------------+---------------+--------------+")
	fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
	fmt.Println("+-------------------------+--------------+---------------+--------------+")

	for rendah <= tinggi {
		tengah := (rendah + tinggi) / 2

		if rekomendasi[tengah].Industri == industriDicari {
			found = true
			fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
				rekomendasi[tengah].Nama, rekomendasi[tengah].Kecocokan,
				rekomendasi[tengah].Industri, rekomendasi[tengah].Gaji)
			kiri := tengah - 1
			for kiri >= 0 && rekomendasi[kiri].Industri == industriDicari {
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					rekomendasi[kiri].Nama, rekomendasi[kiri].Kecocokan,
					rekomendasi[kiri].Industri, rekomendasi[kiri].Gaji)
				kiri--
			}
			kanan := tengah + 1
			for kanan < len(rekomendasi) && rekomendasi[kanan].Industri == industriDicari {
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					rekomendasi[kanan].Nama, rekomendasi[kanan].Kecocokan,
					rekomendasi[kanan].Industri, rekomendasi[kanan].Gaji)
				kanan++
			}
			break
		} else if rekomendasi[tengah].Industri < industriDicari {
			rendah = tengah + 1
		} else {
			tinggi = tengah - 1
		}
	}

	if !found {
		fmt.Println("| Tidak ditemukan karir dengan industri tersebut |")
	}
	fmt.Println("+-------------------------+--------------+---------------+--------------+")
}
package main

import "fmt"

var (
	minatList    = []string{}
	keahlianList = []string{}
)

type Karir struct {
	Nama      string
	Kecocokan int
	Industri  string
	Gaji      string
}

type Pekerjaan struct {
	Nama string
	Gaji string
}

func main() {
	var keluar bool
	var pilihan int

	for !keluar {
		fmt.Println("\nMenu Pilihan:")
		fmt.Println("1. Tambah minat")
		fmt.Println("2. Tambah keahlian")
		fmt.Println("3. Edit minat")
		fmt.Println("4. Edit keahlian")
		fmt.Println("5. Karir rekomendasi")
		fmt.Println("6. Tampilkan Minat Dan Keahlian ")
		fmt.Println("7. Pekerjaan ")
		fmt.Println("8. Keluar")
		fmt.Print("Masukkan pilihan (1-8): ")
		fmt.Scan(&pilihan)
		if pilihan >= 9 {
			fmt.Println("Input tidak valid, silakan coba lagi")
			continue
		}

		switch pilihan {
		case 1:
			tambahMinat()
		case 2:
			tambahKeahlian()
		case 3:
			editMinat()
		case 4:
			editKeahlian()
		case 5:
			karirRekomendasi()
		case 6:
			tampilMinatDanKeahlian(minatList, keahlianList)
		case 7:
			pekerjaan()
		case 8:
			fmt.Println("Terima kasih, program selesai.")
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid, silakan pilih 1-8")
		}
	}
}

func tambahMinat() {
	var input string

	daftarMinat := []string{
		"Membaca",
		"Menulis",
		"Menggambar",
		"Fotografi",
		"Menghitung",
		"Menganalisis",
		"Melukis",
		"Bermain alat musik",
	}

	fmt.Println("\n=== Tambah Minat ===")
	fmt.Println("Daftar minat yang tersedia:")
	for i, minat := range daftarMinat {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	sisa := 3 - len(minatList)
	if sisa <= 0 {
		fmt.Println("\nYahh kamu sudah menambahkan 3 minat")
		return
	}

	fmt.Printf("\nKamu bisa mengisi %d minat (ketik '-' untuk menyudahi tambah minat)\n", sisa)
	for i := 0; i < sisa; i++ {
		fmt.Printf("Pilihan %d (1-8 atau - untuk berhenti): ", i+1)
		fmt.Scan(&input)

		if input == "-" {
			fmt.Println("Kamu Menghentikan Penambahan.")
			break
		}

		var pilihan int
		_, err := fmt.Sscan(input, &pilihan)
		if err != nil || pilihan < 1 || pilihan  > len(daftarMinat) {
			fmt.Println("Input tidak valid")
			return
		}

		duplikat := false
		for _, m := range minatList {
			if m == daftarMinat[pilihan-1] {
				duplikat = true
				break
			}
		}
		if duplikat {
			fmt.Println("Minat sudah ada, coba yang lain.")
			i--
			continue
		}

		minatList = append(minatList, daftarMinat[pilihan-1])
		fmt.Println("--Minat berhasil ditambahkan--")
	}

	fmt.Println("Minat kamu sekarang:", minatList)
}

func tambahKeahlian() {
	var pilihan, sisa int
	sisa = 1 - len(keahlianList)

	fmt.Println("\n=== Tambah Keahlian ===")
	daftarKeahlian := []string{
		"Coding",
		"Analisis",
		"Design Grafis",
		"Editing Video",
		"Menghafal",
		"Mengamati",
		"Melukis",
	}

	fmt.Println("Daftar Keahlian yang tersedia:")
	for i, ahli := range daftarKeahlian {
		fmt.Printf("%d. %s\n", i+1, ahli)
	}

	if sisa <= 0 {
		fmt.Println("\nYahh kamu sudah menambahkan 1 keahlian")
		return
	}

	fmt.Printf("\nKamu bisa mengisi %d keahlian\n", sisa)
	for i := 0; i < sisa; i++ {
		fmt.Printf("Pilihan %d (1-7): ", i+1)
		fmt.Scan(&pilihan)

		if pilihan < 1 || pilihan > 7 {
			fmt.Println("Input tidak valid")
			return
		}

		keahlianList = append(keahlianList, daftarKeahlian[pilihan-1])
		fmt.Println("--Keahlian berhasil ditambahkan--")
	}
}

func editMinat() {
	var pilihan int

	fmt.Println("\n=== Edit Minat ===")
	if len(minatList) == 0 {
		fmt.Println("Belum ada minat yang tersedia")
		return
	}
	fmt.Println("Minat Anda saat ini:")
	for i, minat := range minatList {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	fmt.Println("\n pilih menu dibawah ini:")
	fmt.Println("1. Tambah Minat")
	fmt.Println("2. Hapus Minat (pilih 1)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilihan kamu: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahMinat()
	case 2:
		hapusMinat()
	case 3:
		return
	default:
		fmt.Println("Pilihan Tidak Valid")
	}
}

func editKeahlian() {
	var pilihan int

	fmt.Println("\n=== Edit Keahlian ===")
	if len(keahlianList) == 0 {
		fmt.Println("Belum ada Keahlian yang tersedia")
		return
	}
	fmt.Println("Keahlian Anda saat ini:")
	for i, keahlian := range keahlianList {
		fmt.Printf("%d. %s\n", i+1, keahlian)
	}
	fmt.Println("\n pilih menu dibawah ini:")
	fmt.Println("1. Tambah keahlian")
	fmt.Println("2. Hapus keahlian (pilih 1)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilihan kamu: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahKeahlian()
	case 2:
		hapusKeahlian()
	case 3:
		return
	default:
		fmt.Println("Pilihan Tidak Valid")
	}
}

func hapusMinat() {
	fmt.Println("\n=== Hapus Minat ===")
	var pilihan int

	fmt.Print("\nPilih nomor minat yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(minatList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	minatTerhapus := minatList[pilihan-1]
	minatList = append(minatList[:pilihan-1], minatList[pilihan:]...)
	fmt.Println("Minat berhasil dihapus\n", minatTerhapus, 3-len(minatList))
}

func hapusKeahlian() {
	var pilihan int

	fmt.Print("\nPilih nomor keahlian yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(keahlianList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	keahlianTerhapus := keahlianList[pilihan-1]
	keahlianList = append(keahlianList[:pilihan-1], keahlianList[pilihan:]...)
	fmt.Printf("Keahlian '%s' berhasil dihapus\n", keahlianTerhapus)
}
func karirRekomendasi() {
	var rekomendasi []Karir
	var presentase int
	fmt.Println("\n=== Rekomendasi Karir ===")

	if len(minatList) == 0 && len(keahlianList) == 0 {
		fmt.Println("Belum ada data minat atau keahlian")
		return
	}

	fmt.Println("Berdasarkan minat dan keahlian Kamu:")
	fmt.Println("- Minat:", minatList)
	fmt.Println("- Keahlian:", keahlianList)

	for _, minat := range minatList {
		for _, keahlian := range keahlianList {
			if (minat == "Menggambar" || minat == "Melukis") && keahlian == "Design Grafis" {
				rekomendasi = append(rekomendasi, Karir{"UI/UX Designer", 90, "Teknologi", "6-10 juta"})
			} else if minat == "Membaca" && keahlian == "Coding" {
				rekomendasi = append(rekomendasi, Karir{"Software Engineer", 85, "Teknologi", "7-15 juta"})
			} else if minat == "Mengamati" || keahlian == "Menghitung" || keahlian == "Analisis" {
				rekomendasi = append(rekomendasi, Karir{"Data Analyst", 80, "Data", "6-12 juta"})
			} else if minat == "Fotografi" && keahlian == "Editing Video" {
				rekomendasi = append(rekomendasi, Karir{"Video Editor", 75, "Kreatif", "4-8 juta"})
			} else if (minat == "Menggambar") && (keahlian == "Coding" || keahlian == "Design Grafis") {
				rekomendasi = append(rekomendasi, Karir{"Front-End Developer", 85, "Teknologi", "6-10 juta"})
			} else if minat == "Menulis" && (keahlian == "Menghafal" || keahlian == "Mengamati") {
				rekomendasi = append(rekomendasi, Karir{"Technical Writer", 70, "Penulisan", "6-12 juta"})
			} else if minat == "Bermain alat musik" && keahlian == "Menghafal" {
				rekomendasi = append(rekomendasi, Karir{"Music Content Editor", 65, "Musik", "5-9 juta"})
			} else if minat == "Menulis" && (keahlian == "Mengamati" || keahlian == "Coding") {
				rekomendasi = append(rekomendasi, Karir{"SEO Specialist", 75, "Digital Marketing", "5-8 juta"})
			} else if minat == "Melukis" && keahlian == "Design Grafis" {
				rekomendasi = append(rekomendasi, Karir{"Ilustrator Digital", 80, "Kreatif", "5-9 juta"})
			} else if minat == "Membaca" && (keahlian == "Menulis" || keahlian == "Menggambar") {
				rekomendasi = append(rekomendasi, Karir{"Content Strategist", 70, "Marketing", "6-11 juta"})
			}
		}
	}
	if len(rekomendasi) > 0 {
		fmt.Println("\nIni adalah rekomendasi karir yang cocok:")
		fmt.Println("+-------------------------+--------------+---------------+--------------+")
		fmt.Println("|      PEKERJAAN          | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
		fmt.Println("+-------------------------+--------------+---------------+--------------+")

		for _, karir := range rekomendasi {
			fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
				karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
		}

		fmt.Println("+-------------------------+--------------+---------------+--------------+")
		fmt.Println("Pilih yuk")
		fmt.Println("1. Urut Presentasi kecocokan paling tinggi")
		fmt.Println("2. Urut Presentasi kecocokan paling Rendah")
		fmt.Println("3. Cari Industri")
		fmt.Println("4. Kembali ke menu utama")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&presentase)

		switch presentase {
		case 1:
			for i := 0; i < len(rekomendasi)-1; i++ {
				minIdx := i
				for j := i + 1; j < len(rekomendasi); j++ {
					if rekomendasi[j].Kecocokan > rekomendasi[minIdx].Kecocokan {
						minIdx = j
					}
				}
				rekomendasi[i], rekomendasi[minIdx] = rekomendasi[minIdx], rekomendasi[i]
			}

			fmt.Println("\nUrutan rekomendasi Tertinggi:")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			for i := 0; i < len(rekomendasi); i++ {
				karir := rekomendasi[i]
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
			}
			fmt.Println("+-------------------------+--------------+---------------+--------------+")

		case 2:
			for i := 1; i < len(rekomendasi); i++ {
				key := rekomendasi[i]
				j := i - 1
				for j >= 0 && rekomendasi[j].Kecocokan > key.Kecocokan {
					rekomendasi[j+1] = rekomendasi[j]
					j--
				}
				rekomendasi[j+1] = key
			}

			fmt.Println("\nUrutan rekomendasi Terendah:")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			for i := 0; i < len(rekomendasi); i++ {
				karir := rekomendasi[i]
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
			}
			fmt.Println("+-------------------------+--------------+---------------+--------------+")

		case 3:
			var industriDicari string
			fmt.Print("Masukkan industri yang dicari: ")
			fmt.Scan(&industriDicari)
			binarySearch(rekomendasi, industriDicari)

		case 4:
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	} else {
		fmt.Println("Maaf, belum ada karir yang cocok ditemukan berdasarkan minat dan keahlian kamu.")
	}
}

func tampilMinatDanKeahlian(minatList []string, keahlianList []string) {
	var minat, keahlian string
	fmt.Println("\n=== Daftar Minat & Keahlian ===")
	fmt.Printf("%-4s| %-20s| %-20s\n", "No", "Minat", "Keahlian")
	fmt.Println("----+--------------------+----------------------")

	max := len(minatList)
	if len(keahlianList) > max {
		max = len(keahlianList)
	}

	for i := 0; i < max; i++ {
		if i < len(minatList) {
			minat = minatList[i]
		}
		if i < len(keahlianList) {
			keahlian = keahlianList[i]
		}
		fmt.Printf("%-4d| %-20s| %-20s\n", i+1, minat, keahlian)
	}
}

func pekerjaan() {
	var pilihan, inputAngka int
	pekerjaanData := []Pekerjaan{
		{"UI/UX Designer", "6-10 juta"},
		{"Software Engineer", "7-15 juta"},
		{"Data Analyst", "6-12 juta"},
		{"Video Editor", "4-8 juta"},
		{"Front-End Developer", "6-10 juta"},
		{"Technical Writer", "6-12 juta"},
		{"Music Content Editor", "5-9 juta"},
		{"SEO Specialist", "5-8 juta"},
		{"Illustrator Digital", "5-9 juta"},
		{"Content Strategist", "6-11 juta"},
	}

	fmt.Println("\n+-------------------------+--------------+")
	fmt.Println("|      PEKERJAAN         |     GAJI     |")
	fmt.Println("+-------------------------+--------------+")
	for i := 0; i < len(pekerjaanData); i++ {
		fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
	}
	fmt.Println("+-------------------------+--------------+")
	fmt.Println("\nPilih urutan:")
	fmt.Println("1. Gaji Terbesar ke Terkecil")
	fmt.Println("2. Gaji Terkecil ke Besar")
	fmt.Println("3. Cari pekerjaan berdasarkan gaji")
	fmt.Println("4. Kembali ke Menu Utama")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		for i := 0; i < len(pekerjaanData)-1; i++ {
			maxIdx := i
			for j := i + 1; j < len(pekerjaanData); j++ {
				if pekerjaanData[j].Gaji > pekerjaanData[maxIdx].Gaji {
					maxIdx = j
				}
			}
			pekerjaanData[i], pekerjaanData[maxIdx] = pekerjaanData[maxIdx], pekerjaanData[i]
		}

		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")
		for i := 0; i < len(pekerjaanData); i++ {
			fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
		}
		fmt.Println("+-------------------------+--------------+")

	case 2:
		for i := 1; i < len(pekerjaanData); i++ {
			key := pekerjaanData[i]
			j := i - 1
			for j >= 0 && pekerjaanData[j].Gaji > key.Gaji {
				pekerjaanData[j+1] = pekerjaanData[j]
				j--
			}
			pekerjaanData[j+1] = key
		}

		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")
		for i := 0; i < len(pekerjaanData); i++ {
			fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
		}
		fmt.Println("+-------------------------+--------------+")

	case 3:

		fmt.Print("\nMasukkan angka awal gaji (contoh: 6 untuk '6-10 juta'): ````````````````````````````````                                                                                                `")
		fmt.Scan(&inputAngka)

		found := false
		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")

		for i := 0; i < len(pekerjaanData); i++ {
			gaji := pekerjaanData[i].Gaji
			angkaAwal := int(gaji[0] - '0')

			if angkaAwal == inputAngka {
				fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, gaji)
				found = true
			}
		}

		fmt.Println("+-------------------------+--------------+")
		if !found {
			fmt.Printf("Tidak ada pekerjaan dengan gaji mulai dari %d\n", inputAngka)
		}
	case 4:
		return
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func binarySearch(rekomendasi []Karir, industriDicari string) {
	for i := 0; i < len(rekomendasi)-1; i++ {
		for j := 0; j < len(rekomendasi)-i-1; j++ {
			if rekomendasi[j].Industri > rekomendasi[j+1].Industri {
				rekomendasi[j], rekomendasi[j+1] = rekomendasi[j+1], rekomendasi[j]
			}
		}
	}
	rendah := 0
	tinggi := len(rekomendasi) - 1
	found := false

	fmt.Println("\nHasil pencarian industri", industriDicari, ":")
	fmt.Println("+-------------------------+--------------+---------------+--------------+")
	fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
	fmt.Println("+-------------------------+--------------+---------------+--------------+")

	for rendah <= tinggi {
		tengah := (rendah + tinggi) / 2

		if rekomendasi[tengah].Industri == industriDicari {
			found = true
			fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
				rekomendasi[tengah].Nama, rekomendasi[tengah].Kecocokan,
				rekomendasi[tengah].Industri, rekomendasi[tengah].Gaji)
			kiri := tengah - 1
			for kiri >= 0 && rekomendasi[kiri].Industri == industriDicari {
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					rekomendasi[kiri].Nama, rekomendasi[kiri].Kecocokan,
					rekomendasi[kiri].Industri, rekomendasi[kiri].Gaji)
				kiri--
			}
			kanan := tengah + 1
			for kanan < len(rekomendasi) && rekomendasi[kanan].Industri == industriDicari {
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					rekomendasi[kanan].Nama, rekomendasi[kanan].Kecocokan,
					rekomendasi[kanan].Industri, rekomendasi[kanan].Gaji)
				kanan++
			}
			break
		} else if rekomendasi[tengah].Industri < industriDicari {
			rendah = tengah + 1
		} else {
			tinggi = tengah - 1
		}
	}

	if !found {
		fmt.Println("| Tidak ditemukan karir dengan industri tersebut |")
	}
	fmt.Println("+-------------------------+--------------+---------------+--------------+")
}
package main

import "fmt"

var (
	minatList    = []string{}
	keahlianList = []string{}
)

type Karir struct {
	Nama      string
	Kecocokan int
	Industri  string
	Gaji      string
}

type Pekerjaan struct {
	Nama string
	Gaji string
}

func main() {
	var keluar bool
	var pilihan int

	for !keluar {
		fmt.Println("\nMenu Pilihan:")
		fmt.Println("1. Tambah minat")
		fmt.Println("2. Tambah keahlian")
		fmt.Println("3. Edit minat")
		fmt.Println("4. Edit keahlian")
		fmt.Println("5. Karir rekomendasi")
		fmt.Println("6. Tampilkan Minat Dan Keahlian ")
		fmt.Println("7. Pekerjaan ")
		fmt.Println("8. Keluar")
		fmt.Print("Masukkan pilihan (1-8): ")
		fmt.Scan(&pilihan)
		if pilihan >= 9 {
			fmt.Println("Input tidak valid, silakan coba lagi")
			continue
		}

		switch pilihan {
		case 1:
			tambahMinat()
		case 2:
			tambahKeahlian()
		case 3:
			editMinat()
		case 4:
			editKeahlian()
		case 5:
			karirRekomendasi()
		case 6:
			tampilMinatDanKeahlian(minatList, keahlianList)
		case 7:
			pekerjaan()
		case 8:
			fmt.Println("Terima kasih, program selesai.")
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid, silakan pilih 1-8")
		}
	}
}

func tambahMinat() {
	var input string

	daftarMinat := []string{
		"Membaca",
		"Menulis",
		"Menggambar",
		"Fotografi",
		"Menghitung",
		"Menganalisis",
		"Melukis",
		"Bermain alat musik",
	}

	fmt.Println("\n=== Tambah Minat ===")
	fmt.Println("Daftar minat yang tersedia:")
	for i, minat := range daftarMinat {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	sisa := 3 - len(minatList)
	if sisa <= 0 {
		fmt.Println("\nYahh kamu sudah menambahkan 3 minat")
		return
	}

	fmt.Printf("\nKamu bisa mengisi %d minat (ketik '-' untuk menyudahi tambah minat)\n", sisa)
	for i := 0; i < sisa; i++ {
		fmt.Printf("Pilihan %d (1-8 atau - untuk berhenti): ", i+1)
		fmt.Scan(&input)

		if input == "-" {
			fmt.Println("Kamu Menghentikan Penambahan.")
			break
		}

		var pilihan int
		_, err := fmt.Sscan(input, &pilihan)
		if err != nil || pilihan < 1 || pilihan  > len(daftarMinat) {
			fmt.Println("Input tidak valid")
			return
		}

		duplikat := false
		for _, m := range minatList {
			if m == daftarMinat[pilihan-1] {
				duplikat = true
				break
			}
		}
		if duplikat {
			fmt.Println("Minat sudah ada, coba yang lain.")
			i--
			continue
		}

		minatList = append(minatList, daftarMinat[pilihan-1])
		fmt.Println("--Minat berhasil ditambahkan--")
	}

	fmt.Println("Minat kamu sekarang:", minatList)
}

func tambahKeahlian() {
	var pilihan, sisa int
	sisa = 1 - len(keahlianList)

	fmt.Println("\n=== Tambah Keahlian ===")
	daftarKeahlian := []string{
		"Coding",
		"Analisis",
		"Design Grafis",
		"Editing Video",
		"Menghafal",
		"Mengamati",
		"Melukis",
	}

	fmt.Println("Daftar Keahlian yang tersedia:")
	for i, ahli := range daftarKeahlian {
		fmt.Printf("%d. %s\n", i+1, ahli)
	}

	if sisa <= 0 {
		fmt.Println("\nYahh kamu sudah menambahkan 1 keahlian")
		return
	}

	fmt.Printf("\nKamu bisa mengisi %d keahlian\n", sisa)
	for i := 0; i < sisa; i++ {
		fmt.Printf("Pilihan %d (1-7): ", i+1)
		fmt.Scan(&pilihan)

		if pilihan < 1 || pilihan > 7 {
			fmt.Println("Input tidak valid")
			return
		}

		keahlianList = append(keahlianList, daftarKeahlian[pilihan-1])
		fmt.Println("--Keahlian berhasil ditambahkan--")
	}
}

func editMinat() {
	var pilihan int

	fmt.Println("\n=== Edit Minat ===")
	if len(minatList) == 0 {
		fmt.Println("Belum ada minat yang tersedia")
		return
	}
	fmt.Println("Minat Anda saat ini:")
	for i, minat := range minatList {
		fmt.Printf("%d. %s\n", i+1, minat)
	}

	fmt.Println("\n pilih menu dibawah ini:")
	fmt.Println("1. Tambah Minat")
	fmt.Println("2. Hapus Minat (pilih 1)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilihan kamu: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahMinat()
	case 2:
		hapusMinat()
	case 3:
		return
	default:
		fmt.Println("Pilihan Tidak Valid")
	}
}

func editKeahlian() {
	var pilihan int

	fmt.Println("\n=== Edit Keahlian ===")
	if len(keahlianList) == 0 {
		fmt.Println("Belum ada Keahlian yang tersedia")
		return
	}
	fmt.Println("Keahlian Anda saat ini:")
	for i, keahlian := range keahlianList {
		fmt.Printf("%d. %s\n", i+1, keahlian)
	}
	fmt.Println("\n pilih menu dibawah ini:")
	fmt.Println("1. Tambah keahlian")
	fmt.Println("2. Hapus keahlian (pilih 1)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilihan kamu: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		tambahKeahlian()
	case 2:
		hapusKeahlian()
	case 3:
		return
	default:
		fmt.Println("Pilihan Tidak Valid")
	}
}

func hapusMinat() {
	fmt.Println("\n=== Hapus Minat ===")
	var pilihan int

	fmt.Print("\nPilih nomor minat yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(minatList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	minatTerhapus := minatList[pilihan-1]
	minatList = append(minatList[:pilihan-1], minatList[pilihan:]...)
	fmt.Println("Minat berhasil dihapus\n", minatTerhapus, 3-len(minatList))
}

func hapusKeahlian() {
	var pilihan int

	fmt.Print("\nPilih nomor keahlian yang akan dihapus: ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > len(keahlianList) {
		fmt.Println("Nomor tidak valid")
		return
	}

	keahlianTerhapus := keahlianList[pilihan-1]
	keahlianList = append(keahlianList[:pilihan-1], keahlianList[pilihan:]...)
	fmt.Printf("Keahlian '%s' berhasil dihapus\n", keahlianTerhapus)
}
func karirRekomendasi() {
	var rekomendasi []Karir
	var presentase int
	fmt.Println("\n=== Rekomendasi Karir ===")

	if len(minatList) == 0 && len(keahlianList) == 0 {
		fmt.Println("Belum ada data minat atau keahlian")
		return
	}

	fmt.Println("Berdasarkan minat dan keahlian Kamu:")
	fmt.Println("- Minat:", minatList)
	fmt.Println("- Keahlian:", keahlianList)

	for _, minat := range minatList {
		for _, keahlian := range keahlianList {
			if (minat == "Menggambar" || minat == "Melukis") && keahlian == "Design Grafis" {
				rekomendasi = append(rekomendasi, Karir{"UI/UX Designer", 90, "Teknologi", "6-10 juta"})
			} else if minat == "Membaca" && keahlian == "Coding" {
				rekomendasi = append(rekomendasi, Karir{"Software Engineer", 85, "Teknologi", "7-15 juta"})
			} else if minat == "Mengamati" || keahlian == "Menghitung" || keahlian == "Analisis" {
				rekomendasi = append(rekomendasi, Karir{"Data Analyst", 80, "Data", "6-12 juta"})
			} else if minat == "Fotografi" && keahlian == "Editing Video" {
				rekomendasi = append(rekomendasi, Karir{"Video Editor", 75, "Kreatif", "4-8 juta"})
			} else if (minat == "Menggambar") && (keahlian == "Coding" || keahlian == "Design Grafis") {
				rekomendasi = append(rekomendasi, Karir{"Front-End Developer", 85, "Teknologi", "6-10 juta"})
			} else if minat == "Menulis" && (keahlian == "Menghafal" || keahlian == "Mengamati") {
				rekomendasi = append(rekomendasi, Karir{"Technical Writer", 70, "Penulisan", "6-12 juta"})
			} else if minat == "Bermain alat musik" && keahlian == "Menghafal" {
				rekomendasi = append(rekomendasi, Karir{"Music Content Editor", 65, "Musik", "5-9 juta"})
			} else if minat == "Menulis" && (keahlian == "Mengamati" || keahlian == "Coding") {
				rekomendasi = append(rekomendasi, Karir{"SEO Specialist", 75, "Digital Marketing", "5-8 juta"})
			} else if minat == "Melukis" && keahlian == "Design Grafis" {
				rekomendasi = append(rekomendasi, Karir{"Ilustrator Digital", 80, "Kreatif", "5-9 juta"})
			} else if minat == "Membaca" && (keahlian == "Menulis" || keahlian == "Menggambar") {
				rekomendasi = append(rekomendasi, Karir{"Content Strategist", 70, "Marketing", "6-11 juta"})
			}
		}
	}
	if len(rekomendasi) > 0 {
		fmt.Println("\nIni adalah rekomendasi karir yang cocok:")
		fmt.Println("+-------------------------+--------------+---------------+--------------+")
		fmt.Println("|      PEKERJAAN          | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
		fmt.Println("+-------------------------+--------------+---------------+--------------+")

		for _, karir := range rekomendasi {
			fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
				karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
		}

		fmt.Println("+-------------------------+--------------+---------------+--------------+")
		fmt.Println("Pilih yuk")
		fmt.Println("1. Urut Presentasi kecocokan paling tinggi")
		fmt.Println("2. Urut Presentasi kecocokan paling Rendah")
		fmt.Println("3. Cari Industri")
		fmt.Println("4. Kembali ke menu utama")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&presentase)

		switch presentase {
		case 1:
			for i := 0; i < len(rekomendasi)-1; i++ {
				minIdx := i
				for j := i + 1; j < len(rekomendasi); j++ {
					if rekomendasi[j].Kecocokan > rekomendasi[minIdx].Kecocokan {
						minIdx = j
					}
				}
				rekomendasi[i], rekomendasi[minIdx] = rekomendasi[minIdx], rekomendasi[i]
			}

			fmt.Println("\nUrutan rekomendasi Tertinggi:")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			for i := 0; i < len(rekomendasi); i++ {
				karir := rekomendasi[i]
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
			}
			fmt.Println("+-------------------------+--------------+---------------+--------------+")

		case 2:
			for i := 1; i < len(rekomendasi); i++ {
				key := rekomendasi[i]
				j := i - 1
				for j >= 0 && rekomendasi[j].Kecocokan > key.Kecocokan {
					rekomendasi[j+1] = rekomendasi[j]
					j--
				}
				rekomendasi[j+1] = key
			}

			fmt.Println("\nUrutan rekomendasi Terendah:")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
			fmt.Println("+-------------------------+--------------+---------------+--------------+")
			for i := 0; i < len(rekomendasi); i++ {
				karir := rekomendasi[i]
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					karir.Nama, karir.Kecocokan, karir.Industri, karir.Gaji)
			}
			fmt.Println("+-------------------------+--------------+---------------+--------------+")

		case 3:
			var industriDicari string
			fmt.Print("Masukkan industri yang dicari: ")
			fmt.Scan(&industriDicari)
			binarySearch(rekomendasi, industriDicari)

		case 4:
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	} else {
		fmt.Println("Maaf, belum ada karir yang cocok ditemukan berdasarkan minat dan keahlian kamu.")
	}
}

func tampilMinatDanKeahlian(minatList []string, keahlianList []string) {
	var minat, keahlian string
	fmt.Println("\n=== Daftar Minat & Keahlian ===")
	fmt.Printf("%-4s| %-20s| %-20s\n", "No", "Minat", "Keahlian")
	fmt.Println("----+--------------------+----------------------")

	max := len(minatList)
	if len(keahlianList) > max {
		max = len(keahlianList)
	}

	for i := 0; i < max; i++ {
		if i < len(minatList) {
			minat = minatList[i]
		}
		if i < len(keahlianList) {
			keahlian = keahlianList[i]
		}
		fmt.Printf("%-4d| %-20s| %-20s\n", i+1, minat, keahlian)
	}
}

func pekerjaan() {
	var pilihan, inputAngka int
	pekerjaanData := []Pekerjaan{
		{"UI/UX Designer", "6-10 juta"},
		{"Software Engineer", "7-15 juta"},
		{"Data Analyst", "6-12 juta"},
		{"Video Editor", "4-8 juta"},
		{"Front-End Developer", "6-10 juta"},
		{"Technical Writer", "6-12 juta"},
		{"Music Content Editor", "5-9 juta"},
		{"SEO Specialist", "5-8 juta"},
		{"Illustrator Digital", "5-9 juta"},
		{"Content Strategist", "6-11 juta"},
	}

	fmt.Println("\n+-------------------------+--------------+")
	fmt.Println("|      PEKERJAAN         |     GAJI     |")
	fmt.Println("+-------------------------+--------------+")
	for i := 0; i < len(pekerjaanData); i++ {
		fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
	}
	fmt.Println("+-------------------------+--------------+")
	fmt.Println("\nPilih urutan:")
	fmt.Println("1. Gaji Terbesar ke Terkecil")
	fmt.Println("2. Gaji Terkecil ke Besar")
	fmt.Println("3. Cari pekerjaan berdasarkan gaji")
	fmt.Println("4. Kembali ke Menu Utama")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		for i := 0; i < len(pekerjaanData)-1; i++ {
			maxIdx := i
			for j := i + 1; j < len(pekerjaanData); j++ {
				if pekerjaanData[j].Gaji > pekerjaanData[maxIdx].Gaji {
					maxIdx = j
				}
			}
			pekerjaanData[i], pekerjaanData[maxIdx] = pekerjaanData[maxIdx], pekerjaanData[i]
		}

		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")
		for i := 0; i < len(pekerjaanData); i++ {
			fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
		}
		fmt.Println("+-------------------------+--------------+")

	case 2:
		for i := 1; i < len(pekerjaanData); i++ {
			key := pekerjaanData[i]
			j := i - 1
			for j >= 0 && pekerjaanData[j].Gaji > key.Gaji {
				pekerjaanData[j+1] = pekerjaanData[j]
				j--
			}
			pekerjaanData[j+1] = key
		}

		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")
		for i := 0; i < len(pekerjaanData); i++ {
			fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, pekerjaanData[i].Gaji)
		}
		fmt.Println("+-------------------------+--------------+")

	case 3:

		fmt.Print("\nMasukkan angka awal gaji (contoh: 6 untuk '6-10 juta'): ````````````````````````````````                                                                                                `")
		fmt.Scan(&inputAngka)

		found := false
		fmt.Println("\n+-------------------------+--------------+")
		fmt.Println("|      PEKERJAAN         |     GAJI     |")
		fmt.Println("+-------------------------+--------------+")

		for i := 0; i < len(pekerjaanData); i++ {
			gaji := pekerjaanData[i].Gaji
			angkaAwal := int(gaji[0] - '0')

			if angkaAwal == inputAngka {
				fmt.Printf("| %-23s | %-12s |\n", pekerjaanData[i].Nama, gaji)
				found = true
			}
		}

		fmt.Println("+-------------------------+--------------+")
		if !found {
			fmt.Printf("Tidak ada pekerjaan dengan gaji mulai dari %d\n", inputAngka)
		}
	case 4:
		return
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func binarySearch(rekomendasi []Karir, industriDicari string) {
	for i := 0; i < len(rekomendasi)-1; i++ {
		for j := 0; j < len(rekomendasi)-i-1; j++ {
			if rekomendasi[j].Industri > rekomendasi[j+1].Industri {
				rekomendasi[j], rekomendasi[j+1] = rekomendasi[j+1], rekomendasi[j]
			}
		}
	}
	rendah := 0
	tinggi := len(rekomendasi) - 1
	found := false

	fmt.Println("\nHasil pencarian industri", industriDicari, ":")
	fmt.Println("+-------------------------+--------------+---------------+--------------+")
	fmt.Println("|      PEKERJAAN         | KECOCOKAN (%)|   INDUSTRI    |     GAJI     |")
	fmt.Println("+-------------------------+--------------+---------------+--------------+")

	for rendah <= tinggi {
		tengah := (rendah + tinggi) / 2

		if rekomendasi[tengah].Industri == industriDicari {
			found = true
			fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
				rekomendasi[tengah].Nama, rekomendasi[tengah].Kecocokan,
				rekomendasi[tengah].Industri, rekomendasi[tengah].Gaji)
			kiri := tengah - 1
			for kiri >= 0 && rekomendasi[kiri].Industri == industriDicari {
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					rekomendasi[kiri].Nama, rekomendasi[kiri].Kecocokan,
					rekomendasi[kiri].Industri, rekomendasi[kiri].Gaji)
				kiri--
			}
			kanan := tengah + 1
			for kanan < len(rekomendasi) && rekomendasi[kanan].Industri == industriDicari {
				fmt.Printf("| %-23s | %-12d | %-13s | %-12s |\n",
					rekomendasi[kanan].Nama, rekomendasi[kanan].Kecocokan,
					rekomendasi[kanan].Industri, rekomendasi[kanan].Gaji)
				kanan++
			}
			break
		} else if rekomendasi[tengah].Industri < industriDicari {
			rendah = tengah + 1
		} else {
			tinggi = tengah - 1
		}
	}

	if !found {
		fmt.Println("| Tidak ditemukan karir dengan industri tersebut |")
	}
	fmt.Println("+-------------------------+--------------+---------------+--------------+")
}
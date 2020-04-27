package main

import (
	"fmt"
	"time"
)

func main() {
	//Format Tgl : Tahun, Bulan, Tgl
	inputTglPertama := date(2020, 1, 1)
	inputTglKedua := date(2020, 3, 10)

	jumlahHari := inputTglKedua.Sub(inputTglPertama).Hours() / 24

	fmt.Println("Jumlah hari dari tanggal 1 Januari - 10 Maret adalah", jumlahHari, "hari")
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

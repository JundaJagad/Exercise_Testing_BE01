package main

import (
	"errors"
	"fmt"
	"testing"
)

// func HitungHargaTotal
func HitungHargaTotal(hargaItem, ongkir float64, qty int) (float64, error) {
	if hargaItem <= 0 || ongkir < 0 || qty <= 0 {
		return 0, errors.New("harga atau kuantitas tidak valid")
	}
	return (hargaItem * float64(qty)) + ongkir, nil
}

// func PembayaranBarang
func PembayaranBarang(hargaTotal float64, metodePembayaran string, dicicil bool) error {
	// cek hargaTotal > 0
	if hargaTotal <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	// cek metode pembayaran
	switch metodePembayaran {
	case "cod", "transfer", "debit", "credit", "gerai":
		// metode dikenali
	default:
		return errors.New("metode tidak dikenali")
	}

	// cek dicicil atau tidak
	if dicicil {
		// jika dicicil, metode harus credit, dan hargaTotal harus >= 500.000
		if metodePembayaran != "credit" || hargaTotal < 500000 {
			return errors.New("cicilan tidak memenuhi syarat")
		}
	} else {
		// jika tidak dicicil, metode harus bukan credit
		if metodePembayaran == "credit" {
			return errors.New("credit harus dicicil")
		}
	}

	// tidak ada kesalahan
	return nil
}

func TestHitungHargaTotal(t *testing.T) {
	type args struct {
		hargaItem float64
		ongkir    float64
		qty       int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"case 1", args{100000, 20000, 2}, 220000, false},
		{"case 2", args{0, 20000, 2}, 0, true},
		{"case 3", args{100000, -20000, 2}, 0, true},
		{"case 4", args{100000, 20000, 0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HitungHargaTotal(tt.args.hargaItem, tt.args.ongkir, tt.args.qty)
			if (err != nil) != tt.wantErr {
				t.Errorf("HitungHargaTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HitungHargaTotal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func main() {
	// Tes fungsi HitungHargaTotal
	fmt.Println(HitungHargaTotal(100000, 20000, 2)) // 220000

	// Tes fungsi PembayaranBarang
	fmt.Println(PembayaranBarang(0, "cod", false))          // harga tidak bisa nol
	fmt.Println(PembayaranBarang(600000, "credit", false))  // credit harus dicicil
	fmt.Println(PembayaranBarang(400000, "credit", true))   // cicilan tidak memenuhi syarat
	fmt.Println(PembayaranBarang(600000, "debit", true))    // cicilan tidak memenuhi syarat
	fmt.Println(PembayaranBarang(600000, "credit", true))   // <nil>
	fmt.Println(PembayaranBarang(600000, "transfer", true)) // <nil>
}

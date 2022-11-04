package models

import "time"

type PresetKriteria struct {
	ID                     uint      `gorm:"primaryKey" json:"id"`
	SetoranAwal            float64   `json:"setoran_awal"`
	SetoranLanjutanMinimal float64   `json:"setoran_lanjutan_minimal"`
	SaldoMinimum           float64   `json:"saldo_minimum"`
	SukuBunga              float64   `json:"suku_bunga"`
	Fungsionalitas         float64   `json:"fungsionalitas"`
	BiayaAdmin             float64   `json:"biaya_admin"`
	BiayaPenarikanHabis    float64   `json:"biaya_penarikan_habis"`
	KategoriUmurPengguna   float64   `json:"kategori_umur_pengguna"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

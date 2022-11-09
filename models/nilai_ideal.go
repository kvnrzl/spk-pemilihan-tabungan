package models

import "time"

type NilaiIdeal struct {
	ID                     uint                 `gorm:"primaryKey" json:"id"`
	SetoranAwal            int                  `json:"setoran_awal"`
	SetoranLanjutanMinimal int                  `json:"setoran_lanjutan_minimal"`
	SaldoMinimum           int                  `json:"saldo_minimum"`
	SukuBunga              float64              `json:"suku_bunga"`
	FungsionalitasID       uint                 `json:"fungsionalitas_id"`
	Fungsionalitas         Fungsionalitas       `json:"fungsionalitas"`
	BiayaAdmin             int                  `json:"biaya_admin"`
	BiayaPenarikanHabis    int                  `json:"biaya_penarikan_habis"`
	KategoriUmurPenggunaID uint                 `json:"kategori_umur_pengguna_id"`
	KategoriUmurPengguna   KategoriUmurPengguna `json:"kategori_umur_pengguna"`
	CreatedAt              time.Time            `json:"created_at"`
	UpdatedAt              time.Time            `json:"updated_at"`
}

type Fungsionalitas struct {
	ID            uint `gorm:"primaryKey" json:"id"`
	Investasi     uint `json:"investasi"`
	Bisnis        uint `json:"bisnis"`
	Transaksional uint `json:"transaksional"`
}

type KategoriUmurPengguna struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	Dewasa uint `json:"dewasa"`
	Remaja uint `json:"remaja"`
	Anak   uint `json:"Anak"`
}

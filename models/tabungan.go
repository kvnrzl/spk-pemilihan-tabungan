package models

import "time"

type Tabungan struct {
	ID                     uint      `gorm:"primarykey" json:"id"`
	NamaTabungan           string    `json:"nama_tabungan"`
	SetoranAwal            int       `json:"setoran_awal"`
	SetoranLanjutanMinimal int       `json:"setoran_lanjutan_minimal"`
	SaldoMinimum           int       `json:"saldo_minimum"`
	SukuBunga              float64   `json:"suku_bunga"`
	Fungsionalitas         string    `json:"fungsionalitas"`
	BiayaAdmin             int       `json:"biaya_admin"`
	BiayaPenarikanHabis    int       `json:"biaya_penarikan_habis"`
	KategoriUmurPengguna   string    `json:"kategori_umur_pengguna"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

//! Perbaikan :
/*
	1. buat yg bertipe data string menjadi UPPERCASE / LOWERCASE (penyeragaman data)
	2. buat yg bisa mengakses beberapa endpoint harus memiliki cookie name "jwt" / hak admin
*/

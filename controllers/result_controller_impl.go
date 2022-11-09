package controllers

import (
	"fmt"
	"net/http"
	"project_spk_pemilihan_tabungan/models"
	"project_spk_pemilihan_tabungan/services"

	"github.com/gin-gonic/gin"
)

type Output struct {
	ID           string          `json:"id"`
	NilaiIdealID uint            `json:"nilai_ideal_id"`
	TabunganID   uint            `json:"tabungan_id"`
	Tabungan     models.Tabungan `json:"tabungan"`
	Skor         float64         `json:"skor"`
}

type NewTabungan struct {
	ID                     uint    `gorm:"primaryKey" json:"id"`
	NamaTabungan           string  `json:"nama_tabungan"`
	SetoranAwal            float64 `json:"setoran_awal"`
	SetoranLanjutanMinimal float64 `json:"setoran_lanjutan_minimal"`
	SaldoMinimum           float64 `json:"saldo_minimum"`
	SukuBunga              float64 `json:"suku_bunga"`
	Fungsionalitas         uint    `json:"fungsionalitas"`
	BiayaAdmin             float64 `json:"biaya_admin"`
	BiayaPenarikanHabis    float64 `json:"biaya_penarikan_habis"`
	KategoriUmurPengguna   uint    `json:"kategori_umur_pengguna"`
	// CreatedAt              time.Time `json:"created_at"`
	// UpdatedAt              time.Time `json:"updated_at"`
}

type HasilAkhir struct {
	NamaTabungan string  `json:"nama_tabungan"`
	Skor         float64 `json:"skor"`
}

func interpolasiLinear(nilaiKriteria float64, maxKriteria float64, skorMax float64, skorMin float64) float64 {
	var skor float64
	if nilaiKriteria < 0 {
		skor = ((nilaiKriteria-(-maxKriteria))/(0-(-maxKriteria)))*(skorMax-skorMin) + skorMin
	} else {
		skor = (nilaiKriteria/maxKriteria)*(skorMin-skorMax) + skorMax
	}

	return skor
}

func saw(alternatif NewTabungan, preset *models.PresetKriteria) float64 {
	return (alternatif.SetoranAwal * preset.SetoranAwal) + (alternatif.SetoranLanjutanMinimal * preset.SetoranLanjutanMinimal) + (alternatif.SaldoMinimum * preset.SaldoMinimum) + (alternatif.SukuBunga * preset.SukuBunga) + (alternatif.BiayaAdmin * preset.BiayaAdmin) + (alternatif.BiayaPenarikanHabis * preset.BiayaPenarikanHabis) + (float64(alternatif.Fungsionalitas) * preset.Fungsionalitas) + (float64(alternatif.KategoriUmurPengguna) * preset.KategoriUmurPengguna)
}

type ResultControllerImpl struct {
	tabunganService       services.TabunganService
	presetKriteriaService services.PresetKriteriaService
}

func NewResultController(tabunganService services.TabunganService, presetKriteriaService services.PresetKriteriaService) ResultController {
	return &ResultControllerImpl{tabunganService, presetKriteriaService}
}

func (c *ResultControllerImpl) HitungResult(ctx *gin.Context) {
	var nilaiIdeal models.NilaiIdeal
	err := ctx.ShouldBindJSON(&nilaiIdeal)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}

	result, err := c.tabunganService.FindAllTabungan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}

	// SKORING
	var newTabungans []NewTabungan
	var maxSetoranAwal, maxSetoranLanjutanMinimal, maxSaldoMinimum, maxBiayaAdmin, maxBiayaPenarikanHabis float64
	var maxSukuBunga float64
	for _, tabungan := range result {
		var newTabungan NewTabungan
		newTabungan.NamaTabungan = tabungan.NamaTabungan

		newTabungan.SetoranAwal = float64(tabungan.SetoranAwal - nilaiIdeal.SetoranAwal)
		if newTabungan.SetoranAwal > maxSetoranAwal {
			maxSetoranAwal = newTabungan.SetoranAwal
		}
		newTabungan.SetoranLanjutanMinimal = float64(tabungan.SetoranLanjutanMinimal - nilaiIdeal.SetoranLanjutanMinimal)
		if newTabungan.SetoranLanjutanMinimal > maxSetoranLanjutanMinimal {
			maxSetoranLanjutanMinimal = newTabungan.SetoranLanjutanMinimal
		}
		newTabungan.SaldoMinimum = float64(tabungan.SaldoMinimum - nilaiIdeal.SaldoMinimum)
		if newTabungan.SaldoMinimum > maxSaldoMinimum {
			maxSaldoMinimum = newTabungan.SaldoMinimum
		}
		newTabungan.SukuBunga = tabungan.SukuBunga - nilaiIdeal.SukuBunga
		if newTabungan.SukuBunga > maxSukuBunga {
			maxSukuBunga = newTabungan.SukuBunga
		}
		newTabungan.BiayaAdmin = float64(tabungan.BiayaAdmin - nilaiIdeal.BiayaAdmin)
		if newTabungan.BiayaAdmin > maxBiayaAdmin {
			maxBiayaAdmin = newTabungan.BiayaAdmin
		}
		newTabungan.BiayaPenarikanHabis = float64(tabungan.BiayaPenarikanHabis - nilaiIdeal.BiayaPenarikanHabis)
		if newTabungan.BiayaPenarikanHabis > maxBiayaPenarikanHabis {
			maxBiayaPenarikanHabis = newTabungan.BiayaPenarikanHabis
		}

		if tabungan.Fungsionalitas == "INVESTASI" {
			newTabungan.Fungsionalitas = nilaiIdeal.Fungsionalitas.Investasi
		} else if tabungan.Fungsionalitas == "BISNIS" {
			newTabungan.Fungsionalitas = nilaiIdeal.Fungsionalitas.Bisnis
		} else {
			newTabungan.Fungsionalitas = nilaiIdeal.Fungsionalitas.Transaksional
		}

		if tabungan.KategoriUmurPengguna == "DEWASA" {
			newTabungan.KategoriUmurPengguna = nilaiIdeal.KategoriUmurPengguna.Dewasa
		} else if tabungan.KategoriUmurPengguna == "REMAJA" {
			newTabungan.KategoriUmurPengguna = nilaiIdeal.KategoriUmurPengguna.Remaja
		} else {
			newTabungan.KategoriUmurPengguna = nilaiIdeal.KategoriUmurPengguna.Anak
		}

		newTabungans = append(newTabungans, newTabungan)
	}

	fmt.Println("isi dari newtabungans: ", newTabungans)

	// INTERPOLASI
	var newTabunganInterpolasi []NewTabungan
	for _, tabungan := range newTabungans {
		var newTabungan NewTabungan
		newTabungan.NamaTabungan = tabungan.NamaTabungan
		newTabungan.SetoranAwal = interpolasiLinear(tabungan.SetoranAwal, maxSetoranAwal, 5, 1)
		newTabungan.SetoranLanjutanMinimal = interpolasiLinear(tabungan.SetoranLanjutanMinimal, maxSetoranLanjutanMinimal, 5, 1)
		newTabungan.SaldoMinimum = interpolasiLinear(tabungan.SaldoMinimum, maxSaldoMinimum, 5, 1)
		newTabungan.SukuBunga = interpolasiLinear(tabungan.SukuBunga, maxSukuBunga, 5, 1)
		newTabungan.BiayaAdmin = interpolasiLinear(tabungan.BiayaAdmin, maxBiayaAdmin, 5, 1)
		newTabungan.BiayaPenarikanHabis = interpolasiLinear(tabungan.BiayaPenarikanHabis, maxBiayaPenarikanHabis, 5, 1)
		newTabungan.Fungsionalitas = tabungan.Fungsionalitas
		newTabungan.KategoriUmurPengguna = tabungan.KategoriUmurPengguna

		newTabunganInterpolasi = append(newTabunganInterpolasi, newTabungan)
	}

	fmt.Println("isi dari newtabunganinterpolasi : ", newTabunganInterpolasi)

	preset, err := c.presetKriteriaService.FindFirstPreset(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}

	// METODE SAW
	var hasilAkhir []HasilAkhir
	for _, tabungan := range newTabunganInterpolasi {
		var newHasilAkhir HasilAkhir
		newHasilAkhir.NamaTabungan = tabungan.NamaTabungan
		newHasilAkhir.Skor = saw(tabungan, preset)

		hasilAkhir = append(hasilAkhir, newHasilAkhir)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    hasilAkhir,
	})
}

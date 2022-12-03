package controllers

import (
	"fmt"
	"math"
	"net/http"
	"project_spk_pemilihan_tabungan/models"
	"project_spk_pemilihan_tabungan/services"
	"sort"

	"github.com/gin-gonic/gin"
)

// type Output struct {
// 	ID           string          `json:"id"`
// 	NilaiIdealID uint            `json:"nilai_ideal_id"`
// 	TabunganID   uint            `json:"tabungan_id"`
// 	Tabungan     models.Tabungan `json:"tabungan"`
// 	Skor         float64         `json:"skor"`
// }

type NewTabungan struct {
	ID                     uint    `gorm:"primaryKey" json:"id"`
	NamaTabungan           string  `json:"nama_tabungan"`
	SetoranAwal            float64 `json:"setoran_awal"`
	SetoranLanjutanMinimal float64 `json:"setoran_lanjutan_minimal"`
	SaldoMinimum           float64 `json:"saldo_minimum"`
	SukuBunga              float64 `json:"suku_bunga"`
	BiayaAdmin             float64 `json:"biaya_admin"`
	BiayaPenarikanHabis    float64 `json:"biaya_penarikan_habis"`
	Fungsionalitas         uint    `json:"fungsionalitas"`
	KategoriUmurPengguna   uint    `json:"kategori_umur_pengguna"`
	// CreatedAt              time.Time `json:"created_at"`
	// UpdatedAt              time.Time `json:"updated_at"`
}

type HasilAkhir struct {
	ID           uint    `json:"id"`
	NamaTabungan string  `json:"nama_tabungan"`
	Skor         float64 `json:"skor"`
}

// type InputRecomendation struct {
// 	ID               uint                  `json:"id"`
// 	NilaiIdealID     uint                  `json:"nilai_ideal_id"`
// 	NilaiIdeal       models.NilaiIdeal     `json:"nilai_ideal"`
// 	PresetKriteriaID uint                  `json:"preset_kriteria_id"`
// 	PresetKriteria   models.PresetKriteria `json:"preset_kriteria"`
// }

func interpolasiLinear(nilaiKriteria float64, maxKriteria float64, skorMax float64, skorMin float64) float64 {
	var skor float64
	if nilaiKriteria < 0 {
		skor = ((nilaiKriteria-(-maxKriteria))/(0-(-maxKriteria)))*(skorMax-skorMin) + skorMin
	} else {
		skor = (nilaiKriteria/maxKriteria)*(skorMin-skorMax) + skorMax
	}

	return math.Round(skor*100) / 100
}

func interpolasiLinearBenefit(data float64, dataMin float64, dataMax float64) float64 {
	return ((data - dataMin) / (dataMax - dataMin) * (5 - 1)) + 1
}

func interpolasiLinearCost(data float64, dataMin float64, dataMax float64) float64 {
	return ((data - dataMin) / (dataMax - dataMin) * (1 - 5)) + 5
}

func saw(alternatif NewTabungan, preset models.BobotKriteria) float64 {
	skor := (alternatif.SetoranAwal * preset.SetoranAwal) +
		(alternatif.SetoranLanjutanMinimal * preset.SetoranLanjutanMinimal) +
		(alternatif.SaldoMinimum * preset.SaldoMinimum) +
		(alternatif.SukuBunga * preset.SukuBunga) +
		(alternatif.BiayaAdmin * preset.BiayaAdmin) +
		(alternatif.BiayaPenarikanHabis * preset.BiayaPenarikanHabis) +
		(float64(alternatif.Fungsionalitas) * preset.Fungsionalitas) +
		(float64(alternatif.KategoriUmurPengguna) * preset.KategoriUmurPengguna)

	return math.Round(skor*100) / 100
}

type ResultControllerImpl struct {
	tabunganService           services.TabunganService
	presetKriteriaService     services.PresetKriteriaService
	inputRecomendationService services.InputRecomendationService
}

func NewResultController(tabunganService services.TabunganService, presetKriteriaService services.PresetKriteriaService, inputService services.InputRecomendationService) ResultController {
	return &ResultControllerImpl{
		tabunganService:           tabunganService,
		presetKriteriaService:     presetKriteriaService,
		inputRecomendationService: inputService,
	}
}

func (c *ResultControllerImpl) HitungResult(ctx *gin.Context) {
	// var nilaiIdeal models.NilaiIdeal
	// err := ctx.ShouldBindJSON(&nilaiIdeal)
	// if err != nil {
	// 	ctx.JSON(400, gin.H{
	// 		"code":  400,
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	var input models.InputRecomendation

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}

	_, err = c.inputRecomendationService.Create(ctx, &input)
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
	var maxSetoranAwal, maxSetoranLanjutanMinimal, maxSaldoMinimum, maxBiayaAdmin, maxBiayaPenarikanHabis, maxSukuBunga float64
	var minSukuBunga = result[1].SukuBunga
	var minBiayaAdmin = float64(result[1].BiayaAdmin)
	var minBiayaPenarikanHabis = float64(result[1].BiayaPenarikanHabis)

	for _, tabungan := range result {
		var newTabungan NewTabungan

		newTabungan.ID = tabungan.ID
		newTabungan.NamaTabungan = tabungan.NamaTabungan

		newTabungan.SetoranAwal = float64(tabungan.SetoranAwal - input.NilaiIdeal.SetoranAwal)
		newTabungan.SetoranLanjutanMinimal = float64(tabungan.SetoranLanjutanMinimal - input.NilaiIdeal.SetoranLanjutanMinimal)
		newTabungan.SaldoMinimum = float64(tabungan.SaldoMinimum - input.NilaiIdeal.SaldoMinimum)

		if input.NilaiIdeal.SukuBunga == -1 {
			newTabungan.SukuBunga = float64(tabungan.SukuBunga)
			if tabungan.SukuBunga < minSukuBunga {
				minSukuBunga = tabungan.SukuBunga
			}
			if tabungan.SukuBunga > maxSukuBunga {
				maxSukuBunga = tabungan.SukuBunga
			}
		} else {
			newTabungan.SukuBunga = tabungan.SukuBunga - input.NilaiIdeal.SukuBunga
			if math.Abs(newTabungan.SukuBunga) > maxSukuBunga {
				maxSukuBunga = math.Abs(newTabungan.SukuBunga)
			}
		}
		if input.NilaiIdeal.BiayaAdmin == -1 {
			newTabungan.BiayaAdmin = float64(tabungan.BiayaAdmin)
			if float64(tabungan.BiayaAdmin) < minBiayaAdmin {
				minBiayaAdmin = float64(tabungan.BiayaAdmin)
			}
			if float64(tabungan.BiayaAdmin) > maxBiayaAdmin {
				maxBiayaAdmin = float64(tabungan.BiayaAdmin)
			}
		} else {
			newTabungan.BiayaAdmin = float64(tabungan.BiayaAdmin - input.NilaiIdeal.BiayaAdmin)
			if math.Abs(newTabungan.BiayaAdmin) > maxBiayaAdmin {
				maxBiayaAdmin = math.Abs(newTabungan.BiayaAdmin)
			}
		}
		if input.NilaiIdeal.BiayaPenarikanHabis == -1 {
			newTabungan.BiayaPenarikanHabis = float64(tabungan.BiayaPenarikanHabis)
			if float64(tabungan.BiayaPenarikanHabis) < minBiayaPenarikanHabis {
				minBiayaPenarikanHabis = float64(tabungan.BiayaPenarikanHabis)
			}
			if float64(tabungan.BiayaPenarikanHabis) > maxBiayaPenarikanHabis {
				maxBiayaPenarikanHabis = float64(tabungan.BiayaPenarikanHabis)
			}
		} else {
			newTabungan.BiayaPenarikanHabis = float64(tabungan.BiayaPenarikanHabis - input.NilaiIdeal.BiayaPenarikanHabis)
			if math.Abs(newTabungan.BiayaPenarikanHabis) > maxBiayaPenarikanHabis {
				maxBiayaPenarikanHabis = math.Abs(newTabungan.BiayaPenarikanHabis)
			}
		}

		if math.Abs(newTabungan.SetoranAwal) > maxSetoranAwal {
			maxSetoranAwal = math.Abs(newTabungan.SetoranAwal)
		}
		if math.Abs(newTabungan.SetoranLanjutanMinimal) > maxSetoranLanjutanMinimal {
			maxSetoranLanjutanMinimal = math.Abs(newTabungan.SetoranLanjutanMinimal)
		}
		if math.Abs(newTabungan.SaldoMinimum) > maxSaldoMinimum {
			maxSaldoMinimum = math.Abs(newTabungan.SaldoMinimum)
		}
		// if math.Abs(newTabungan.SukuBunga) > maxSukuBunga {
		// 	maxSukuBunga = math.Abs(newTabungan.SukuBunga)
		// }
		// if math.Abs(newTabungan.BiayaAdmin) > maxBiayaAdmin {
		// 	maxBiayaAdmin = math.Abs(newTabungan.BiayaAdmin)
		// }
		// if math.Abs(newTabungan.BiayaPenarikanHabis) > maxBiayaPenarikanHabis {
		// 	maxBiayaPenarikanHabis = math.Abs(newTabungan.BiayaPenarikanHabis)
		// }

		if tabungan.Fungsionalitas == "INVESTASI" {
			newTabungan.Fungsionalitas = input.NilaiIdeal.Fungsionalitas.Investasi
		} else if tabungan.Fungsionalitas == "BISNIS" {
			newTabungan.Fungsionalitas = input.NilaiIdeal.Fungsionalitas.Bisnis
		} else {
			newTabungan.Fungsionalitas = input.NilaiIdeal.Fungsionalitas.Transaksional
		}

		if tabungan.KategoriUmurPengguna == "DEWASA" {
			newTabungan.KategoriUmurPengguna = input.NilaiIdeal.KategoriUmurPengguna.Dewasa
		} else if tabungan.KategoriUmurPengguna == "REMAJA" {
			newTabungan.KategoriUmurPengguna = input.NilaiIdeal.KategoriUmurPengguna.Remaja
		} else {
			newTabungan.KategoriUmurPengguna = input.NilaiIdeal.KategoriUmurPengguna.Anak
		}

		newTabungans = append(newTabungans, newTabungan)
	}
	fmt.Println("isi dari maxSetoranAwal: ", maxSetoranAwal)
	fmt.Println("isi dari maxSetoranLanjutanMinimal: ", maxSetoranLanjutanMinimal)
	fmt.Println("isi dari maxSaldoMinimum: ", maxSaldoMinimum)
	fmt.Println("isi dari maxsukubunga: ", maxSukuBunga)
	fmt.Println("isi dari maxBiayaAdmin: ", maxBiayaAdmin)
	fmt.Println("isi dari maxbiayapenarikan habis: ", maxBiayaPenarikanHabis)
	fmt.Println("isi dari newtabungans: ", newTabungans)

	fmt.Println("isi dari minSukuBunga: ", minSukuBunga)
	fmt.Println("isi dari minBiayaAdmin: ", minBiayaAdmin)
	fmt.Println("isi dari minBiayaPenarikanHabis: ", minBiayaPenarikanHabis)

	// INTERPOLASI
	var newTabunganInterpolasi []NewTabungan
	for _, tabungan := range newTabungans {
		var newTabungan NewTabungan

		newTabungan.ID = tabungan.ID
		newTabungan.NamaTabungan = tabungan.NamaTabungan
		newTabungan.SetoranAwal = interpolasiLinear(tabungan.SetoranAwal, maxSetoranAwal, 5, 1)
		newTabungan.SetoranLanjutanMinimal = interpolasiLinear(tabungan.SetoranLanjutanMinimal, maxSetoranLanjutanMinimal, 5, 1)
		newTabungan.SaldoMinimum = interpolasiLinear(tabungan.SaldoMinimum, maxSaldoMinimum, 5, 1)
		if input.NilaiIdeal.SukuBunga == -1 {
			newTabungan.SukuBunga = interpolasiLinearBenefit(tabungan.SukuBunga, minSukuBunga, maxSukuBunga)
		} else {
			newTabungan.SukuBunga = interpolasiLinear(tabungan.SukuBunga, maxSukuBunga, 5, 1)
		}

		if input.NilaiIdeal.BiayaAdmin == -1 {
			newTabungan.BiayaAdmin = interpolasiLinearCost(tabungan.BiayaAdmin, minBiayaAdmin, maxBiayaAdmin)
		} else {
			newTabungan.BiayaAdmin = interpolasiLinear(tabungan.BiayaAdmin, maxBiayaAdmin, 5, 1)
		}

		if input.NilaiIdeal.BiayaPenarikanHabis == -1 {
			newTabungan.BiayaPenarikanHabis = interpolasiLinearCost(tabungan.BiayaPenarikanHabis, minBiayaPenarikanHabis, maxBiayaPenarikanHabis)
		} else {
			newTabungan.BiayaPenarikanHabis = interpolasiLinear(tabungan.BiayaPenarikanHabis, maxBiayaPenarikanHabis, 5, 1)
		}

		newTabungan.Fungsionalitas = tabungan.Fungsionalitas
		newTabungan.KategoriUmurPengguna = tabungan.KategoriUmurPengguna

		newTabunganInterpolasi = append(newTabunganInterpolasi, newTabungan)
	}

	fmt.Println("isi dari newtabunganinterpolasi : ", newTabunganInterpolasi)

	// preset, err := c.presetKriteriaService.FindFirstPreset(ctx.Request.Context())
	// if err != nil {
	// 	ctx.JSON(500, gin.H{
	// 		"code":  500,
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// var preset models.PresetKriteria

	// if err := ctx.ShouldBindJSON(&preset); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"code":  http.StatusBadRequest,
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	fmt.Println("isi dari input.bobot: ", input.BobotKriteria)

	// METODE SAW
	var hasilAkhir []HasilAkhir
	for _, tabungan := range newTabunganInterpolasi {
		var newHasilAkhir HasilAkhir

		newHasilAkhir.ID = tabungan.ID
		newHasilAkhir.NamaTabungan = tabungan.NamaTabungan
		newHasilAkhir.Skor = saw(tabungan, input.BobotKriteria)

		hasilAkhir = append(hasilAkhir, newHasilAkhir)
	}

	sort.Slice(hasilAkhir, func(i, j int) bool {
		return hasilAkhir[i].Skor > hasilAkhir[j].Skor
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    hasilAkhir,
	})
}

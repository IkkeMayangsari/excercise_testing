package main

import (
	"testing"
)

func TestHitungLuasPersegi(t *testing.T) {
	sisi := 4.0
	hasilYangDiharapkan := 16.0
	hasilHitung := HitungLuasPersegi(sisi)

	if hasilHitung != hasilYangDiharapkan {
		t.Errorf("Hasil hitung luas persegi salah. Hasil yang diharapkan: %f, Hasil: %f", hasilYangDiharapkan, hasilHitung)
	}
}

func TestHitungLuasPersegiPanjang(t *testing.T) {
	panjang := 5.0
	lebar := 7.0 
	hasilYangDiharapkan := 35.0
	hasilHitung := HitungLuasPersegiPanjang(panjang, lebar)

	if hasilHitung != hasilYangDiharapkan {
		t.Errorf("Hasil hitung luas persegi panjang salah. Hasil yang diharpkan: %f, Hasil: %f", hasilYangDiharapkan, hasilHitung)
	}
}
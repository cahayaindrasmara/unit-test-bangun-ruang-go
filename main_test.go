package main

import "testing"

var (
	kubus                    Kubus   = Kubus{4}
	volumeKubusSeharusnya    float64 = 64
	luasKubusSeharusnya      float64 = 96
	KelilingKubusSeharusnya  float64 = 48
	tabung                   Tabung  = Tabung{3, 4}
	volumeTabungSeharusnya   float64 = 113.04
	luasTabungSeharusnya     float64 = 131.88
	kelilingTabungSeharusnya float64 = 18.84
	balok                    Balok   = Balok{2, 3, 4}
	volumeBalokSeharusnya    float64 = 24
	luasBalokSeharusnya      float64 = 52
	kelilingBalokSeharusnya  float64 = 36
)

func TestHitungVolumeKubus(t *testing.T) {
	t.Logf("Volume: %.2f", kubus.Volume())

	if kubus.Volume() != volumeKubusSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", volumeKubusSeharusnya)
	}
}

func TestHitungLuasKubus(t *testing.T) {
	t.Logf("Luas: %.2f", kubus.Luas())

	if kubus.Luas() != luasKubusSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", luasKubusSeharusnya)
	}
}

func TestHitungKelilingKubus(t *testing.T) {
	t.Logf("Keliling: %.2f", kubus.Keliling())

	if kubus.Keliling() != KelilingKubusSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", KelilingKubusSeharusnya)
	}
}

func TestHitungVolumeTabung(t *testing.T) {
	t.Logf("Volume: %.2f", tabung.Volume())

	if tabung.Volume() != volumeTabungSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", volumeTabungSeharusnya)
	}
}

func TestHitungLuasTabung(t *testing.T) {
	t.Logf("Luas: %.2f", tabung.Luas())

	if tabung.Luas() != luasTabungSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", luasTabungSeharusnya)
	}
}

func TestHitungKelilingTabung(t *testing.T) {
	t.Logf("Keliling: %.2ff", tabung.Keliling())

	if tabung.Keliling() != kelilingTabungSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", kelilingTabungSeharusnya)
	}
}

func TestHitungVolumeBalok(t *testing.T) {
	t.Logf("Volume: %.2f", balok.Volume())

	if balok.Volume() != volumeBalokSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", volumeBalokSeharusnya)
	}
}

func TestHitungLuasBalok(t *testing.T) {
	t.Logf("Luas: %.2f", balok.Luas())

	if balok.Luas() != luasBalokSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", luasBalokSeharusnya)
	}
}

func TestHitungKelilingBalok(t *testing.T) {
	t.Logf("Luas: %.2f", balok.Keliling())

	if balok.Keliling() != kelilingBalokSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", kelilingBalokSeharusnya)
	}
}

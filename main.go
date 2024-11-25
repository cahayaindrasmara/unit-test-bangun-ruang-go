package main

import (
	"fmt"
	"math"
)

type Kubus struct {
	Sisi float64
}

type Tabung struct {
	Jarijari float64
	Tinggi   float64
}

type Balok struct {
	Panjang float64
	Lebar   float64
	Tinggi  float64
}

func (k Kubus) Volume() float64 {
	return math.Pow(k.Sisi, 3)
}

func (k Kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}

func (k Kubus) Keliling() float64 {
	return k.Sisi * 12
}

func (t Tabung) Volume() float64 {
	return 3.14 * math.Pow(t.Jarijari, 2) * t.Tinggi
}

func (t Tabung) Luas() float64 {
	return 2 * 3.14 * t.Jarijari * (t.Jarijari + t.Tinggi)
}

func (t Tabung) Keliling() float64 {
	return 2 * 3.14 * t.Jarijari
}

func (b Balok) Volume() float64 {
	return b.Panjang * b.Lebar * b.Tinggi
}

func (b Balok) Luas() float64 {
	return 2 * ((b.Panjang * b.Lebar) + (b.Panjang * b.Tinggi) + (b.Lebar * b.Tinggi))
}

func (b Balok) Keliling() float64 {
	return 4 * (b.Panjang + b.Lebar + b.Tinggi)
}

func main() {
	kubus := Kubus{Sisi: 5}
	fmt.Println(kubus.Luas())
}

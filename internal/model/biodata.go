package model

type Biodata struct {
	Nama string
	Foto string
	Email string
	Umur uint8
	Telepon string
	Pernikahan bool
	Pendidikan []Pendidikan
}

type Pendidikan struct {
	Nama string
	Jurusan string
}
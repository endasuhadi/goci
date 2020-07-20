package model

import (
	"app/conf"
)

type Siswa struct {
	No           int
	Nama         string `binding:"required"`
	NamaWali     string `json:"nama_wali", gorm:"column:nama_wali" binding:"required"`
	TanggalLahir string `json:"tanggal_lahir", gorm:"column:tanggal_lahir" binding:"required"`
	Telepon      string `binding:"required"`
	Alamat       string `binding:"required"`
}

type ResponseSiswa struct {
	ResponseData ResponseData
	Data         []Siswa
}

func GetSiswa(limit int) []Siswa{
	var siswa []Siswa
	db := conf.Db
	db.Limit(limit).Find(&siswa)
	return siswa
}

func SaveSiswa(siswa *Siswa) {
	db := conf.Db
	db.Create(siswa)
}

func UpdateSiswa(siswa *Siswa, no int) {
	db := conf.Db
	db.Table("siswa").Where("no = (?)", []int{no}).Update(siswa)
}

func DeleteSiswa(no string) {
	db := conf.Db
	db.Delete(Siswa{}, "no = ?", no)
}
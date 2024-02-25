package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (mh *MonHoc) TableName() string {
	return "monhoc"
}

func (dk *DangKi) TableName() string {
	return "dangki"
}

func (sv *SinhVien) TableName() string {
	return "sinhvien"
}

func (lop *Lop) TableName() string {
	return "lop"
}

type MonHoc struct {
	ID       int `gorm:"primaryKey"`
	TenMH    string
	SinhVien []SinhVien `gorm:"many2many:dangki"`
}

type DangKi struct {
	ID         int `gorm:"primaryKey"`
	SinhVienID int `gorm:"primaryKey"`
	MonHocID   int
	MonHoc     MonHoc
	SinhVien   SinhVien
}

type SinhVien struct {
	ID     int      `gorm:"primaryKey"`
	TenSV  string   `gorm:"column:ten_sv"`
	LopID  int      `gorm:"column:lop_id"`
	MonHoc []MonHoc `gorm:"many2many:dangki"`
}

type Lop struct {
	ID       uint `gorm:"primaryKey"`
	TenLop   string
	SinhVien []SinhVien
}

type LopSVReg struct {
	TenSV  string
	TenLop string
}

func main() {
	dsn := "host=postgresvm.postgres.database.azure.com user=mexd password=Saccar2108@ dbname=postgres port=5432 sslmode=require"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("Something went wrong")
	}

	// Hiển thị lớp sinh viên đang học
	var result []LopSVReg
	db.Table("sinhvien").Select("sinhvien.ten_sv, lop.ten_lop").Joins("JOIN lop ON sinhvien.lop_id = lop.id").Scan(&result)
	for _, r := range result {
		fmt.Printf("%+v\n", r)
	}

	// Hiển thị các môn sinh viên đã đăng kí
	var s []SinhVien
	db.Preload("MonHoc").Find(&s)
	for _, mh := range s {
		fmt.Printf("Tên sinh viên: %s\nCác môn đã đăng kí:\n", mh.TenSV)
		for _, r := range mh.MonHoc {
			fmt.Printf("%s\n", r.TenMH)
		}
	}
}

package main

import (
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"log"
)

type MonHoc struct {
	ID    int
	tenMH string
}

type DangKi struct {
	ID   int
	idSV int
	idMH int
}

type SinhVien struct {
	ID    int
	ten   string
	idLop int
}

type Lop struct {
	ID     int
	tenLop string
}

func main() {
	dsn := "host=postgresvm.postgres.database.azure.com user=mexd password=Saccar2108@ dbname=postgres port=5432 sslmode=require"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln("Something went wrong")
	}
}

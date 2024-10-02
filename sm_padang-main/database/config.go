package database

import (
	"log"

	"github.com/aldysp34/sm_padang/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabaseConfig() *gorm.DB {
	// dbURL := "postgres://postgres:postgres@localhost:5432/sm_padangreal"
    dbURL := "postgresql://postgres:keDmYjoYivszpziPhYiCLndFiWbpYriJ@autorack.proxy.rlwy.net:14267/railway"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.Barang{}, &model.BarangIn{}, &model.BarangOut{}, &model.Brand{}, &model.Request{}, &model.Role{}, &model.Satuan{}, &model.Supplier{}, &model.User{})

	return db

}

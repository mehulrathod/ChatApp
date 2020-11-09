package models

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var gormDB *gorm.DB

type Model struct {
	ID        uint64     `gorm:"primary_key;auto_increment:false" json:"id_msb" sql:"index"`
	CreatedAt time.Time  `json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type DefaultTimeStamp struct {
	CreatedAt time.Time  `json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type AutoIncrement struct {
	Id uint64 `gorm:"primary_key;" json:"id" sql:"index"`
}

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	schemaURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, dbHost, dbPort, dbName)
	fmt.Println(schemaURL)

	conn, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}

	gormCon, err := gorm.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}

	db = conn
	gormDB = gormCon
}

func GetDB() *sql.DB {
	return db
}

func GetGormDB() *gorm.DB {
	return gormDB.LogMode(true)
}

func AutoMigrate(conn *gorm.DB) {
	conn.Debug().AutoMigrate(
		// For auto migrate database tables, need to add models below

		&User{},
	)
}

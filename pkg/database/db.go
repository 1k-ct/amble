package database

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// DatabaseConnection
// DB_USER=
// DB_PASS=
// DB_ADDRESS=database.***.location.rds.amazomaws.com
// DB_NAME=
// DB_PORT=
// db, err := DatabaseConnection()
// 	if err != nil {
// 		c.JSON(500, appErrors.ErrMeatdataMsg(err, appErrors.ServerError))
// 		return
// 	}
// 	defer db.Close()
func DatabaseConnection() (*gorm.DB, error) {
	config := ConfigList{
		DbDriverName:   "mysql",
		DbName:         os.Getenv("DB_NAME"),
		DbUserName:     os.Getenv("DB_USER"),
		DbUserPassword: os.Getenv("DB_PASS"),
		DbHost:         os.Getenv("DB_ADDRESS"),
		DbPort:         os.Getenv("DB_PORT"),
	}
	PROTCOL := "@tcp(" + config.DbHost + ":" + config.DbPort + ")"
	CONNECT := config.DbUserName + ":" + config.DbUserPassword + PROTCOL + "/" + config.DbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(config.DbDriverName, CONNECT)
	if err != nil {
		return db, err
	}
	return db, nil
}

// ------------------------------------------------------

type Databaser interface {
	LocalDatabase() (*gorm.DB, error)
}

func NewData() Databaser { return &ConfigList{} }
func ConnectionDatabase(d Databaser) (*gorm.DB, error) {
	db, err := d.LocalDatabase()
	if err != nil {
		return nil, err
	}
	return db, nil
}
func (c *ConfigList) LocalDatabase() (*gorm.DB, error) {
	db, err := c.Connect()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ------------------------------------------------------
type ConfigList struct {
	DbDriverName   string
	DbName         string
	DbUserName     string
	DbUserPassword string
	DbHost         string
	DbPort         string
	ServerPort     string
}

func NewLocalDB(user, pass, dbName string) (*ConfigList, error) {
	config := &ConfigList{
		DbDriverName:   "mysql",
		DbName:         dbName,
		DbUserName:     user,
		DbUserPassword: pass,
		DbHost:         "127.0.0.1",
		DbPort:         "3306",
		ServerPort:     "8080",
	}

	return config, nil
}
func NewDB(file string) (*ConfigList, error) {
	if err := godotenv.Load(file); err != nil {
		return nil, err
	}
	config := &ConfigList{
		DbDriverName:   os.Getenv("DRIVER_NAME"),
		DbName:         os.Getenv("DB_NAME"),
		DbUserName:     os.Getenv("DB_USER"),
		DbUserPassword: os.Getenv("DB_PASSWORD"),
		DbHost:         os.Getenv("DB_HOST"),
		DbPort:         os.Getenv("DB_PORT"),
		ServerPort:     os.Getenv("PORT"),
	}
	return config, nil
}
func (c *ConfigList) Connect() (*gorm.DB, error) {
	// _ "github.com/go-sql-driver/mysql"

	DBMS := c.DbDriverName
	USER := c.DbUserName
	PASS := c.DbUserPassword
	PROTOCOL := "tcp(" + c.DbHost + ":" + c.DbPort + ")"
	DBNAME := c.DbName

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		return nil, err
	}

	return db, nil
}
func (c *ConfigList) NewMakeDB(tables ...interface{}) error {
	db, err := c.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	db.AutoMigrate(tables...)
	return nil
}

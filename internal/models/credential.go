package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Credential struct {
	gorm.Model
	Name        string `gorm:"index;column:name;unique"`
	Login       string `gorm:"column:login"`
	Password    string `gorm:"column:password"`
	Description string `gorm:"column:description"`
}

func databasePath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.passager.db", dirname)
}

func openConnection() *gorm.DB {
	dbPath := databasePath()
	conn, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}
	conn.AutoMigrate(&Credential{})
	return conn
}

func Create(cred *Credential) error {
	conn := openConnection()
	return conn.Create(cred).Error
}

func ReadAll(name string) ([]Credential, error) {
	conn := openConnection()
	var credentials []Credential
	err := conn.Find(&credentials, "LOWER(name) LIKE LOWER(?)", fmt.Sprintf("%%%s%%", name)).Error
	return credentials, err
}

func Delete(name string) error {
	conn := openConnection()
	return conn.Delete(&Credential{}, "name", name).Error
}

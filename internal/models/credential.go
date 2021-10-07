package models

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Credential struct {
	gorm.Model
	Name     string `gorm:"index;column:name;unique"`
	Login    string `gorm:"column:login"`
	Password string `gorm:"column:password"`
}

func (c Credential) String() string {
	return fmt.Sprintf("| %-24s| %-24s| %-24s|", c.Name, c.Login, c.Password)
}

func openConnection() *gorm.DB {
	dbPath := "test.db"
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

func Read(name string) (Credential, error) {
	conn := openConnection()
	var credential Credential
	err := conn.First(&credential, "name = ?", name).Error
	return credential, err
}

func ReadAll() ([]Credential, error) {
	conn := openConnection()
	var credentials []Credential
	err := conn.Find(&credentials).Error
	return credentials, err
}

func Delete(name string) error {
	conn := openConnection()
	return conn.Delete(&Credential{}, "name", name).Error
}

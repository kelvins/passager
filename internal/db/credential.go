package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Credential struct {
	gorm.Model
	Name     string `gorm:"index"`
	Login    string
	Password string
}

func openConnection() *gorm.DB {
	conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	conn.AutoMigrate(&Credential{})
	return conn
}

func Create(cred *Credential) {
	fmt.Println("Create")
	conn := openConnection()
	conn.Create(cred)
}

func Read(name string) (Credential, error) {
	fmt.Println("Read")
	conn := openConnection()
	var credential Credential
	err := conn.First(&credential, "name = ?", name).Error
	return credential, err
}

func Delete(name string) {
	fmt.Println("Delete")
}

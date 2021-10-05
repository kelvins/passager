package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Credential struct {
	gorm.Model
	Name     string `gorm:"index" gorm:"column:name"`
	Login    string `gorm:"column:login"`
	Password string `gorm:"column:password"`
}

func (c Credential) String() string {
	return fmt.Sprintf("| %-24s| %-24s| %-24s|", c.Name, c.Login, c.Password)
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

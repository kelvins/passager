package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Credential is the representation of a specific credential in the database.
type Credential struct {
	gorm.Model
	Name        string `gorm:"index;column:name;unique"`
	Login       string `gorm:"column:login"`
	Password    string `gorm:"column:password"`
	Description string `gorm:"column:description"`
}

func databasePath() string {
	if path := os.Getenv("PASSAGER_DATABASE"); path != "" {
		return path
	}
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

// Create creates a new credential in the database.
func Create(cred *Credential) error {
	conn := openConnection()
	return conn.Create(cred).Error
}

// Save saves an existing credential in the database.
func Save(cred *Credential) error {
	conn := openConnection()
	return conn.Save(&cred).Error
}

// ReadAll reads all credentials from the database that matches the provided name (using LIKE).
func ReadAll(name string) ([]Credential, error) {
	conn := openConnection()
	var credentials []Credential
	err := conn.Order("name").Find(&credentials, "LOWER(name) LIKE LOWER(?)", fmt.Sprintf("%%%s%%", name)).Error
	return credentials, err
}

// ReadFirst reads the first credential from the database that matches the provided name.
func ReadFirst(name string) (Credential, error) {
	conn := openConnection()
	var credential Credential
	err := conn.Where("name = ?", name).First(&credential).Error
	return credential, err
}

// Delete deletes a credential from the database based on the provided name.
func Delete(name string) error {
	if _, err := ReadFirst(name); err != nil {
		return err
	}
	conn := openConnection()
	return conn.Unscoped().Delete(&Credential{}, "name", name).Error
}

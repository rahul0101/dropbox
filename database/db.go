package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/rahularcota/dropbox/config"
	"sync"
)

var once sync.Once
var dbConn *gorm.DB

func NewPostgresConnection(c *config.DatabaseConfig) (*gorm.DB, error) {
	once.Do(func() {
		connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.IP, c.Port, c.Username, c.Name, c.Password)
		var err error
		dbConn, err = gorm.Open("postgres", connStr)
		if err != nil {
			fmt.Printf("failed to connect database with error: %s", err.Error())
			panic(nil)
		}
	})
	return dbConn, nil
}

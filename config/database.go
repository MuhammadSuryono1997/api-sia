package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func (dbConfig *DBConfig) GetConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

var GormDb *gorm.DB

func Init(config DBConfig) {
	db, err := gorm.Open("mysql", config.GetConnectionString())
	if err != nil {
		fmt.Println("error ", err)
	}
	GormDb = db
	//GormDb.DB().SetMaxOpenConns(5)
	//GormDb.DB().SetMaxIdleConns(5)
	//GormDb.DB().SetConnMaxLifetime(5 * time.Minute)
	pingTicker := time.NewTicker(15 * time.Second)
	pingDone := make(chan bool)
	go func() {
		for {
			select {
			case <-pingDone:
				return
			case <-pingTicker.C:
				b := pingDb(GormDb.DB())
				if !b {
					pingDone <- true
				}
			}
		}
	}()

}

func GetDb() *gorm.DB {
	return GormDb
}

func pingDb(db *sql.DB) bool {
	er := db.Ping()
	if er != nil {
		log.Print("mysql error ping", er)
		return false
	} else {
		log.Print("mysql success ping")
		return true
	}
}

package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var PG *gorm.DB

func NewPG() *gorm.DB {
	if PG == nil {

		dsn := "host=localhost user=postgres password=postgres dbname=postgres port=9920 sslmode=disable TimeZone=Asia/Shanghai"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// use singular table name, table for `User` would be `user` with this option enabled
				SingularTable: true,
			},
		})
		if err != nil {
			panic("PG db connect fail")
		}
		sqlDB, err := db.DB()
		if err != nil {
			panic(fmt.Errorf("PG db err when get DB %s", err))
		}

		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
		PG = db
	}
	return PG
}

package db

import (
	"fmt"
	"time"

	"github.com/fengshux/blog2/backend/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var PG *gorm.DB

func NewPG(config *conf.Conf) *gorm.DB {

	if PG == nil {

		c := config.Postgres
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			c.Host,
			c.User,
			c.Password,
			c.DB,
			c.Port,
		)

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

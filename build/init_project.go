package build

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgresConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

func loadDBConfig() *PostgresConf {
	// create and with some options
	initConfig := config.NewWithOptions("init-conf", config.ParseEnv)
	initConfig.AddDriver(yaml.Driver)

	configName := "config/config.yaml"

	err := initConfig.LoadFiles(configName)
	if err != nil {
		panic(err)
	}

	data := PostgresConf{}
	err = initConfig.BindStruct("postgres", &data)
	if err != nil {
		panic(err)
	}
	log.Printf("load config form :%s data:%+v\n", configName, data)
	return &data

}

func InitProject() {
	log.Println("check init start")
	c := loadDBConfig()

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
		panic(err)
	}

	result := db.Exec("SELECT * FROM pg_tables WHERE schemaname = 'public' AND tablename  = 'user';")

	// 如果不报错，就不用初始化
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(result.Error)
	}

	if result.RowsAffected != 0 {
		log.Printf("no need to init")
		return
	}

	log.Println("init database start")
	query, err := ioutil.ReadFile("build/sql/blog.sql")
	if err != nil {
		panic(err)
	}
	result = db.Exec(string(query))
	if result.Error != nil {
		panic(result.Error)
	}

	log.Println("init database finished")
}

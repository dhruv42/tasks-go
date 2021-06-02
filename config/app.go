package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	db *gorm.DB
)

func loadConfig() {
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}

func dsn() string {

	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		viper.GetString("database.dbuser"),
		viper.GetString("database.dbpassword"),
		viper.GetString("database.dbhost"),
		viper.GetString("database.dbname"))
}

func DbOpen() {
	loadConfig()
	d, err := gorm.Open("mysql", dsn())
	if err != nil {
		panic(err)
	}
	db = d
	fmt.Println("=========== Database connected successfully ===========")
}

func GetDB() *gorm.DB {
	return db
}

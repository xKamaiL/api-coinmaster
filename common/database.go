package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	dbName := viper.GetString("mysql.connection.database_name")
	dbHost := viper.GetString("mysql.connection.host")
	dbUsername := viper.GetString("mysql.connection.username")
	dbPassword := viper.GetString("mysql.connection.password")

	db, err := gorm.Open("mysql", dbUsername+":"+dbPassword+"@("+dbHost+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Database err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

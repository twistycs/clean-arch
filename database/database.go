package database

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

//SetUpDb Database
func SetUpDb() *DBConfig {
	start := time.Now()
	fmt.Println("==========Start Read Config DB==========")
	viper.SetConfigFile("./dbconfig.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Can not read config file.")
	}
	var configDBs DBConfig
	viper.Unmarshal(&configDBs)
	fmt.Println("==========End Read Config DB==========")
	end := time.Now()
	fmt.Printf("Use Time for : %s", end.Sub(start))
	return &configDBs
}

//ReadConfigDBFile print
func (db *DBConfig) ReadConfigDBFile() string {
	fmt.Printf("%+v", db)
	return fmt.Sprintf(db.User + ":" + db.Password + "@tcp(" + db.Host + ")/" + db.DBName + "?charset=utf8&parseTime=True&loc=Local")
}

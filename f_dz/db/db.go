package db

import (
	"f_dz/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "user"
	Password = "password"
	Dbname   = "testdb"
)

var DB *gorm.DB

func Connect(params configs.Database) {
	logDbs := make(map[string]interface{})
	logDbs["ConnectDB"] = "Connect() - Connection Opened to Database"
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", params.Host, params.Port, params.User, params.Password, params.DBName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logDbs["ConnectDB"] = "Failed to connect to database"
		log.Fatal("Failed to connect to database\n", err)
		logs.LogConfigsParams("dbs", logDbs)
		os.Exit(2)
	}

	logs.LogConfigsParams("dbs", logDbs)
	log.Println("Connect()- ok, Connection Opened to Database")
}

func Migrate() {
	logDbs := make(map[string]interface{})
	logDbs["MigrateDB"] = "Model USERS Migration Completed"
	DB.AutoMigrate(&models.Users{})
	logs.LogConfigsParams("dbs", logDbs)
	log.Println("Model USERS Migration Completed")
	DB.AutoMigrate(&models.Transactions{})
	logDbs["MigrateDB"] = "Model TRANSACTIONS Migration Completed"
	logs.LogConfigsParams("dbs", logDbs)
	log.Println("Migrate(): Model Transactions Migration Completed")
}

func Close() {
	logDbs := make(map[string]interface{})
	logDbs["CloseDB"] = "Database CLOSE Completed"
	sqlDB, err := DB.DB()
	if err != nil {
		logDbs["CloseDB"] = "Failed to close database connection"
		log.Fatal("Failed to close database connection\n", err)
	}
	sqlDB.Close()
	logs.LogConfigsParams("dbs", logDbs)
	log.Println("Close() - Database CLOSE Completed")
}

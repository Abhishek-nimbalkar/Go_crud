package db

import (
	"fmt"
	"strconv"

	"example.com/crud/configs/env"
	"example.com/crud/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
// ConnectDatabase creates a new database connection
func ConnectDatabase(envConfig *env.Env) (*gorm.DB, error) {
	host := envConfig.DbHost
	port, _ := strconv.Atoi(envConfig.DbPort) // don't forget to convert int since port is int type.
	user := envConfig.DbUser
	dbname := envConfig.DbName
	pass := envConfig.DbPassword
	// set up postgres sql to open it.
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)
	// dsn := "user=postgres password=postgres dbname=golang host=localhost port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(psqlSetup), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Optionally, you can perform database migrations here
	db.AutoMigrate(&models.User{})
	DB = db
	return DB, nil
}

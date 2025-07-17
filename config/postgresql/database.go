package postgresql

import (
	"fmt"
	"os"
	"strconv"
	"time"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConnection struct {
	Db *gorm.DB
}

func CreateConnection() *DbConnection {
	var db *gorm.DB
	var err error

	i := 0
	for {
		db, err = gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			if i == 0 {
				fmt.Printf("CONNECT %v : %v \n", err, os.Getenv("CONNECTION_STRING"))
			} else {
				fmt.Printf("RECCONECT(%d) %v : %v \n", i, err, os.Getenv("CONNECTION_STRING"))
			}
			time.Sleep(3 * time.Second)
			i++
			continue
		}
		break
	}

	num, _ := strconv.Atoi(os.Getenv("MAX_CONNECTION_POOL"))
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("FAILED TO CONNECT DB")
	}
	maxConnIdle, err := strconv.Atoi(os.Getenv("MAX_CONNECTION_IDLE"))
	if err != nil {
		maxConnIdle = 5
	}
	dbSQL.SetConnMaxIdleTime(time.Duration(maxConnIdle))
	dbSQL.SetMaxOpenConns(num)
	dbSQL.SetConnMaxLifetime(time.Hour)
	return &DbConnection{db}
}

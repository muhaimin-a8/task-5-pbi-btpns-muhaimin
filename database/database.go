package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"
	"pbi-btpns-api/app"
	"time"
)

func init() {
	app.LoadConfig()
}

func NewDB() (*sql.DB, error) {
	var host string
	var port int
	var dbname string
	var user string
	var pass string
	var sslmode string

	if e := os.Getenv("ENV"); e == "production" {
		// db config for production
		host = viper.Get("database.postgres.host").(string)
		port = viper.Get("database.postgres.port").(int)
		dbname = viper.Get("database.postgres.dbname").(string)
		user = viper.Get("database.postgres.user").(string)
		pass = viper.Get("database.postgres.pass").(string)
		sslmode = viper.Get("database.postgres.sslmode").(string)
	} else {
		// db config for test
		host = viper.Get("database.postgresTest.host").(string)
		port = viper.Get("database.postgresTest.port").(int)
		dbname = viper.Get("database.postgresTest.dbname").(string)
		user = viper.Get("database.postgresTest.user").(string)
		pass = viper.Get("database.postgresTest.pass").(string)
		sslmode = viper.Get("database.postgresTest.sslmode").(string)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, pass, sslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// configure sql pool
	maxIdle := viper.Get("database.pool.maxIdle").(int)
	maxOpen := viper.Get("database.pool.maxOpen").(int)
	maxIdleTime := viper.Get("database.pool.maxIdleTime").(int)
	maxLifeTime := viper.Get("database.pool.maxLifeTime").(int)

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)
	db.SetConnMaxIdleTime(time.Second * time.Duration(maxIdleTime))
	db.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTime))

	return db, nil
}

package config

import (
	"os"

	"upper.io/db.v3/mysql"
)

var (
	host     = "127.0.0.1:3307"
	database = "secret_santa_dev"
	user     = "root"
	password = "password"
)

func init() {
	if os.Getenv("MYSQL_HOST") != "" {
		host = os.Getenv("MYSQL_HOST")
	}
	if os.Getenv("MYSQL_DATABASE") != "" {
		database = os.Getenv("MYSQL_DATABASE")
	}
	if os.Getenv("MYSQL_USER") != "" {
		user = os.Getenv("MYSQL_USER")
	}
	if os.Getenv("MYSQL_PASSWORD") != "" {
		password = os.Getenv("MYSQL_PASSWORD")
	}
}

func GetDbSettings() mysql.ConnectionURL {
	return mysql.ConnectionURL{
		Database: database,
		Host:     host,
		User:     user,
		Password: password,
	}
}

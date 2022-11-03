package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var JWT_SECRET_KEY = ""
var JWT_EXPIRE_DURATION time.Duration
var DB_USERNAME = ""
var DB_PASSWORD = ""
var DB_HOST = ""
var DB_PORT = ""
var DB_NAME = ""

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	duration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_DURATION"))
	JWT_EXPIRE_DURATION = time.Hour * time.Duration(duration)
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
}

package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT   = 0
	DBURL  = ""
	DBNAME = ""
)

func Load() {

	var err error
	err = godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 9999
	}

	DBNAME = os.Getenv("DB_NAME")
	DBURL = fmt.Sprintf("%s://%s:%s", os.Getenv("DB_DRIVER"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
}

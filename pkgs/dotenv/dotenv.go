package dotenv

import (
	"log"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

var mySqlUrl string

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetMySQL() string {
	if mySqlUrl == "" {
		mySqlUrl = goDotEnvVariable("MYSQL")
	}
	return mySqlUrl
}

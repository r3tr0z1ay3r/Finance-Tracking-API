package Config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(path string) {

	if err := godotenv.Load(path); err != nil {

		log.Fatal("No .env file")

	}
}

func GetEnv(key string) string {

	return os.Getenv(key)

}

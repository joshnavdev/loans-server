package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshnavdev/loans-server/utils"
)

type AppConfig struct {
  Port string
  Host string
  Addr string
} 

type DatabaseConfig struct {
  User string
  Pass string
  Host string
  Name string
}

type Config struct {
  Application AppConfig
  Database DatabaseConfig
}

func (db *DatabaseConfig) GetUri() string {
  uri := fmt.Sprintf(
    "mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
    db.User,
    db.Pass,
    db.Host,
  )

  return uri
}

func LoadConfig() (config Config, err error) {
  err = godotenv.Load()

  if err != nil {
    return
  }

  port := utils.GetValue(os.Getenv("PORT"), "3000")
  host := utils.GetValue(os.Getenv("HOST"), "127.0.0.1")
  addr := fmt.Sprintf("%s:%s", host, port)

  dbUser := utils.GetValue(os.Getenv("DB_USER"), "root")
  dbPass := os.Getenv("DB_PASS")
  dbHost := os.Getenv("DB_HOST")
  dbName := os.Getenv("DB_NAME")

 
  config = Config{
    Application: AppConfig{
      Port: port,
      Host: host,
      Addr: addr,
    },
    Database: DatabaseConfig {
      User: dbUser,
      Pass: dbPass,
      Host: dbHost,
      Name: dbName,
    },
  }

  return
}

func GetEnv() Config {
  config, err := LoadConfig()

  if err != nil {
    log.Fatal("Cannot load env", err)
  }

  return config
}

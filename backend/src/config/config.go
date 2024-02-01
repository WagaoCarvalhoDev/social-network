package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	apiPortEnvVar       = "API_PORT"
	mysqlUserEnvVar     = "MYSQL_USER"
	mysqlPasswordEnvVar = "MYSQL_PASSWORD"
	mysqlDatabaseEnvVar = "MYSQL_DATABASE"
)

var (
	Port          int
	ServerAddress string
)

func SetupEnvironment() error {
	var err error

	// Carrega as variáveis de ambiente do arquivo .env no diretório atual
	if err = godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load environment variables: %v", err)
	}

	Port, err = strconv.Atoi(os.Getenv(apiPortEnvVar))
	if err != nil {
		Port = 9000
	}

	ServerAddress = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv(mysqlUserEnvVar),
		os.Getenv(mysqlPasswordEnvVar),
		os.Getenv(mysqlDatabaseEnvVar),
	)

	return nil
}

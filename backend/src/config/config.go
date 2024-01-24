package config

import (
	"fmt"
	"os"
	"path/filepath"
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

	// Obtém o diretório pai do diretório atual
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %v", err)
	}

	parentDir := filepath.Dir(currentDir)

	// Monta o caminho para o arquivo .env no diretório pai
	envFilePath := filepath.Join(parentDir, ".env")

	// Carrega as variáveis de ambiente do arquivo .env
	if err = godotenv.Load(envFilePath); err != nil {
		return fmt.Errorf("failed to load environment variables from %s: %v", envFilePath, err)
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

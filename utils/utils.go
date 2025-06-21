package utils

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

var (
	AppName		string
	AppHost		string
	AppPort		int

	AppStatus	bool

	DBHost		string
	DBPort		int
	DBUser		string
	DBPassword	string
	DBName		string
	TimeZone	string
)

const (
	defaultAppName		= "Go Fiber"
	defaultAppHost		= "127.0.0.1"
	defaultAppPort		= 8000

	defaultAppStatus	= false

	defaultDBHost		= "127.0.0.1"
	defaultDBPort		= 5432
	defaultDBUser		= "postgres"
	defaultDBPassword	= "postgres"
	defaultDBName		= "postgres"
	defaultTimeZone		= "Asia/Jakarta"
)

func init() {
	// Configurations File
	viper.SetConfigName(".env.example")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")

	// setdefault App
	viper.SetDefault("APP_NAME", defaultAppName)
	viper.SetDefault("APP_HOST", defaultAppHost)
	viper.SetDefault("APP_PORT", defaultAppPort)
	viper.SetDefault("APP_STATUS",defaultAppStatus)

	// setdefault Database
	viper.SetDefault("DB_HOST", defaultDBHost)
	viper.SetDefault("DB_PORT", defaultDBPort)
	viper.SetDefault("DB_USER", defaultDBUser)
	viper.SetDefault("DB_PASSWORD", defaultDBPassword)
	viper.SetDefault("DB_NAME", defaultDBName)
	viper.SetDefault("TimeZone", defaultTimeZone)

	err := viper.ReadInConfig()
	if err != nil {
		var configNotFound viper.ConfigFileNotFoundError
		if errors.As(err, &configNotFound) {
			log.Println("File Not Found")
		} else {
			log.Printf("Error Reading File %s", err)
		}
	} else {
		log.Println("Using Config File : ", viper.ConfigFileUsed())
	}

	// Configurations App
	AppName = viper.GetString("APP_NAME")
	AppHost = viper.GetString("APP_HOST")
	AppPort = viper.GetInt("APP_PORT")
	AppStatus = viper.GetBool("APP_STATUS")

	// Configurations Database
	DBHost = viper.GetString("DB_HOST")
	DBPort = viper.GetInt("DB_PORT")
	DBUser = viper.GetString("DB_USER")
	DBPassword = viper.GetString("DB_PASSWORD")
	TimeZone = viper.GetString("TimeZone")
}

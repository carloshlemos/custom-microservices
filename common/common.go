package common

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type Configuration struct {
	Port                string `json:"port"`
	EnableGinConsoleLog bool   `json:"enableGinConsoleLog"`
	EnableGinFileLog    bool   `json:"enableGinFileLog"`

	LogFilename   string `json:"logFilename"`
	LogMaxSize    int    `json:"logMaxSize"`
	LogMaxBackups int    `json:"logMaxBackups"`
	LogMaxAge     int    `json:"logMaxAge"`

	PgAddrs      string `json:"pgAddrs"`
	PgPort       string `json:"pgPort"`
	PgDbName     string `json:"pgDbName"`
	PgDbUsername string `json:"pgDbUsername"`
	PgDbPassword string `json:"pgDbPassword"`
}

// Config shares the global configuration
var (
	Config *Configuration
)

// Status Code
const (
	StatusCodeUnknown = -1
	StatusCodeOK      = 1000

	StatusMismatch = 10
)

func LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("./config/config.json")
	if err != nil {
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	// Setting Service Logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   Config.LogFilename,
		MaxSize:    Config.LogMaxSize,    // megabytes after which new file is created
		MaxBackups: Config.LogMaxBackups, // number of backups
		MaxAge:     Config.LogMaxAge,     // days
	})
	log.SetLevel(log.DebugLevel)

	// log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}

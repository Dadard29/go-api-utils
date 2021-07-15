package database

import (
	"fmt"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"reflect"
	"time"
)

func NewConnector(configMap map[string]string, verbose bool, modelList []interface{}) *Connector {

	dbConfig := DbConfig{
		usernameKey:  configMap["usernameKey"],
		passwordKey:  configMap["passwordKey"],
		databaseName: configMap["database"],
		host:         configMap["host"],
		port:         configMap["port"],
	}

	loggerName := fmt.Sprintf("%s_connector", dbConfig.databaseName)
	logger := log.NewLogger(loggerName, logLevel.LevelFromBool(verbose))

	usernameValue := os.Getenv(dbConfig.usernameKey)
	passwordValue := os.Getenv(dbConfig.passwordKey)

	// added so gorm can parse the date object in the sql format
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	parseTime := "?parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", usernameValue, passwordValue,
		dbConfig.host, dbConfig.port, dbConfig.databaseName, parseTime)

	logger.Debug(fmt.Sprintf("connecting to %s...", dbConfig.databaseName))
	var err error
	var db *gorm.DB
	for err != nil {
		db, err = gorm.Open("mysql", dsn)
		if err != nil {
			logger.Error("failed to connect to database")
			time.Sleep(15 * time.Second)
		}
	}

	if db == nil {
		logger.Fatal("failed to setup DB connection")
		return nil
	}

	logger.CheckErrFatal(err)
	logger.Info(fmt.Sprintf("connected to %s...", dbConfig.databaseName))

	for _, v := range modelList {
		if !db.HasTable(v) {
			msg := fmt.Sprintf("Model %v does not have existing table\n", reflect.TypeOf(v))
			logger.Warning(msg)
		}
	}

	return &Connector{
		Orm:      db,
		dbConfig: dbConfig,
		logger:   logger,
	}
}

package dao

import (
	"os"
	"reflect"
	"strings"

	storm "github.com/asdine/storm/v3"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var name = "DAO"
var DBVersion = 1
var DB *storm.DB
var DBName string = "default"

func Initialise(cfg *commonConfig.Settings) error {
	clock := timing.Start(name, "Initialise", "")
	logHandler.InfoLogger.Printf("[%v] Initialising...", strings.ToUpper(name))

	DBVersion = cfg.GetDatabaseVersion()
	DBName = cfg.GetDatabaseName()

	logHandler.InfoLogger.Printf("[%v] Initialised", strings.ToUpper(name))
	clock.Stop(1)
	return nil
}

func GetDBNameFromPath(t string) string {
	dbName := t
	// split dbName on "/"
	dbNameArr := strings.Split(dbName, string(os.PathSeparator))
	noparts := len(dbNameArr)
	dbName = dbNameArr[noparts-1]
	logHandler.InfoLogger.Printf("dbName: %v\n", dbName)
	return dbName
}

func IsValidFieldInStruct(fromField string, data any) error {
	_, isValidField := reflect.TypeOf(data).FieldByName(fromField)
	if !isValidField {
		logHandler.ErrorLogger.Panic(commonErrors.WrapInvalidFieldError(fromField))
		return commonErrors.WrapInvalidFieldError(fromField)
	}
	return nil
}

func CheckDAOReadyState(table string, action audit.Action, isDaoReady bool) {
	if !isDaoReady {
		err := commonErrors.WrapDAONotInitialisedError(table, action.Description())
		logHandler.ErrorLogger.Panic(err)
	}
}

func GetStructType(data any) string {
	rtnType := reflect.TypeOf(data).String()
	// If the type is a pointer, get the underlying type
	if strings.Contains(rtnType, "*") {
		rtnType = reflect.TypeOf(data).Elem().String()
	}
	// If the type is a struct, get the name of the struct
	if strings.Contains(rtnType, ".") {
		rtnType = strings.Split(rtnType, ".")[1]
	}
	return rtnType
}

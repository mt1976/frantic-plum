package database

import (
	"os"
	"strings"

	storm "github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/index"
	validator "github.com/go-playground/validator/v10"
	"github.com/mt1976/frantic-plum/commonErrors"
	"github.com/mt1976/frantic-plum/io"
	"github.com/mt1976/frantic-plum/logger"
	stopwatch "github.com/mt1976/frantic-plum/timing"
)

var Version = 1
var CONNECTION *storm.DB
var domain = "database"

var dataValidator *validator.Validate

func init() {
	Connect()
	dataValidator = validator.New(validator.WithRequiredStructEnabled())
}

func Connect() {
	connect := stopwatch.Start(domain, "OpenDatabaseConnection", "")
	var err error
	CONNECTION, err = storm.Open(io.GetDBFileName(domain), storm.BoltOptions(0666, nil))
	if err != nil {

		logger.ErrorLogger.Printf("[%v] Opening [%v.db] connection Error=[%v]", strings.ToUpper(domain), strings.ToLower(domain), err.Error())
		os.Exit(1)
	}
	logger.DatabaseLogger.Printf("[%v] Open [%v.db] data connection", strings.ToUpper(domain), domain)
	connect.Stop(1)
}

func Backup(loc string) {
	timer := stopwatch.Start(domain, "BackupDatabase", "")
	logger.EventLogger.Printf("[BACKUP] Backup [%v.db] data started...", domain)
	Disconnect()
	io.Backup(domain, loc)
	Connect()
	logger.EventLogger.Printf("[BACKUP] Backup [%v.db] data ends", domain)
	timer.Stop(1)
	logger.DatabaseLogger.Printf("[%v] Backup [%v.db] data connection", strings.ToUpper(domain), domain)
}

func Disconnect() {
	logger.EventLogger.Printf("[%v] Close [%v.db] data file", strings.ToUpper(domain), domain)
	err := CONNECTION.Close()
	if err != nil {
		logger.ErrorLogger.Printf("[%v] Closing %e ", strings.ToUpper(domain), err)
		panic(err)
	}
	logger.DatabaseLogger.Printf("[%v] Close [%v.db] data connection", strings.ToUpper(domain), domain)
}

func Retrieve(fieldName string, value, to any) error {
	logger.DatabaseLogger.Printf("Retrieve [%+v][%+v][%+v]", fieldName, value, to)
	return CONNECTION.One(fieldName, value, to)
}

func GetAll(to any, options ...func(*index.Options)) error {
	logger.DatabaseLogger.Printf("GetAll [%+v][%+v]", to, options)
	return CONNECTION.All(to, options...)
}

func Delete(data any) error {
	logger.DatabaseLogger.Printf("Delete [%+v]", data)
	return CONNECTION.DeleteStruct(data)
}

func Drop(data any) error {
	logger.DatabaseLogger.Printf("Drop [%+v]", data)
	return CONNECTION.Drop(data)
}

func Update(data any) error {
	err := validate(data)
	if err != nil {
		return err
	}
	logger.DatabaseLogger.Printf("Update [%+v]", data)
	return CONNECTION.Update(data)
}

func Create(data any) error {
	err := validate(data)
	if err != nil {
		return err
	}
	logger.DatabaseLogger.Printf("Create [%+v]", data)
	return CONNECTION.Save(data)
}

func validate(data any) error {
	err := commonErrors.HandleGoValidatorError(dataValidator.Struct(data))
	if err != nil {
		logger.ErrorLogger.Printf("[%v] Validation  %v", strings.ToUpper(domain), err.Error())
		return err
	}
	return nil
}

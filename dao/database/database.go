package database

import (
	"strings"

	storm "github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/index"
	validator "github.com/go-playground/validator/v10"
	"github.com/mt1976/frantic-plum/commonErrors"
	"github.com/mt1976/frantic-plum/io"
	"github.com/mt1976/frantic-plum/logger"
	"github.com/mt1976/frantic-plum/timing"
)

var Version = 1
var CONNECTION *storm.DB
var domain = "database"
var dbFileName string

var dataValidator *validator.Validate

func init() {
	//	Connect()
	dataValidator = validator.New(validator.WithRequiredStructEnabled())
}

func Connect() {
	connect(domain)
}

func NamedConnect(name string) {
	connect(name)
}

func connect(name string) {
	dbFileName = name
	connect := timing.Start(domain, "Connect", "")
	var err error
	CONNECTION, err = storm.Open(io.GetDBFileName(dbFileName), storm.BoltOptions(0666, nil))
	if err != nil {

		logger.ErrorLogger.Printf("[%v] Opening [%v.db] connection Error=[%v]", strings.ToUpper(domain), strings.ToLower(dbFileName), err.Error())
		panic(commonErrors.ConnectError(err))
		//os.Exit(1)
	}
	logger.DatabaseLogger.Printf("[%v] Open [%v.db] data connection", strings.ToUpper(domain), dbFileName)
	connect.Stop(1)
}

func Backup(loc string) {
	timer := timing.Start(domain, "Backup", dbFileName)
	logger.EventLogger.Printf("[BACKUP] Backup [%v.db] data started...", dbFileName)
	Disconnect()
	io.Backup(dbFileName, loc)
	Connect()
	logger.EventLogger.Printf("[BACKUP] Backup [%v.db] data ends", dbFileName)
	timer.Stop(1)
	logger.DatabaseLogger.Printf("[%v] Backup [%v.db] data connection", strings.ToUpper(domain), dbFileName)
}

func Disconnect() {
	timer := timing.Start(domain, "Disconnect", dbFileName)
	logger.EventLogger.Printf("[%v] Close [%v.db] data file", strings.ToUpper(domain), dbFileName)
	err := CONNECTION.Close()
	if err != nil {
		logger.ErrorLogger.Printf("[%v] Closing %v ", strings.ToUpper(domain), err)
		panic(commonErrors.DisconnectError(err))
	}
	logger.DatabaseLogger.Printf("[%v] Close [%v.db] data connection", strings.ToUpper(domain), dbFileName)
	timer.Stop(1)
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
		return commonErrors.UpdateError(err)
	}
	logger.DatabaseLogger.Printf("Update [%+v]", data)
	return CONNECTION.Update(data)
}

func Create(data any) error {
	err := validate(data)
	if err != nil {
		return commonErrors.CreateError(err)
	}
	logger.DatabaseLogger.Printf("Create [%+v]", data)
	return CONNECTION.Save(data)
}

func validate(data any) error {
	err := commonErrors.HandleGoValidatorError(dataValidator.Struct(data))
	if err != nil {
		logger.ErrorLogger.Printf("[%v] Validation  %v", strings.ToUpper(domain), err.Error())
		return commonErrors.ValidateError(err)
	}
	return nil
}

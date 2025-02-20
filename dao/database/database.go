package database

import (
	"strings"

	storm "github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/index"
	validator "github.com/go-playground/validator/v10"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var Version = 1
var connection *storm.DB
var domain = "database"
var dbFileName string

var dataValidator *validator.Validate

func init() {
	//	Connect()å
	dataValidator = validator.New(validator.WithRequiredStructEnabled())
}

func Connect() *storm.DB {
	return connect(domain)
}

func NamedConnect(name string) *storm.DB {
	return connect(name)
}

func connect(name string) *storm.DB {
	dbFileName = name
	connect := timing.Start(domain, "Connect", name)
	var err error
	connection, err = storm.Open(ioHelpers.GetDBFileName(dbFileName), storm.BoltOptions(0666, nil))
	if err != nil {
		connect.Stop(0)
		logHandler.ErrorLogger.Panicf("[%v] Opening [%v.db] connection Error=[%v]", strings.ToUpper(domain), strings.ToLower(dbFileName), err.Error())
		panic(commonErrors.WrapConnectError(err))
	}
	logHandler.DatabaseLogger.Printf("[%v] Opened [%v.db] data connection", strings.ToUpper(domain), dbFileName)
	connect.Stop(1)
	return connection
}

func Backup(loc string) {
	timer := timing.Start(domain, "Backup", dbFileName)
	logHandler.EventLogger.Printf("[BACKUP] Backup [%v.db] data started...", dbFileName)
	Disconnect()
	ioHelpers.Backup(dbFileName, loc)
	connect(dbFileName)
	logHandler.EventLogger.Printf("[BACKUP] Backup [%v.db] data ends", dbFileName)
	timer.Stop(1)
	logHandler.DatabaseLogger.Printf("[%v] Backup [%v.db] data connection", strings.ToUpper(domain), dbFileName)
}

func Disconnect() {
	timer := timing.Start(domain, "Disconnect", dbFileName)
	logHandler.EventLogger.Printf("[%v] Close [%v.db] data file", strings.ToUpper(domain), dbFileName)
	err := connection.Close()
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Closing %v ", strings.ToUpper(domain), err)
		panic(commonErrors.WrapDisconnectError(err))
	}
	logHandler.DatabaseLogger.Printf("[%v] Close [%v.db] data connection", strings.ToUpper(domain), dbFileName)
	timer.Stop(1)
}

func Retrieve(fieldName string, value, to any) error {
	logHandler.DatabaseLogger.Printf("Retrieve [%+v][%+v][%+v]", fieldName, value, to)
	return connection.One(fieldName, value, to)
}

func GetAll(to any, options ...func(*index.Options)) error {
	logHandler.DatabaseLogger.Printf("GetAll [%+v][%+v]", to, options)
	return connection.All(to, options...)
}

func Delete(data any) error {
	logHandler.DatabaseLogger.Printf("Delete [%+v]", data)
	return connection.DeleteStruct(data)
}

func Drop(data any) error {
	logHandler.DatabaseLogger.Printf("Drop [%+v]", data)
	return connection.Drop(data)
}

func Update(data any) error {
	err := validate(data)
	if err != nil {
		return commonErrors.WrapError(err)
	}
	logHandler.DatabaseLogger.Printf("Update [%+v]", data)
	return connection.Update(data)
}

func Create(data any) error {
	err := validate(data)
	if err != nil {
		return commonErrors.WrapCreateError(err)
	}
	logHandler.DatabaseLogger.Printf("Create [%+v]", data)
	return connection.Save(data)
}

func validate(data any) error {
	err := commonErrors.HandleGoValidatorError(dataValidator.Struct(data))
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Validation  %v", strings.ToUpper(domain), err.Error())
		return commonErrors.WrapValidationError(err)
	}
	return nil
}

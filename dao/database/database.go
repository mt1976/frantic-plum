package database

import (
	"strings"

	storm "github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/index"
	validator "github.com/go-playground/validator/v10"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var domain = "database"

type DB struct {
	connection   *storm.DB
	name         string
	databaseName string
	initialised  bool
}

var connectionPool map[string]*DB
var connectionPoolMaxSize int
var cfg *commonConfig.Settings

var dataValidator *validator.Validate

func init() {
	//	Connect()
	cfg := commonConfig.Get()

	dataValidator = validator.New(validator.WithRequiredStructEnabled())
	connectionPool = make(map[string]*DB)
	connectionPoolMaxSize = cfg.GetDatabasePoolSize()
}

func Connect() *DB {
	return connect(domain)
}

func NamedConnect(name string) *DB {
	return connect(name)
}

func connect(name string) *DB {
	if connectionPool[domain] != nil {
		logHandler.DatabaseLogger.Printf("[%v] db already open [%v.db] data connection", strings.ToUpper(domain), connectionPool[domain].name)
		return connectionPool[domain]
	}
	db := DB{}
	db.name = name
	db.databaseName = ioHelpers.GetDBFileName(name)
	connect := timing.Start(domain, actions.CONNECT.GetCode(), db.databaseName)
	var err error
	db.connection, err = storm.Open(db.databaseName, storm.BoltOptions(0777, nil))
	if err != nil {
		connect.Stop(0)
		logHandler.ErrorLogger.Panicf("[%v] Opening [%v.db] connection Error=[%v]", strings.ToUpper(domain), strings.ToLower(db.databaseName), err.Error())
		panic(commonErrors.WrapConnectError(err))
	}
	db.initialised = true
	// Add to connection pool
	storeConnectionInPool(db)
	logHandler.DatabaseLogger.Printf("[%v] Opened [%v.db] data connection", strings.ToUpper(domain), db.databaseName)
	connect.Stop(1)
	return &db
}

func storeConnectionInPool(db DB) {
	if len(connectionPool) >= connectionPoolMaxSize {
		logHandler.DatabaseLogger.Panicf("[%v] Connection pool full [%v]", strings.ToUpper(domain), connectionPoolMaxSize)
		return
	}
	connectionPool[domain] = &db
}

func (db *DB) Backup(loc string) {
	timer := timing.Start(domain, actions.BACKUP.GetCode(), db.databaseName)
	logHandler.DatabaseLogger.Printf("[BACKUP] Backup [%v.db] data started...", db.databaseName)
	db.Disconnect()
	ioHelpers.Backup(db.databaseName, loc)
	connect(db.name)
	logHandler.DatabaseLogger.Printf("[BACKUP] Backup [%v.db] data ends", db.databaseName)
	timer.Stop(1)
	logHandler.DatabaseLogger.Printf("[%v] Backup [%v.db] data connection", strings.ToUpper(domain), db.databaseName)
}

func (db *DB) Disconnect() {
	timer := timing.Start(domain, actions.DISCONNECT.Code, db.databaseName)
	logHandler.DatabaseLogger.Printf("[%v] Close [%v.db] data file", strings.ToUpper(domain), db.databaseName)
	err := db.connection.Close()
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Closing %v ", strings.ToUpper(domain), err)
		panic(commonErrors.WrapDisconnectError(err))
	}
	logHandler.DatabaseLogger.Printf("[%v] Close [%v.db] data connection", strings.ToUpper(domain), db.databaseName)
	timer.Stop(1)
}

func (db *DB) Retrieve(fieldName string, value, to any) error {
	logHandler.DatabaseLogger.Printf("Retrieve [%+v][%+v][%+v]", fieldName, value, to)
	return db.connection.One(fieldName, value, to)
}

func (db *DB) GetAll(to any, options ...func(*index.Options)) error {
	logHandler.DatabaseLogger.Printf("GetAll [%+v][%+v]", to, options)
	return db.connection.All(to, options...)
}

func (db *DB) Delete(data any) error {
	logHandler.DatabaseLogger.Printf("Delete [%+v]", data)
	return db.connection.DeleteStruct(data)
}

func (db *DB) Drop(data any) error {
	logHandler.DatabaseLogger.Printf("Drop [%+v]", data)
	return db.connection.Drop(data)
}

func (db *DB) Update(data any) error {
	err := validate(data)
	if err != nil {
		return commonErrors.WrapError(err)
	}
	logHandler.DatabaseLogger.Printf("Update [%+v]", data)
	return db.connection.Update(data)
}

func (db *DB) Create(data any) error {
	err := validate(data)
	if err != nil {
		return commonErrors.WrapCreateError(err)
	}
	logHandler.DatabaseLogger.Printf("Create [%+v]", data)
	return db.connection.Save(data)
}

func validate(data any) error {
	timer := timing.Start(domain, actions.VALIDATE.GetCode(), "")
	logHandler.DatabaseLogger.Printf("Validate [%+v]", data)
	err := commonErrors.HandleGoValidatorError(dataValidator.Struct(data))
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] Validation  %v", strings.ToUpper(domain), err.Error())
		timer.Stop(0)
		return commonErrors.WrapValidationError(err)
	}
	timer.Stop(1)
	return nil
}

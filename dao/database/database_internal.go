package database

import (
	"strings"

	"github.com/asdine/storm/v3"
	"github.com/go-playground/validator/v10"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var (
	connectionPool        map[string]*DB         = make(map[string]*DB)                                 // map of database connections, indexed by domain.
	connectionPoolMaxSize int                    = 10                                                   // maximum number of connections
	cfg                   *commonConfig.Settings = commonConfig.Get()                                   // configuration settings
	dataValidator         *validator.Validate    = validator.New(validator.WithRequiredStructEnabled()) // data validator
)

func init() {
	//	Connect()
	//	cfg = commonConfig.Get()
	//dataValidator = validator.New(validator.WithRequiredStructEnabled())
	//connectionPool = make(map[string]*DB)
	connectionPoolMaxSize = cfg.GetDatabasePoolSize()
}

func connect(name string) *DB {
	// Ensure the name is lowercase
	name = strings.ToLower(name)
	logHandler.DatabaseLogger.Printf("Opening Connection to [%v.db] data", name)
	// list the connection pool
	for key, value := range connectionPool {
		logHandler.DatabaseLogger.Printf("Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	// check if connection already exists
	if connectionPool[name] != nil && connectionPool[name].name == name {
		logHandler.DatabaseLogger.Printf("Connection already open [%v], using connection pool [%v] [codec=%v]", connectionPool[name].name, connectionPool[name].databaseName, connectionPool[name].connection.Node.Codec().Name())
		return connectionPool[name]
	}

	db := DB{}
	db.name = name
	db.databaseName = ioHelpers.GetDBFileName(db.name)
	connect := timing.Start(db.name, actions.CONNECT.GetCode(), db.databaseName)
	var err error
	db.connection, err = storm.Open(db.databaseName, storm.BoltOptions(0777, nil))
	if err != nil {
		connect.Stop(0)
		logHandler.ErrorLogger.Panicf("Opening [%v.db] connection Error=[%v]", strings.ToLower(db.databaseName), err.Error())
		panic(commonErrors.WrapConnectError(err))
	}
	db.initialised = true
	// Add to connection pool
	storeConnectionInPool(db, db.name)
	logHandler.DatabaseLogger.Printf("Opened [%v.db] data connection [codec=%v]", db.databaseName, db.connection.Node.Codec().Name())
	connect.Stop(1)
	return &db
}

func storeConnectionInPool(db DB, key string) {
	if len(connectionPool) >= connectionPoolMaxSize {
		logHandler.DatabaseLogger.Panicf("Connection pool full [%v]", connectionPoolMaxSize)
		return
	}
	connectionPool[key] = &db
}

func validate(data any, db *DB) error {
	timer := timing.Start(db.name, actions.VALIDATE.GetCode(), "")
	logHandler.DatabaseLogger.Printf("Validate [%+v] [%v.db]", dao.GetStructType(data), db.name)
	err := commonErrors.HandleGoValidatorError(dataValidator.Struct(data))
	if err != nil {
		logHandler.ErrorLogger.Printf("Validating %v %v [%v.db]", err.Error(), dao.GetStructType(data), db.name)
		timer.Stop(0)
		return commonErrors.WrapValidationError(err)
	}
	timer.Stop(1)
	return nil
}

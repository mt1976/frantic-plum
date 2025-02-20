package database

import (
	"reflect"
	"strings"

	storm "github.com/asdine/storm/v3"
	"github.com/go-playground/validator/v10"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var (
	domain                string                 = "database"                                           // domain for this code module
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

func validate(data any, db *DB) error {
	timer := timing.Start(domain, actions.VALIDATE.GetCode(), "")
	logHandler.DatabaseLogger.Printf("Validate [%+v] [%v.db]", getType(data), db.name)
	err := commonErrors.HandleGoValidatorError(dataValidator.Struct(data))
	if err != nil {
		logHandler.ErrorLogger.Printf("Validating %v %v [%v.db]", err.Error(), getType(data), db.name)
		timer.Stop(0)
		return commonErrors.WrapValidationError(err)
	}
	timer.Stop(1)
	return nil
}

func getType(data any) string {
	rtnType := reflect.TypeOf(data).String()
	return rtnType
}

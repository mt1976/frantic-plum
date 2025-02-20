package database

import (
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/index"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

type DB struct {
	connection   *storm.DB
	name         string
	databaseName string
	initialised  bool
}

func Connect() *DB {
	return connect("database")
}

func NamedConnect(name string) *DB {
	return connect(name)
}

func (db *DB) Backup(loc string) {
	timer := timing.Start(db.name, actions.BACKUP.GetCode(), db.databaseName)
	logHandler.DatabaseLogger.Printf("Backup [%v.db] data started...", db.databaseName)
	db.Disconnect()
	ioHelpers.Backup(db.databaseName, loc)
	connect(db.name)
	logHandler.DatabaseLogger.Printf("Backup [%v.db] data ends", db.databaseName)
	timer.Stop(1)
	logHandler.DatabaseLogger.Printf("Backup [%v.db] data connection", db.databaseName)
}

func (db *DB) Disconnect() {
	timer := timing.Start(db.name, actions.DISCONNECT.Code, db.databaseName)
	logHandler.DatabaseLogger.Printf("Closing [%v.db] connection", db.name)
	err := db.connection.Close()
	if err != nil {
		logHandler.ErrorLogger.Printf("Closing [%v.db] %v ", db.name, err.Error())
		panic(commonErrors.WrapDisconnectError(err))
	}
	logHandler.DatabaseLogger.Printf("Closed [%v.db] connection", db.name)
	timer.Stop(1)
}

func (db *DB) Retrieve(fieldName string, value, to any) error {
	logHandler.DatabaseLogger.Printf("Retrieve (%+v=%+v)[%+v] [%v.db]", fieldName, value, getType(to), db.name)
	return db.connection.One(fieldName, value, to)
}

func (db *DB) GetAll(to any, options ...func(*index.Options)) error {
	logHandler.DatabaseLogger.Printf("GetAll [%+v][%+v] [%v.db]", getType(to), options, db.name)
	return db.connection.All(to, options...)
}

func (db *DB) Delete(data any) error {
	logHandler.DatabaseLogger.Printf("Delete [%+v] [%v.db]", getType(data), db.name)
	return db.connection.DeleteStruct(data)
}

func (db *DB) Drop(data any) error {
	logHandler.DatabaseLogger.Printf("Drop [%+v] [%v.db]", getType(data), db.name)
	return db.connection.Drop(data)
}

func (db *DB) Update(data any) error {
	err := validate(data, db)
	if err != nil {
		return commonErrors.WrapError(err)
	}
	logHandler.DatabaseLogger.Printf("Update [%+v] [%v.db]", getType(data), db.name)
	return db.connection.Update(data)
}

func (db *DB) Create(data any) error {
	err := validate(data, db)
	if err != nil {
		return commonErrors.WrapCreateError(err)
	}
	logHandler.DatabaseLogger.Printf("Create [%+v] [%v.db]", getType(data), db.name)
	return db.connection.Save(data)
}

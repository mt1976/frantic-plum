package database

import (
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/index"
	"github.com/asdine/storm/v3/q"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

type DB struct {
	connection   *storm.DB
	Name         string
	databaseName string
	initialised  bool
}

func Connect() *DB {
	return connect("database")
}

func ConnectToNamedDB(name string) *DB {
	return connect(name)
}

func (db *DB) Reconnect() {
	logHandler.InfoLogger.Printf("[RECONNECT] Reconnecting [%v.db] data - %+v", db.Name, db)
	logHandler.InfoLogger.Printf("[RECONNECT] Connection Pool [%+v]", connectionPool)
	for key, value := range connectionPool {
		logHandler.InfoLogger.Printf("[RECONNECT] Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	connect(db.Name)
	logHandler.InfoLogger.Printf("[RECONNECT] Reconnected [%v.db] data", db.Name)
}

func (db *DB) Backup(loc string) {
	timer := timing.Start(db.Name, actions.BACKUP.GetCode(), db.databaseName)
	logHandler.InfoLogger.Printf("Backup [%v.db] data started... %v", db.Name, loc)
	db.Disconnect()
	logHandler.InfoLogger.Printf("Backup [%v.db] disconnected", db.Name)
	ioHelpers.Backup(db.Name, loc)
	logHandler.InfoLogger.Printf("Backup [%v.db] backup done ends", db.Name)
	db.Reconnect()
	//connect(db.Name)
	logHandler.InfoLogger.Printf("Backup [%v.db] (re)connected", db.Name)
	timer.Stop(1)
	logHandler.InfoLogger.Printf("Backup [%v.db] data connection", db.Name)
}

func (db *DB) Disconnect() {
	timer := timing.Start(db.Name, actions.DISCONNECT.Code, db.databaseName)
	logHandler.InfoLogger.Printf("[DISCONNECT] Disconnecting [%v.db] connection", db.Name)
	err := db.connection.Close()
	if err != nil {
		logHandler.ErrorLogger.Printf("[DISCONNECT] Closing [%v.db] %v ", db.Name, err.Error())
		panic(commonErrors.WrapDisconnectError(err))
	}
	releaseFromConnectionPool(db)
	logHandler.InfoLogger.Printf("[DISCONNECT] Closed [%v.db] connection", db.Name)
	for key, value := range connectionPool {
		logHandler.InfoLogger.Printf("[DISCONNECT] Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	timer.Stop(1)
}

func (db *DB) Retrieve(fieldName string, value, to any) error {
	logHandler.DatabaseLogger.Printf("Retrieve (%+v=%+v)[%+v] [%v.db]", fieldName, value, dao.GetStructType(to), db.Name)
	return db.connection.One(fieldName, value, to)
}

func (db *DB) GetAll(to any, options ...func(*index.Options)) error {
	logHandler.DatabaseLogger.Printf("GetAll [%+v][%+v] [%v.db]", dao.GetStructType(to), options, db.Name)
	return db.connection.All(to, options...)
}

func (db *DB) Delete(data any) error {
	logHandler.DatabaseLogger.Printf("Delete [%+v] [%v.db]", dao.GetStructType(data), db.Name)
	return db.connection.DeleteStruct(data)
}

func (db *DB) Drop(data any) error {
	logHandler.DatabaseLogger.Printf("Drop [%+v] [%v.db]", dao.GetStructType(data), db.Name)
	return db.connection.Drop(data)
}

func (db *DB) Update(data any) error {
	err := validate(data, db)
	if err != nil {
		return commonErrors.WrapError(err)
	}
	logHandler.DatabaseLogger.Printf("Update [%+v] [%v.db]", dao.GetStructType(data), db.Name)
	return db.connection.Update(data)
}

func (db *DB) Create(data any) error {
	err := validate(data, db)
	if err != nil {
		return commonErrors.WrapCreateError(err)
	}
	logHandler.DatabaseLogger.Printf("Create [%+v] [%v.db]", dao.GetStructType(data), db.Name)
	return db.connection.Save(data)
}

func (db *DB) Count(data any) (int, error) {
	logHandler.DatabaseLogger.Printf("Count [%+v] [%v.db]", dao.GetStructType(data), db.Name)
	logHandler.InfoLogger.Printf("Count [%+v] [%v.db] [%+v]", dao.GetStructType(data), db.Name, db)
	for key, value := range connectionPool {
		logHandler.InfoLogger.Printf("Connection Pool [%v] [%v] [codec=%v]", key, value.databaseName, value.connection.Node.Codec().Name())
	}
	//logHandler.InfoLogger.Printf("connectionPool [%+v]", connectionPool)
	return db.connection.Count(data)
}

func (db *DB) CountWhere(fieldName string, value any, to any) (int, error) {
	logHandler.DatabaseLogger.Printf("CountWhere (%+v=%+v)[%+v] [%v.db]", fieldName, value, dao.GetStructType(to), db.Name)
	if err := dao.IsValidFieldInStruct(fieldName, to); err != nil {
		return 0, err
	}
	query := db.connection.Select(q.Eq(fieldName, value))
	count, err := query.Count(to)
	return count, err
}

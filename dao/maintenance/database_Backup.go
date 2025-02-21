package maintenance

import (
	"strings"
	"time"

	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/frantic-core/timing"
)

type DatabaseBackupJob struct {
	funcs []func() (*database.DB, error)
}

func (job *DatabaseBackupJob) Run() error {

	jobs.Announce(job.Name(), "Started")

	performDatabaseBackup(job)

	jobs.NextRun(job.Name(), job.Schedule())
	jobs.Announce(job.Name(), "Completed")

	return nil
}

func (job *DatabaseBackupJob) Service() func() {
	return func() {
		_ = job.Run()
	}
}

func (job *DatabaseBackupJob) Schedule() string {
	return "55 11 * * *"
}

func (job *DatabaseBackupJob) Name() string {
	//name, _ := translation.Get("Scheduled Database Backup")
	return "Maintenance - Backup Database"
}

func performDatabaseBackup(job *DatabaseBackupJob) {
	logHandler.InfoLogger.Printf("[%v] [%v] Started", domain, job.Name())
	j := timing.Start(job.Name(), actions.BACKUP.Code, "All")

	dateTime := time.Now().Format(dateHelpers.Format.BackupFolder)
	logHandler.InfoLogger.Printf("[%v] [BACKUP] Date=[%v]", domain, dateTime)

	destPath := paths.Backups().String() + paths.Seperator() + dateTime
	fullBackupPath := paths.Application().String() + destPath

	//create a folder
	err := ioHelpers.MkDir(fullBackupPath)
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] [%v] Error=[%v]", domain, strings.ToUpper(job.Name()), err.Error())
	}

	for _, thisFunc := range job.funcs {
		logHandler.ServiceLogger.Printf("[%v] [%v]", domain, job.Name())
		db, err := thisFunc()
		if err != nil {
			logHandler.ErrorLogger.Panicf("[%v] [%v] Error=[%v]", domain, strings.ToUpper(job.Name()), err.Error())
		}
		logHandler.InfoLogger.Printf("[%v] [%v] Backup [%v]", domain, job.Name(), db.Name)
		db.Disconnect()
		logHandler.InfoLogger.Printf("[%v] [%v] Disconnected [%v]", domain, job.Name(), db.Name)
		db.Backup(fullBackupPath)
		logHandler.InfoLogger.Printf("[%v] [%v] Backup Done [%v]", domain, job.Name(), db.Name)
		db.Reconnect()
		logHandler.InfoLogger.Printf("[%v] [%v] Reconnected [%v]", domain, job.Name(), db.Name)
	}
	j.Stop(len(job.funcs))
	logHandler.InfoLogger.Printf("[%v] [%v] Completed", domain, job.Name())
}

func (job *DatabaseBackupJob) AddFunction(fn func() (*database.DB, error)) {
	job.funcs = append(job.funcs, fn)
}

func (job *DatabaseBackupJob) Description() string {
	return "Scheduled Database Backup, runs at 11:55 daily"
}

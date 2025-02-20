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
	logHandler.ServiceLogger.Printf("[%v] [%v] Started", domain, job.Name())
	j := timing.Start(job.Name(), actions.BACKUP.Code, "All")

	dateTime := time.Now().Format(dateHelpers.Format.BackupFolder)
	logHandler.ServiceLogger.Printf("[%v] [BACKUP] Date=[%v]", domain, dateTime)

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
		db.Backup(fullBackupPath)
	}
	j.Stop(len(job.funcs))
	logHandler.ServiceLogger.Printf("[%v] [%v] Completed", domain, job.Name())
}

func (job *DatabaseBackupJob) AddFunction(f func() (*database.DB, error)) {
	job.funcs = append(job.funcs, f)
}

func (job *DatabaseBackupCleanerJob) Description() string {
	return "Scheduled Database Maintenance - Prune Old Backups"
}

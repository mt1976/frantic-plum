package maintenance

import (
	"strings"
	"time"

	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/frantic-core/timing"
)

type DatabaseBackupJob struct {
	db   *database.DB
	name string
}

func (job DatabaseBackupJob) Run() error {
	jobs.Announce(job, "Started")
	performDatabaseBackup(job)
	jobs.NextRun(job)
	jobs.Announce(job, "Completed")
	return nil
}

func (job DatabaseBackupJob) Service() func() {
	return func() {
		_ = job.Run()
	}
}

func (job DatabaseBackupJob) Schedule() string {
	return "55 11 * * *"
}

func (job DatabaseBackupJob) Name() string {
	//name, _ := translation.Get("Scheduled Database Backup")
	return "Maintenance - Backup - " + job.name
}

func performDatabaseBackup(job DatabaseBackupJob) {
	j := timing.Start(job.Name(), "Backup", "All")

	dateTime := time.Now().Format(dateHelpers.Format.BackupFolder)
	logHandler.ServiceLogger.Printf("[%v] [BACKUP] Date=[%v]", domain, dateTime)

	destPath := paths.Backups().String() + paths.Seperator() + dateTime
	fullBackupPath := paths.Application().String() + destPath

	//create a folder
	err := ioHelpers.MkDir(fullBackupPath)
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] [%v] Error=[%v]", domain, strings.ToUpper(job.Name()), err.Error())
	}

	job.db.Backup(fullBackupPath)

	j.Stop(6)
}

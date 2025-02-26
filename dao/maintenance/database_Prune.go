package maintenance

import (
	"fmt"
	"time"

	"github.com/mt1976/frantic-core/application"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/frantic-core/timing"
)

type DatabaseBackupCleanerJob struct {
}

func (job *DatabaseBackupCleanerJob) Run() error {
	jobs.PreRun(job)
	pruneExpiredBackups(job)
	jobs.PostRun(job)
	return nil
}

func (job *DatabaseBackupCleanerJob) Service() func() {
	return func() {
		_ = job.Run()
	}
}

func (job *DatabaseBackupCleanerJob) Schedule() string {
	return "25 0 * * *"
}

func (job *DatabaseBackupCleanerJob) Name() string {
	//name, _ := translation.Get("Scheduled Database Maintenance - Prune Old Backups")
	return "Maintenance - Prune Old Backups"
}

func pruneExpiredBackups(job *DatabaseBackupCleanerJob) {

	settings := commonConfig.Get()
	// Do something every day at midnight
	name := jobs.CodedName(job)

	j := timing.Start(job.Name(), actions.MAINTENANCE.GetCode(), job.Description())
	// Get Settings

	retainBackupDays := settings.GetHistory_MaxHistoryEntries()

	logHandler.ServiceLogger.Printf("[%v] RetainBackupDays=[%v]", name, retainBackupDays)
	today := jobs.StartOfDay(time.Now())

	// get today's date
	DMY := dateHelpers.Format.DMY
	todayStr := today.Format(DMY)
	logHandler.ServiceLogger.Printf("[%v] Today=[%v]", name, todayStr)
	deleteBeforeDate := today.AddDate(0, 0, -retainBackupDays)
	deleteBeforeDateStr := deleteBeforeDate.Format(DMY)
	logHandler.ServiceLogger.Printf("[%v] DeleteBeforeDate=[%v]", name, deleteBeforeDateStr)

	// Get Backups path
	path := paths.Backups().String()
	logHandler.ServiceLogger.Printf("[%v] Path=[%v]", name, path)
	full := paths.Application().String()
	logHandler.ServiceLogger.Printf("[%v] AppPath=[%v]", name, full)
	backupPath := full + path
	logHandler.ServiceLogger.Printf("[%v] BackupPath=[%v]", name, backupPath)

	// Get all folders in the backup directory
	folders, err := ioHelpers.Dir(backupPath)
	if err != nil {
		logHandler.ServiceLogger.Panicf("[%v] Error=[%v]", name, err.Error())
		return
	}
	logHandler.ServiceLogger.Printf("[%v] No Folders=[%v]", name, len(folders))
	count := 0
	// For each folder check if it is before the deleteBeforeDate
	for _, folder := range folders {
		// Get the date from the folder name
		backupDate, err := getDateFromBackupFolderName(folder)
		if err != nil {
			logHandler.ServiceLogger.Panicf("[%v] Error=[%v]", name, err.Error())
			return
		}
		// Check if the backupDate is before the deleteBeforeDate
		if backupDate.Before(deleteBeforeDate) {
			// Delete the folder
			logHandler.ServiceLogger.Printf("[%v] Deleting=[%v] FolderDate=[%v] DeleteDate=[%v]", name, folder, backupDate.Format(DMY), deleteBeforeDateStr)
			count++
			err := ioHelpers.DeleteFolder(backupPath + folder)
			if err != nil {
				logHandler.ErrorLogger.Printf("[%v] Error=[%v]", name, err.Error())
				return
			}
			msg := "Backup Pruned Folder=[%v] On=[%v]"
			msg = fmt.Sprintf(msg, folder, application.HostName())
			logHandler.ServiceLogger.Printf("[%v] [%v]", name, msg)
		}
	}
	j.Stop(count)
}

func getDateFromBackupFolderName(folder string) (date time.Time, err error) {
	// Get the date from the folder name
	date, err = time.Parse(dateHelpers.Format.BackupFolder, folder)
	if err != nil {
		logHandler.ServiceLogger.Panicf("[%v] [%v] Error=[%v]", domain, "BACKUP", err.Error())
		return
	}
	return
}

func (job *DatabaseBackupCleanerJob) AddDatabaseAccessFunctions(fn func() ([]*database.DB, error)) {
	// do nothing
	panic("Not Implemented")
}

func (job *DatabaseBackupCleanerJob) Description() string {
	return "Scheduled Database Maintenance - Prunes Old Backups, Retaining the last 30 days, run at 00:25 every day"
}

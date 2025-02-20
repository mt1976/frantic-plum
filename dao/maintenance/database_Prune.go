package maintenance

import (
	"fmt"
	"strings"
	"time"

	"github.com/mt1976/frantic-core/application"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/ioHelpers"
	"github.com/mt1976/frantic-core/jobs"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/frantic-core/timing"
)

type DatabaseBackupCleanerJob struct {
}

func (p DatabaseBackupCleanerJob) Run() error {
	jobs.Announce(p, "Started")
	pruneExpiredBackups()
	jobs.NextRun(p)
	jobs.Announce(p, "Completed")
	return nil
}

func (p DatabaseBackupCleanerJob) Service() func() {
	return func() {
		_ = p.Run()
	}
}

func (p DatabaseBackupCleanerJob) Schedule() string {
	return "25 0 * * *"
}

func (p DatabaseBackupCleanerJob) Name() string {
	//name, _ := translation.Get("Scheduled Database Maintenance - Prune Old Backups")
	return "Maintenance - Prune Old Backups"
}

func pruneExpiredBackups() {

	settings := commonConfig.Get()
	// Do something every day at midnight
	name := "Prune"
	j := timing.Start(strings.ToUpper(name), "Prune", "Old Backups")
	// Get Settings

	retainBackupDays := settings.GetMaxHistoryEntries()

	logHandler.ServiceLogger.Printf("[%v] RetainBackupDays=[%v]", strings.ToUpper(name), retainBackupDays)
	today := jobs.StartOfDay(time.Now())

	// get today's date
	DMY := dateHelpers.Format.DMY
	todayStr := today.Format(DMY)
	logHandler.ServiceLogger.Printf("[%v] Today=[%v]", strings.ToUpper(name), todayStr)
	deleteBeforeDate := today.AddDate(0, 0, -retainBackupDays)
	deleteBeforeDateStr := deleteBeforeDate.Format(DMY)
	logHandler.ServiceLogger.Printf("[%v] DeleteBeforeDate=[%v]", strings.ToUpper(name), deleteBeforeDateStr)

	// Get Backups path
	path := paths.Backups().String()
	logHandler.ServiceLogger.Printf("[%v] Path=[%v]", strings.ToUpper(name), path)
	full := paths.Application().String()
	logHandler.ServiceLogger.Printf("[%v] AppPath=[%v]", strings.ToUpper(name), full)
	backupPath := full + path
	logHandler.ServiceLogger.Printf("[%v] BackupPath=[%v]", strings.ToUpper(name), backupPath)

	// Get all folders in the backup directory
	folders, err := ioHelpers.Dir(backupPath)
	if err != nil {
		logHandler.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		return
	}
	logHandler.ServiceLogger.Printf("[%v] No Folders=[%v]", strings.ToUpper(name), len(folders))
	count := 0
	// For each folder check if it is before the deleteBeforeDate
	for _, folder := range folders {
		// Get the date from the folder strings.ToUpper(name)
		backupDate, err := getDateFromBackupFolderName(folder)
		if err != nil {
			logHandler.ErrorLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
			return
		}
		// Check if the backupDate is before the deleteBeforeDate
		if backupDate.Before(deleteBeforeDate) {
			// Delete the folder
			logHandler.ServiceLogger.Printf("[%v] Deleting=[%v] FolderDate=[%v] DeleteDate=[%v]", strings.ToUpper(name), folder, backupDate.Format(DMY), deleteBeforeDateStr)
			count++
			err := ioHelpers.DeleteFolder(backupPath + folder)
			if err != nil {
				logHandler.ErrorLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
				return
			}
			msg := "Backup Pruned Folder=[%v] On=[%v]"
			msg = fmt.Sprintf(msg, folder, application.HostName())
			logHandler.ServiceLogger.Printf("[%v] [%v]", strings.ToUpper(name), msg)
		}
	}
	j.Stop(count)
}

func getDateFromBackupFolderName(folder string) (date time.Time, err error) {
	// Get the date from the folder strings.ToUpper(name)
	date, err = time.Parse(dateHelpers.Format.BackupFolder, folder)
	if err != nil {
		logHandler.ErrorLogger.Printf("[%v] [%v] Error=[%v]", domain, "BACKUP", err.Error())
		return
	}
	return
}

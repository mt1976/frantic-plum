package importExportHelper

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/frantic-core/timing"
)

// TemplateImportData is a struct to hold the data from the CSV file
// it is used to import the data into the database
// The struct tags are used to map the fields to the CSV columns
// this struct should be customised to suit the specific requirements of the entryination table/DAO.

var FIELDSEPARATOR = '|'
var importString = "Import"
var exportString = "Export"

func ExportCSV[T any](exportName string, exportList []T) error {
	clock := timing.Start(exportName, actions.EXPORT.GetCode(), "")

	exportFile := openTargetFile(exportName, exportString, logHandler.ExportLogger)
	defer exportFile.Close()

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = FIELDSEPARATOR // Use tab-delimited format
		writer.UseCRLF = true
		return gocsv.NewSafeCSVWriter(writer)
	})

	_, err := gocsv.MarshalString(exportList) // Get all texts as CSV string
	if err != nil {
		logHandler.ExportLogger.Panicf("error exporting %v: %v", exportName, err.Error())
	}

	err = gocsv.MarshalFile(exportList, exportFile) // Get all texts as CSV string
	if err != nil {
		logHandler.ExportLogger.Panicf("error exporting %v: %v", exportName, err.Error())
	}

	msg := fmt.Sprintf("# Generated (%v) %vs at %v on %v", len(exportList), exportName, time.Now().Format("15:04:05"), time.Now().Format("2006-01-02"))
	exportFile.WriteString(msg)

	exportFile.Close()

	logHandler.ExportLogger.Printf("Exported (%v/%v) %v(s) to [%v]", len(exportList), len(exportList), exportName, exportFile.Name())
	logHandler.EventLogger.Printf("Exported (%v/%v) %v(s) to [%v]", len(exportList), len(exportList), exportName, exportFile.Name())
	clock.Stop(len(exportList))
	return nil
}

func openTargetFile(in, action string, useLog *log.Logger) *os.File {
	defaultPath := paths.Defaults()
	templateDataFileName := strings.ToLower(in) + "s.csv"
	fileName := fmt.Sprintf("%s%s/%s", paths.Application().String(), defaultPath.String(), templateDataFileName)

	dataFileHandle, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		useLog.Fatalf("error opening file: %v", err)
		panic(err)
	}
	useLog.Printf("%ving %vs from File=[%v]", action, in, dataFileHandle.Name())
	return dataFileHandle
}

func ImportCSV[T any](importName string, entryTypeToInsert T, importProcessor func(*T) (string, error)) error {

	clock := timing.Start(importName, actions.IMPORT.GetCode(), "")
	// Create a slice of entryTypeToInsert to hold the data from the CSV file
	insertEntriesList := []T{}

	importFile := openTargetFile(importName, importString, logHandler.ImportLogger)
	defer importFile.Close()

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)    // Allows use pipe as delimiter
		r.Comma = FIELDSEPARATOR  // Use tab-delimited format
		r.Comment = '#'           // Ignore comment lines
		r.TrimLeadingSpace = true // Trim leading space
		return r                  // Allows use pipe as delimiter
	})

	if err := gocsv.UnmarshalFile(importFile, &insertEntriesList); err != nil { // Load clients from file
		logHandler.ImportLogger.Printf("Importing %v: %v - No Content, nothing to import.", importName, err.Error())
		importFile.Close()
		clock.Stop(0)
		return nil
	}

	if _, err := importFile.Seek(0, 0); err != nil { // Go to the start of the file
		logHandler.ImportLogger.Panicf("Importing %v: %v - Unable to fet to start of file.", importName, err.Error())
		clock.Stop(0)
		panic(err)
	}

	totalImportEntries := len(insertEntriesList)

	count := 0
	for thisPos, insertEntry := range insertEntriesList {
		//logHandler.ImportLogger.Printf("Import %v (%v/%v)", importName, thisPos+1, totalImportEntries)
		// the load function is a helper function to create a new entry instance and save it to the database
		// the parameters should be customised to suit the specific requirements of the entryination table/DAO.
		entryIdentifier, err := importProcessor(&insertEntry)
		if err != nil {
			logHandler.ImportLogger.Panicf("Error importing %v [%v] Error=[%v]", importName, entryIdentifier, err.Error())
			continue
		}
		logHandler.ImportLogger.Printf("Imported %v (%v/%v) %v=[%v]", importName, thisPos+1, totalImportEntries, importName, entryIdentifier)

		count++
	}

	logHandler.ImportLogger.Printf("Imported (%v/%v) %v(s) from [%v]", count, totalImportEntries, importName, importFile.Name())
	logHandler.EventLogger.Printf("Imported (%v/%v) %v(s) from [%v]", count, totalImportEntries, importName, importFile.Name())
	importFile.Close()
	clock.Stop(count)
	return nil
}

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
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
)

// TemplateImportData is a struct to hold the data from the CSV file
// it is used to import the data into the database
// The struct tags are used to map the fields to the CSV columns
// this struct should be customised to suit the specific requirements of the entryination table/DAO.

var FIELDSEPARATOR = '|'
var domain = "impex"
var importString = "Import"
var exportString = "Export"

func ExportCSV[T any](exportName string, exportList []T) error {
	logHandler.ExportLogger.Printf("Exporting %v", exportName)

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

	logHandler.ExportLogger.Printf("Exported (%v) %vs", len(exportList), exportName)
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
	useLog.Printf("%v=[%v] File=[%v]", action, in, dataFileHandle.Name())
	return dataFileHandle
}

func ImportCSV[T any](importName string, entryTypeToInsert T, importProcessor func(*T) (string, error)) error {

	// Create a slice of entryTypeToInsert to hold the data from the CSV file
	insertEntriesList := []T{}

	csvFile := openTargetFile(importName, importString, logHandler.ImportLogger)
	defer csvFile.Close()

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)    // Allows use pipe as delimiter
		r.Comma = FIELDSEPARATOR  // Use tab-delimited format
		r.Comment = '#'           // Ignore comment lines
		r.TrimLeadingSpace = true // Trim leading space
		return r                  // Allows use pipe as delimiter
	})

	if err := gocsv.UnmarshalFile(csvFile, &insertEntriesList); err != nil { // Load clients from file
		logHandler.ImportLogger.Printf("Importing %v: %v - No Content, nothing to import.", importName, err.Error())
		csvFile.Close()
		return nil
	}

	if _, err := csvFile.Seek(0, 0); err != nil { // Go to the start of the file
		logHandler.ImportLogger.Printf("Importing %v: %v", importName, err.Error())
		panic(err)
	}

	totalImportEntries := len(insertEntriesList)

	count := 0
	for thisPos, insertEntry := range insertEntriesList {
		logHandler.ImportLogger.Printf("Import %v (%v/%v)", importName, thisPos+1, totalImportEntries)
		// the load function is a helper function to create a new entry instance and save it to the database
		// the parameters should be customised to suit the specific requirements of the entryination table/DAO.
		entryIdentifier, err := importProcessor(&insertEntry)
		if err != nil {
			logHandler.ImportLogger.Panicf("Error importing %v [%v] [%v]", importName, entryIdentifier, err.Error())
			continue
		}
		count++
		logHandler.ImportLogger.Printf("Import %v (%v/%v) - %v=[%v]", importName, thisPos, totalImportEntries, domain, entryIdentifier)
	}

	logHandler.ImportLogger.Printf("Imported (%v/%v) %v", count, totalImportEntries, importName)
	csvFile.Close()
	return nil
}

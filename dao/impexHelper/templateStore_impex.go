package impexHelper

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

func ExportCSV(exportName string, exports []any) error {
	logHandler.ExportLogger.Printf("Exporting %v", exportName)

	// initialiser(context.TODO())

	exportFile := openTargetFile(exportString, exportName, logHandler.ExportLogger)
	defer exportFile.Close()

	// exports, err := getter()
	// if err != nil {
	// 	logHandler.ExportLogger.Panicf("Error Getting all texts: %v", err.Error())
	// }

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = FIELDSEPARATOR // Use tab-delimited format
		writer.UseCRLF = true
		return gocsv.NewSafeCSVWriter(writer)
	})

	_, err := gocsv.MarshalString(exports) // Get all texts as CSV string
	if err != nil {
		logHandler.ExportLogger.Panicf("error exporting texts: %v", err.Error())
	}

	err = gocsv.MarshalFile(&exports, exportFile) // Get all texts as CSV string
	if err != nil {
		logHandler.ExportLogger.Panicf("error exporting texts: %v", err.Error())
	}

	msg := fmt.Sprintf("# Generated (%v) %vs at %v on %v", len(exports), exportName, time.Now().Format("15:04:05"), time.Now().Format("2006-01-02"))
	exportFile.WriteString(msg)

	exportFile.Close()

	logHandler.ExportLogger.Printf("Exported (%v) %vs", len(exports), exportName)
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

func ImportCSV(importName string, entriesToInsert []any, importMapper func(any) (string, error)) error {
	logHandler.ImportLogger.Printf("Importing %v", importName)

	csvFile := openTargetFile(importString, importName, logHandler.ImportLogger)
	defer csvFile.Close()

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)    // Allows use pipe as delimiter
		r.Comma = FIELDSEPARATOR  // Use tab-delimited format
		r.Comment = '#'           // Ignore comment lines
		r.TrimLeadingSpace = true // Trim leading space
		return r                  // Allows use pipe as delimiter
	})

	if err := gocsv.UnmarshalFile(csvFile, &entriesToInsert); err != nil { // Load clients from file
		logHandler.ImportLogger.Printf("Importing %v: %v - No Content, nothing to import.", domain, err.Error())
		csvFile.Close()
		return nil
	}

	if _, err := csvFile.Seek(0, 0); err != nil { // Go to the start of the file
		logHandler.ImportLogger.Printf("Importing %v: %v", domain, err.Error())
		panic(err)
	}
	totalImportEntries := len(entriesToInsert)
	for thisPos, insertEntry := range entriesToInsert {
		logHandler.ImportLogger.Printf("Importing %v (%v/%v)", domain, thisPos+1, totalImportEntries)
		// the load function is a helper function to create a new entry instance and save it to the database
		// the parameters should be customised to suit the specific requirements of the entryination table/DAO.
		entryIdentifier, err := importMapper(insertEntry)
		if err != nil {
			logHandler.ImportLogger.Panicf("Error importing %v [%v] [%v]", domain, entryIdentifier, err.Error())
		}
		logHandler.ImportLogger.Printf("Imported %v [%v]", domain, entryIdentifier)
	}

	logHandler.ImportLogger.Printf("Imported (%v) %v", len(entriesToInsert), domain)
	csvFile.Close()
	return nil
}

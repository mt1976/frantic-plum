package logger

import (
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/mt1976/frantic-plum/colours"
	"github.com/mt1976/frantic-plum/common"
	"github.com/mt1976/frantic-plum/paths"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	WarningLogger     *log.Logger
	InfoLogger        *log.Logger
	ErrorLogger       *log.Logger
	PanicLogger       *log.Logger
	TimingLogger      *log.Logger
	EventLogger       *log.Logger
	ServiceLogger     *log.Logger
	TraceLogger       *log.Logger
	AuditLogger       *log.Logger
	TranslationLogger *log.Logger
	SecurityLogger    *log.Logger
	DatabaseLogger    *log.Logger
	ApiLogger         *log.Logger
)

var Reset string
var Red string
var Green string
var Yellow string
var Blue string
var Magenta string
var Cyan string
var Gray string
var White string

func init() {
	settings := common.Get()
	//prefix := "data/logs/"
	prefix := paths.Application().String() + paths.Logs().String() + string(os.PathSeparator)
	name := prefix + settings.ApplicationName() + "-"

	generalWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: fileName(name, "general"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	timingWriter := io.MultiWriter(&lumberjack.Logger{Filename: fileName(name, "timing"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	serviceWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: fileName(name, "service"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	auditWriter := io.MultiWriter(&lumberjack.Logger{Filename: fileName(name, "audit"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	errorWriter := io.MultiWriter(os.Stdout, os.Stderr, &lumberjack.Logger{Filename: fileName(name, "error"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	panicWriter := io.MultiWriter(os.Stdout, os.Stderr, &lumberjack.Logger{Filename: fileName(name, "panic"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	translationWriter := io.MultiWriter(&lumberjack.Logger{Filename: fileName(name, "translation"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	traceWriter := io.MultiWriter(&lumberjack.Logger{Filename: fileName(name, "trace"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	warningWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: fileName(name, "warning"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	eventWriter := io.MultiWriter(os.Stderr, &lumberjack.Logger{Filename: fileName(name, "event"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	securityWriter := io.MultiWriter(os.Stderr, &lumberjack.Logger{Filename: fileName(name, "security"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	databaseWriter := io.MultiWriter(&lumberjack.Logger{Filename: fileName(name, "database"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	apiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: fileName(name, "api"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})

	//fmt.Printf("name: %v\n", name)
	//os.Exit(1)
	setColoursNormal()
	if runtime.GOOS == "windows" {
		setColoursWindows()
	}

	msgStructure := log.Ldate | log.Ltime | log.Lshortfile

	InfoLogger = log.New(generalWriter, nameIt(Cyan, "Info"), msgStructure)
	WarningLogger = log.New(warningWriter, nameIt(Yellow, "Warning"), msgStructure)
	ErrorLogger = log.New(errorWriter, nameIt(Red, "Error"), msgStructure)
	PanicLogger = log.New(panicWriter, nameIt(Red, "Panic"), msgStructure)
	TimingLogger = log.New(timingWriter, nameIt(Gray, "Timing"), msgStructure)
	EventLogger = log.New(eventWriter, nameIt(Green, "Event"), msgStructure)
	ServiceLogger = log.New(serviceWriter, nameIt(Green, "Service"), msgStructure)
	TraceLogger = log.New(traceWriter, nameIt(White, "Trace"), msgStructure)
	AuditLogger = log.New(auditWriter, nameIt(Yellow, "Audit"), msgStructure)
	TranslationLogger = log.New(translationWriter, nameIt(Cyan, "Translation"), msgStructure)
	SecurityLogger = log.New(securityWriter, nameIt(Magenta, "Security"), msgStructure)
	DatabaseLogger = log.New(databaseWriter, nameIt(Blue, "Database"), msgStructure)
	ApiLogger = log.New(apiWriter, nameIt(Green, "API"), msgStructure)
}

func fileName(in, name string) string {
	return in + name + ".log"
}

func nameIt(colour, name string) string {
	name = strings.ToUpper(name)
	return colour + sBracket(name) + Reset
}

// sBracket adds square brackets to a string
func sBracket(s string) string {
	return "[" + s + "]"
}
func TestIt() {
	InfoLogger.Println("Starting the application...")
	InfoLogger.Println("Something noteworthy happened")
	WarningLogger.Println("There is something you should know about")
	PanicLogger.Println("Something went wrong")
	ErrorLogger.Println("Something went wrong")
	TimingLogger.Println("Timing")
	EventLogger.Println("Events")
	ServiceLogger.Println("Service")
	TraceLogger.Println("Trace")
	AuditLogger.Println("Audit")
	TranslationLogger.Println("Translation")
	SecurityLogger.Println("Security")
	DatabaseLogger.Println("Database")
	ApiLogger.Println("API")
}

func setColoursNormal() {
	Reset = colours.Reset
	Red = colours.Red
	Green = colours.Green
	Yellow = colours.Yellow
	Blue = colours.Blue
	Magenta = colours.Magenta
	Cyan = colours.Cyan
	Gray = colours.Gray
	White = colours.White
}

func setColoursWindows() {
	Reset = ""
	Red = ""
	Green = ""
	Yellow = ""
	Blue = ""
	Magenta = ""
	Cyan = ""
	Gray = ""
	White = ""
}

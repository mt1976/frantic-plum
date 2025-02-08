package logger

import (
	"io"
	"log"
	"os"
	"runtime"
	"strings"

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
	if settings.DisableLoggingGeneral() || settings.DisableAllLogging() {
		generalWriter = io.MultiWriter(io.Discard)
	}

	timingWriter := io.MultiWriter(&lumberjack.Logger{Filename: fileName(name, "timing"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	if settings.DisableLoggingTiming() || settings.DisableAllLogging() {
		timingWriter = io.MultiWriter(io.Discard)
	}

	serviceWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: fileName(name, "service"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	if settings.DisableLoggingService() || settings.DisableAllLogging() {
		serviceWriter = io.MultiWriter(io.Discard)
	}

	auditWriter := io.MultiWriter(&lumberjack.Logger{Filename: fileName(name, "audit"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	if settings.DisableLoggingAudit() || settings.DisableAllLogging() {
		auditWriter = io.MultiWriter(io.Discard)
	}

	errorWriter := io.MultiWriter(os.Stdout, os.Stderr, &lumberjack.Logger{Filename: fileName(name, "error"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	if settings.DisableAllLogging() {
		errorWriter = io.MultiWriter(io.Discard)
	}

	panicWriter := io.MultiWriter(os.Stdout, os.Stderr, &lumberjack.Logger{Filename: fileName(name, "panic"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	if settings.DisableAllLogging() {
		panicWriter = io.MultiWriter(io.Discard)
	}

	translationWriter := io.MultiWriter(&lumberjack.Logger{Filename: fileName(name, "translation"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	if settings.DisableLoggingTranslation() || settings.DisableAllLogging() {
		translationWriter = io.MultiWriter(io.Discard)
	}

	traceWriter := io.MultiWriter(&lumberjack.Logger{Filename: fileName(name, "trace"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	if settings.DisableLoggingTrace() || settings.DisableAllLogging() {
		traceWriter = io.MultiWriter(io.Discard)
	}

	warningWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: fileName(name, "warning"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	if settings.DisableLoggingWarning() || settings.DisableAllLogging() {
		warningWriter = io.MultiWriter(io.Discard)
	}

	eventWriter := io.MultiWriter(os.Stderr, &lumberjack.Logger{Filename: fileName(name, "event"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	if settings.DisableLoggingEvent() || settings.DisableAllLogging() {
		eventWriter = io.MultiWriter(io.Discard)
	}

	securityWriter := io.MultiWriter(os.Stderr, &lumberjack.Logger{Filename: fileName(name, "security"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	if settings.DisableLoggingSecurity() || settings.DisableAllLogging() {
		securityWriter = io.MultiWriter(io.Discard)
	}

	databaseWriter := io.MultiWriter(&lumberjack.Logger{Filename: fileName(name, "database"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	if settings.DisableLoggingDatabase() || settings.DisableAllLogging() {
		databaseWriter = io.MultiWriter(io.Discard)
	}

	apiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{Filename: fileName(name, "api"), MaxSize: 10, MaxBackups: 3, MaxAge: 28, Compress: true})
	if settings.DisableLoggingApi() || settings.DisableAllLogging() {
		apiWriter = io.MultiWriter(io.Discard)
	}

	//fmt.Printf("name: %v\n", name)
	//os.Exit(1)
	setColoursNormal()
	if runtime.GOOS == "windows" {
		setColoursWindows()
	}

	msgStructure := log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix

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

var hdr = "------------------------------------------------------------------------"

func Banner(class, name, action string) {
	InfoLogger.Println(hdr)
	InfoLogger.Printf("[%v] Activity=[%v] - %v", strings.ToUpper(class), name, action)
	InfoLogger.Println(hdr)
}

func ServiceBanner(class, name, action string) {
	ServiceLogger.Println(hdr)
	ServiceLogger.Printf("[%v] Activity=[%v] - %v", strings.ToUpper(class), name, action)
	ServiceLogger.Println(hdr)
}

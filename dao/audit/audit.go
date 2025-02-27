package audit

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mt1976/frantic-core/application"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/contextHandler"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var name = "Audit"
var cfg *commonConfig.Settings

type Action struct {
	code        string
	short       string
	description string
	silent      bool
}

func init() {
	cfg = commonConfig.Get()
}

func getDBVersion() int {
	// Implement the logic to get the DB version without importing the dao package
	return cfg.GetDatabase_Version()
}

type Audit struct {
	CreatedAt        time.Time
	CreatedBy        string
	CreatedOn        string
	CreatedAtDisplay string
	Updates          []AuditUpdateInfo
	DeletedAt        time.Time
	DeletedBy        string
	DeletedOn        string
	DeletedAtDisplay string
	AuditSequence    int
	DBVersion        int
	//Empty     time.Time // Convience Field - Used to avoid erros with dates.
}

type AuditUpdateInfo struct {
	UpdatedAt        time.Time
	UpdateAction     string
	UpdatedBy        string
	UpdatedOn        string
	UpdatedAtDisplay string
	UpdateNotes      string
}

const (
	AUDITMSG = "[%v] Action=[%s] At=[%v] By=[%v] On=[%v] Notes=[%v]"
)

var (
	CREATE       Action
	DELETE       Action
	UPDATE       Action
	ERASE        Action
	CLONE        Action
	NOTIFICATION Action
	SERVICE      Action
	SILENT       Action
	GRANT        Action
	REVOKE       Action
	PROCESS      Action
	IMPORT       Action
	EXPORT       Action
	GET          Action
	REPAIR       Action
	audit        Action
	BACKUP       Action
)

func init() {
	CREATE = Action{code: "NEW", description: "New Record", silent: false, short: "Create"}
	DELETE = Action{code: "DEL", description: "Delete Record", silent: false, short: "Delete"}
	UPDATE = Action{code: "UPD", description: "Update Record", silent: false, short: "Update"}
	ERASE = Action{code: "ERS", description: "Erase Record", silent: false, short: "Erase"}
	CLONE = Action{code: "CLN", description: "Clone Record", silent: false, short: "Clone"}
	NOTIFICATION = Action{code: "NTE", description: "Notification Sent", silent: false, short: "Notification"}
	SERVICE = Action{code: "SVC", description: "Service Action", silent: false, short: "Service"}
	SILENT = Action{code: "SIL", description: "Silent Action", silent: true, short: "Silent"}
	GRANT = Action{code: "GNT", description: "Grant", silent: false, short: "Grant"}
	REVOKE = Action{code: "RVK", description: "Revoke", silent: false, short: "Revoke"}
	PROCESS = Action{code: "PRC", description: "Process", silent: false, short: "Process"}
	IMPORT = Action{code: "IMP", description: "Import", silent: false, short: "Import"}
	EXPORT = Action{code: "EXP", description: "Export", silent: false, short: "Export"}
	REPAIR = Action{code: "REP", description: "Repaired", silent: false, short: "Repair"}
	audit = Action{code: "AUD", description: "Audit", silent: true, short: "Audit"}
	BACKUP = Action{code: "BAK", description: "Backup", silent: true, short: "Backup"}
}

func (a *Action) WithMessage(in string) Action {
	a.description = in
	return *a
}

func (a *Action) popMessage() string {
	message := a.description
	a.description = ""
	return message
}

func (a *Audit) Action(ctx context.Context, action Action) error {

	message := action.popMessage()
	timingMessage := fmt.Sprintf("Action=[%v] Message=[%v]", action.Code(), message)
	clock := timing.Start("Audit", actions.AUDIT.GetCode(), timingMessage)

	auditTime := time.Now()
	auditDisplay := dateHelpers.FormatAudit(auditTime)
	// auditUser := support.GetActiveUserCode()
	auditUser, err := getAuditUserCode(ctx)
	if err != nil {
		logHandler.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err)
	}
	auditHost := application.HostName()

	if auditUser == "" {
		logHandler.ErrorLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), "No Active User")
		logHandler.InfoLogger.Printf("[%v] Action=[%v] Message=[%v]", strings.ToUpper(name), action.code, message)
		os.Exit(0)
	}
	//updateAction := action

	if action.Is(CREATE) {
		a.CreatedAt = auditTime
		a.CreatedBy = auditUser
		a.CreatedOn = auditHost
		a.CreatedAtDisplay = auditDisplay
	}

	if action.Is(DELETE) {
		a.DeletedAt = auditTime
		a.DeletedBy = auditUser
		a.DeletedOn = auditHost
		a.DeletedAtDisplay = auditDisplay
	}

	if a.AuditSequence == 0 {
		a.AuditSequence = 1
	} else {
		a.AuditSequence++
	}

	update := AuditUpdateInfo{}

	update.UpdatedAt = auditTime
	update.UpdatedBy = auditUser
	update.UpdatedOn = auditHost
	update.UpdatedAtDisplay = auditDisplay
	update.UpdateAction = action.code
	update.UpdateNotes = message
	// a.DBVersion = dao.Version
	a.DBVersion = getDBVersion()
	if !(action.Is(SERVICE) || action.Is(SILENT) || action.IsSilent()) {
		a.Updates = append(a.Updates, update)
	}

	a.DBVersion = getDBVersion()

	logHandler.AuditLogger.Printf(AUDITMSG, strings.ToUpper(name), action.code, auditDisplay, auditUser, auditHost, message)
	clock.Stop(1)
	return nil
}

func (a *Audit) Spew() error {
	// Spew the Audit Data
	noAudit := len(a.Updates)
	//logger.InfoLogger.Printf(" No Updates=[%v]", noAudit)
	if noAudit > 0 {
		for i := 0; i < noAudit; i++ {
			upd := a.Updates[i]
			logHandler.TraceLogger.Printf(AUDITMSG, strings.ToUpper(name), upd.UpdateAction, upd.UpdatedAtDisplay, upd.UpdatedBy, upd.UpdatedOn, upd.UpdateNotes)
		}
	}
	return nil
}

func (a *Action) Is(inAction Action) bool {
	return a.code == inAction.code
}

func (a *Action) IsSilent() bool {
	return a.silent
}

func (a *Action) Description() string {
	return a.description
}

func (a *Action) ShortNameRaw() string {
	return a.short
}

func (a *Action) ShortName() string {
	return strings.ToUpper(a.ShortNameRaw())
}

func (a *Action) Text() string {
	return strings.ToUpper(a.code)
}

func (a *Action) SetMessage(in string) {
	a.description = in
}

func (a *Action) GetMessage() string {
	return a.description
}

func (a *Action) SetText(in string) {
	a.code = in
}

func (a *Action) Code() string {
	return a.code
}

func getAuditUserCode(ctx context.Context) (string, error) {
	// Implement the logic to get the user without importing the dao package
	defaultUser := cfg.GetServiceUser_UserCode()
	if ctx == context.TODO() || ctx == nil {
		return defaultUser, nil
	}

	// Get the current user from the context
	sessionUser := contextHandler.GetUserCode(ctx)
	//ctx.Value(cfg.GetSecuritySessionKey_UserCode())
	if sessionUser != "" {
		return sessionUser, nil
	}
	return defaultUser, commonErrors.ErrContextCannotGetUserCode
}

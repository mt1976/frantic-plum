package commonConfig

import (
	"strconv"
	"strings"
)

const (
	MODE_DEVELOPMENT = "development"
	MODE_PRODUCTION  = "production"
	MODE_TEST
)

func (s *Settings) GetServerPort() int {
	return s.Server.Port
}

func (s *Settings) GetServerPortAsString() string {
	a := s.Server.Port
	return strconv.Itoa(a)
}

func (s *Settings) GetApplicationName() string {
	return s.Application.Name
}

func (s *Settings) GetApplicationPrefix() string {
	return s.Application.Prefix
}

func (s *Settings) GetApplicationHomePath() string {
	return s.Application.Home
}

func (s *Settings) GetApplicationDescription() string {
	return s.Application.Description
}

func (s *Settings) GetServerProtocol() string {
	return s.Server.Protocol
}

func (s *Settings) IsApplicationMode(inMode string) bool {
	// If first three chars of environment are "dev" then return "development"
	if strings.ToLower(s.Server.Environment[:3]) == strings.ToLower(inMode[:3]) {
		return true
	}
	return false
}

func (s *Settings) GetLogoPath() string {
	return s.Assets.Logo
}

func (s *Settings) GetFaviconPath() string {
	return s.Assets.Favicon
}

func (s *Settings) GetDateTimeFormat() string {
	return s.Dates.DateTimeFormat
}

func (s *Settings) GetDateFormat() string {
	return s.Dates.DateFormat
}

func (s *Settings) GetTimeFormat() string {
	return s.Dates.TimeFormat
}

func (s *Settings) DateFormatBackup() string {
	return s.Dates.Backup
}

func (s *Settings) GetDateFormatForBackupDirectory() string {
	return s.Dates.BackupFolder
}

func (s *Settings) GetHumanReadableDateFormat() string {
	return s.Dates.Human
}

func (s *Settings) GetDateFormatDMY2() string {
	return s.Dates.DMY2
}

func (s *Settings) GetDateFormatYMD() string {
	return s.Dates.YMD
}

func (s *Settings) GetInternalDateFormat() string {
	return s.Dates.Internal
}

func (s *Settings) GetMaxHistoryEntries() int {
	return s.History.MaxEntries
}

func (s *Settings) GetDisplayDelimiter() string {
	return s.Display.Delimiter
}

func (s *Settings) GetApplicationEnvironment() string {
	return s.Application.Environment
}

func (s *Settings) GetApplicationVersion() string {
	return s.Application.Version
}

func (s *Settings) GetApplicationReleaseDate() string {
	return s.Application.ReleaseDate
}

func (s *Settings) GetApplicationCopyright() string {
	return s.Application.Copyright
}

func (s *Settings) GetApplicationAuthor() string {
	return s.Application.Author
}

func (s *Settings) GetMessageTypeKey() string {
	return s.Message.TypeKey
}

func (s *Settings) GetMessageTitleKey() string {
	return s.Message.TitleKey
}

func (s *Settings) GetMessageContentKey() string {
	return s.Message.ContentKey
}

func (s *Settings) GetMessageActionKey() string {
	return s.Message.ActionKey
}

func (s *Settings) SEP() string {
	return s.Display.Delimiter
}

func (s *Settings) Delimiter() string {
	return s.SEP()
}

func (s *Settings) GetApplicationLocale() string {
	return s.Application.Locale
}

func (s *Settings) GetServerHost() string {
	return s.Server.Host
}

func (s *Settings) GetTranslationServerHost() string {
	return s.Translation.Host
}

func (s *Settings) GetTranslationServerPort() int {
	return s.Translation.Port
}

func (s *Settings) GetTranslationLocale() string {
	return s.Translation.Locale
}

func (s *Settings) GetTranslationServerProtocol() string {
	return s.Translation.Protocol
}

func (s *Settings) GetTranslationServerPortAsString() string {
	return strconv.Itoa(s.Translation.Port)
}

// func (s *Settings) MaxEntries() int {
// 	return s.History.MaxEntries
// }

func (s *Settings) GetPushoverUserKey() string {
	return s.Pushover.UserKey
}

func (s *Settings) GetPushoverToken() string {
	return s.Pushover.APIToken
}

func (s *Settings) GetStatus_Unknown() string {
	if s.Status.UNKNOWN == "" {
		return "UNKN"
	}
	return s.Status.UNKNOWN
}

func (s *Settings) GetStatus_Online() string {
	if s.Status.ONLINE == "" {
		return "ONLN"
	}
	return s.Status.ONLINE
}

func (s *Settings) GetStatus_Offline() string {
	if s.Status.OFFLINE == "" {
		return "OFLN"
	}
	return s.Status.OFFLINE
}

func (s *Settings) GetStatus_Error() string {
	if s.Status.ERROR == "" {
		return "ERRO"
	}
	return s.Status.ERROR
}

func (s *Settings) GetStatus_Warning() string {
	if s.Status.WARNING == "" {
		return "WARN"
	}
	return s.Status.WARNING
}

func (s *Settings) GetSecurityServiceUserName() string {
	return s.Security.ServiceUserName
}

func (s *Settings) GetSecurityServiceUserCode() string {
	return s.Security.ServiceUserCode
}

func (s *Settings) GetSecuritySessionExpiryPeriodKey() int {
	return s.Security.SessionExpiry
}

func (s *Settings) GetSecuritySessionUserCodeKey() string {
	return s.Security.SessionUserCodeKey
}

func (s *Settings) GetSecuritySessionUserIDKey() string {
	return s.Security.SessionUserIDKey
}

func (s *Settings) GetSecuritySessionExpiryKey() string {
	return s.Security.SessionExpiryKey
}

func (s *Settings) GetSecuritySessionTokenKey() string {
	return s.Security.SessionTokenKey
}

func (s *Settings) GetSecuritySessionKeyName() string {
	return s.Security.SessionKeyName
}

func (s *Settings) GetValidOrigins() []string {
	var origins []string
	for _, v := range s.AllowedOrigins {
		if v.Name != "" {
			origins = append(origins, v.Name)
		}
	}
	return origins
}

func (s *Settings) GetValidHosts() []struct {
	Name string "toml:\"name\""
	FQDN string "toml:\"fqdn\""
	IP   string "toml:\"ip\""
	Zone string "toml:\"zone\""
} {
	return s.Hosts
}

func (s *Settings) GetLocales() []struct {
	Key  string "toml:\"key\""
	Name string "toml:\"name\""
} {
	return s.Locales
}

func (s *Settings) IsValidLocale(in string) bool {
	for _, v := range s.Locales {
		if v.Key == in {
			return true
		}
	}
	return false
}

func (s *Settings) IsValidOrigin(in string) bool {
	for _, v := range s.AllowedOrigins {
		if v.Name == in {
			return true
		}
	}
	return false
}

func (s *Settings) GetLocaleName(in string) string {
	for _, v := range s.Locales {
		if v.Key == in {
			return v.Name
		}
	}
	return ""
}

func isTrueFalse(s string) bool {
	logTrue := "true"
	//var logFalse = "false"
	// We only disable the logging if the value is "true"
	if s == logTrue {
		return true
	}
	return false
}

func (s *Settings) GetEmailHost() string {
	return s.Email.Host
}

func (s *Settings) GetEmailPort() int {
	return s.Email.Port
}

func (s *Settings) GetEmailUser() string {
	return s.Email.User
}

func (s *Settings) GetEmailPassword() string {
	return s.Email.Password
}

func (s *Settings) GetEmailSender() string {
	return s.Email.From
}

func (s *Settings) GetEmailFooter() string {
	return s.Email.Footer
}

func (s *Settings) GetEmailPortAsString() string {
	return strconv.Itoa(s.Email.Port)
}

func (s *Settings) GetAdminEmail() string {
	return s.Email.Admin
}

func (s *Settings) IsGeneralLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.General)

}

func (s *Settings) IsTimingLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Timing)
}

func (s *Settings) IsServiceLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Service)
}

func (s *Settings) IsAuditLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Audit)
}

func (s *Settings) IsSecurityLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Security)
}

func (s *Settings) IsDatabaseLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Database)
}

func (s *Settings) IsTranslationLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Translation)
}

func (s *Settings) IsTraceLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Trace)
}

func (s *Settings) IsWarningLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Warning)
}

func (s *Settings) IsEventLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Event)
}

func (s *Settings) IsApiLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Api)
}

func (s *Settings) IsImportLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Import)
}

func (s *Settings) IsExportLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Export)
}

func (s *Settings) AreAllLogsDisabled() bool {
	return isTrueFalse(s.Logging.Disable.All)
}

func (s *Settings) GetLogsMaxSize() int {
	a, _ := strconv.Atoi(s.Logging.Defaults.MaxSize)
	if a == 0 {
		a = 10
	}
	return a
}

func (s *Settings) GetLogsMaxBackups() int {
	a, _ := strconv.Atoi(s.Logging.Defaults.MaxBackups)
	return a
}

func (s *Settings) GetLogsMaxAge() int {
	a, _ := strconv.Atoi(s.Logging.Defaults.MaxAge)
	if a == 0 {
		a = 30
	}
	return a
}

func (s *Settings) IsLogCompressionEnabled() bool {
	return isTrueFalse(s.Logging.Defaults.Compress)
}

func (s *Settings) GetDatabaseVersion() int {
	return s.Database.Version
}

func (s *Settings) GetDatabaseType() string {
	return s.Database.Type
}

func (s *Settings) GetDatabaseName() string {
	return s.Database.Name
}

func (s *Settings) GetDatabasePath() string {
	return s.Database.Path
}

func (s *Settings) GetDatabaseHost() string {
	return s.Database.Host
}

func (s *Settings) GetDatabasePort() int {
	return s.Database.Port
}

func (s *Settings) GetDatabaseUser() string {
	return s.Database.User
}

func (s *Settings) GetDatabasePassword() string {
	return s.Database.Pass
}

func (s *Settings) GetDatabasePortAsString() string {
	return strconv.Itoa(s.Database.Port)
}

func (s *Settings) GetDatabasePoolSize() int {
	if s.Database.PoolSize == 0 {
		return 10 // Default to 10 connections
	}
	return s.Database.PoolSize
}

func (s *Settings) GetDatabaseTimeout() int {
	if s.Database.Timeout == 0 {
		return 30 // Default to 30 seconds
	}
	return s.Database.Timeout
}

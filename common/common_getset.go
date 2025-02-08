package common

import (
	"strconv"
	"strings"
)

const (
	MODE_DEVELOPMENT = "development"
	MODE_PRODUCTION  = "production"
	MODE_TEST
)

func (s *Settings) ApplicationPort() int {
	return s.Server.Port
}

func (s *Settings) ApplicationPortString() string {
	a := s.Server.Port
	return strconv.Itoa(a)
}

func (s *Settings) ApplicationName() string {
	return s.Application.Name
}

func (s *Settings) ApplicationPrefix() string {
	return s.Application.Prefix
}

func (s *Settings) ApplicationHome() string {
	return s.Application.Home
}

func (s *Settings) ApplicationDescription() string {
	return s.Application.Description
}

func (s *Settings) ServerProtocol() string {
	return s.Server.Protocol
}

func (s *Settings) ApplicationModeIs(inMode string) bool {
	// If first three chars of environment are "dev" then return "development"
	if strings.ToLower(s.Server.Environment[:3]) == strings.ToLower(inMode[:3]) {
		return true
	}
	return false
}

func (s *Settings) AssetsLogo() string {
	return s.Assets.Logo
}

func (s *Settings) AssetsFavicon() string {
	return s.Assets.Favicon
}

func (s *Settings) DateFormatDateTime() string {
	return s.Dates.DateTimeFormat
}

func (s *Settings) DateFormatDate() string {
	return s.Dates.DateFormat
}

func (s *Settings) DateFormatTime() string {
	return s.Dates.TimeFormat
}

func (s *Settings) DateFormatBackup() string {
	return s.Dates.Backup
}

func (s *Settings) DateFormatBackupFolder() string {
	return s.Dates.BackupFolder
}

func (s *Settings) DateFormatHuman() string {
	return s.Dates.Human
}

func (s *Settings) DateFormatDMY2() string {
	return s.Dates.DMY2
}

func (s *Settings) DateFormatYMD() string {
	return s.Dates.YMD
}

func (s *Settings) DateFormatInternal() string {
	return s.Dates.Internal
}

func (s *Settings) HistoryMaxEntries() int {
	return s.History.MaxEntries
}

func (s *Settings) DisplayDelimiter() string {
	return s.Display.Delimiter
}

func (s *Settings) ApplicationEnvironment() string {
	return s.Application.Environment
}

func (s *Settings) ApplicationVersion() string {
	return s.Application.Version
}

func (s *Settings) ApplicationReleaseDate() string {
	return s.Application.ReleaseDate
}

func (s *Settings) ApplicationCopyright() string {
	return s.Application.Copyright
}

func (s *Settings) ApplicationAuthor() string {
	return s.Application.Author
}

func (s *Settings) MessageTypeKey() string {
	return s.Message.TypeKey
}

func (s *Settings) MessageTitleKey() string {
	return s.Message.TitleKey
}

func (s *Settings) MessageContentKey() string {
	return s.Message.ContentKey
}

func (s *Settings) MessageActionKey() string {
	return s.Message.ActionKey
}

func (s *Settings) SEP() string {
	return s.Display.Delimiter
}

func (s *Settings) Delimiter() string {
	return s.SEP()
}

func (s *Settings) ApplicationLocale() string {
	return s.Application.Locale
}

func (s *Settings) ApplicationHost() string {
	return s.Server.Host
}

func (s *Settings) TranslationHost() string {
	return s.Translation.Host
}

func (s *Settings) TranslationPort() int {
	return s.Translation.Port
}

func (s *Settings) TranslationLocale() string {
	return s.Translation.Locale
}

func (s *Settings) TranslationProtocol() string {
	return s.Translation.Protocol
}

func (s *Settings) TranslationPortString() string {
	return strconv.Itoa(s.Translation.Port)
}

func (s *Settings) MaxEntries() int {
	return s.History.MaxEntries
}

func (s *Settings) PushoverUserKey() string {
	return s.Pushover.UserKey
}

func (s *Settings) PushoverAPIToken() string {
	return s.Pushover.APIToken
}

func (s *Settings) StatusUNKNOWN() string {
	if s.Status.UNKNOWN == "" {
		return "UNKN"
	}
	return s.Status.UNKNOWN
}

func (s *Settings) StatusONLINE() string {
	if s.Status.ONLINE == "" {
		return "ONLN"
	}
	return s.Status.ONLINE
}

func (s *Settings) StatusOFFLINE() string {
	if s.Status.OFFLINE == "" {
		return "OFLN"
	}
	return s.Status.OFFLINE
}

func (s *Settings) StatusERROR() string {
	if s.Status.ERROR == "" {
		return "ERRO"
	}
	return s.Status.ERROR
}

func (s *Settings) StatusWARNING() string {
	if s.Status.WARNING == "" {
		return "WARN"
	}
	return s.Status.WARNING
}

func (s *Settings) SecurityServiceUserName() string {
	return s.Security.ServiceUserName
}

func (s *Settings) SecurityServiceUserCode() string {
	return s.Security.ServiceUserCode
}

func (s *Settings) SecuritySessionExpiryPeriod() int {
	return s.Security.SessionExpiry
}

func (s *Settings) SecuritySessionUserCodeKey() string {
	return s.Security.SessionUserCodeKey
}

func (s *Settings) SecuritySessionUserIDKey() string {
	return s.Security.SessionUserIDKey
}

func (s *Settings) SecuritySessionExpiryKey() string {
	return s.Security.SessionExpiryKey
}

func (s *Settings) SecuritySessionTokenKey() string {
	return s.Security.SessionTokenKey
}

func (s *Settings) SecuritySessionKey() string {
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

func (s *Settings) EmailHost() string {
	return s.Email.Host
}

func (s *Settings) EmailPort() int {
	return s.Email.Port
}

func (s *Settings) EmailUser() string {
	return s.Email.User
}

func (s *Settings) EmailPassword() string {
	return s.Email.Password
}

func (s *Settings) EmailFrom() string {
	return s.Email.From
}

func (s *Settings) EmailFooter() string {
	return s.Email.Footer
}

func (s *Settings) EmailPortString() string {
	return strconv.Itoa(s.Email.Port)
}

func (s *Settings) EmailAdmin() string {
	return s.Email.Admin
}

var logTrue = "true"
var logFalse = "false"

func (s *Settings) DisableLoggingGeneral() bool {
	if s.DisableLogging.General == logTrue {
		return true
	}
	return false
}

func (s *Settings) DisableLoggingTiming() bool {
	return isTrueFalse(s.DisableLogging.General)
}

func isTrueFalse(s string) bool {
	// We only disable the logging if the value is "true"
	if s == logTrue {
		return true
	}
	return false
}

func (s *Settings) DisableLoggingService() bool {
	return isTrueFalse(s.DisableLogging.Service)
}

func (s *Settings) DisableLoggingAudit() bool {
	return isTrueFalse(s.DisableLogging.Audit)
}

func (s *Settings) DisableLoggingSecurity() bool {
	return isTrueFalse(s.DisableLogging.Security)
}

func (s *Settings) DisableLoggingDatabase() bool {
	return isTrueFalse(s.DisableLogging.Database)
}

func (s *Settings) DisableLoggingTranslation() bool {
	return isTrueFalse(s.DisableLogging.Translation)
}

func (s *Settings) DisableLoggingTrace() bool {
	return isTrueFalse(s.DisableLogging.Trace)
}

func (s *Settings) DisableLoggingWarning() bool {
	return isTrueFalse(s.DisableLogging.Warning)
}

func (s *Settings) DisableLoggingEvent() bool {
	return isTrueFalse(s.DisableLogging.Event)
}

func (s *Settings) DisableLoggingApi() bool {
	return isTrueFalse(s.DisableLogging.Api)
}

func (s *Settings) DisableAllLogging() bool {
	return isTrueFalse(s.DisableLogging.All)
}

func (s *Settings) LoggingMaxSize() int {
	a, _ := strconv.Atoi(s.Logging.MaxSize)
	return a
}

func (s *Settings) LoggingMaxBackups() int {
	a, _ := strconv.Atoi(s.Logging.MaxBackups)
	return a
}

func (s *Settings) LoggingMaxAge() int {
	a, _ := strconv.Atoi(s.Logging.MaxAge)
	return a
}

func (s *Settings) LoggingCompress() bool {
	return isTrueFalse(s.Logging.Compress)
}

package config

import (
	"strconv"
	"strings"
)

const (
	MODE_DEVELOPMENT = "development"
	MODE_PRODUCTION  = "production"
	MODE_TEST
)

func (d *Configuration) ApplicationPort() int {
	return d.Server.Port
}

func (d *Configuration) ApplicationPortString() string {
	a := d.Server.Port
	return strconv.Itoa(a)
}

func (d *Configuration) ApplicationName() string {
	return d.Application.Name
}

func (d *Configuration) ApplicationPrefix() string {
	return d.Application.Prefix
}

func (d *Configuration) ApplicationHome() string {
	return d.Application.Home
}

func (d *Configuration) ApplicationDescription() string {
	return d.Application.Description
}

func (d *Configuration) ServerProtocol() string {
	return d.Server.Protocol
}

func (d *Configuration) ApplicationModeIs(inMode string) bool {
	// If first three chars of environment are "dev" then return "development"
	if strings.ToLower(d.Server.Environment[:3]) == strings.ToLower(inMode[:3]) {
		return true
	}
	return false
}

func (d *Configuration) AssetsLogo() string {
	return d.Assets.Logo
}

func (d *Configuration) AssetsFavicon() string {
	return d.Assets.Favicon
}

func (d *Configuration) DateFormatDateTime() string {
	return d.Dates.DateTimeFormat
}

func (d *Configuration) DateFormatDate() string {
	return d.Dates.DateFormat
}

func (d *Configuration) DateFormatTime() string {
	return d.Dates.TimeFormat
}

func (d *Configuration) DateFormatBackup() string {
	return d.Dates.Backup
}

func (d *Configuration) DateFormatBackupFolder() string {
	return d.Dates.BackupFolder
}

func (d *Configuration) DateFormatHuman() string {
	return d.Dates.Human
}

func (d *Configuration) DateFormatDMY2() string {
	return d.Dates.DMY2
}

func (d *Configuration) DateFormatYMD() string {
	return d.Dates.YMD
}

func (d *Configuration) DateFormatInternal() string {
	return d.Dates.Internal
}

func (d *Configuration) HistoryMaxEntries() int {
	return d.History.MaxEntries
}

func (d *Configuration) DisplayDelimiter() string {
	return d.Display.Delimiter
}

func (d *Configuration) ApplicationEnvironment() string {
	return d.Application.Environment
}

func (d *Configuration) ApplicationVersion() string {
	return d.Application.Version
}

func (d *Configuration) ApplicationReleaseDate() string {
	return d.Application.ReleaseDate
}

func (d *Configuration) ApplicationCopyright() string {
	return d.Application.Copyright
}

func (d *Configuration) ApplicationAuthor() string {
	return d.Application.Author
}

func (d *Configuration) MessageTypeKey() string {
	return d.Message.TypeKey
}

func (d *Configuration) MessageTitleKey() string {
	return d.Message.TitleKey
}

func (d *Configuration) MessageContentKey() string {
	return d.Message.ContentKey
}

func (d *Configuration) MessageActionKey() string {
	return d.Message.ActionKey
}

func (d *Configuration) SEP() string {
	return d.Display.Delimiter
}

func (d *Configuration) Delimiter() string {
	return d.SEP()
}

func (d *Configuration) ApplicationLocale() string {
	return d.Application.Locale
}

func (d *Configuration) ApplicationHost() string {
	return d.Server.Host
}

func (d *Configuration) TranslationHost() string {
	return d.Translation.Host
}

func (d *Configuration) TranslationPort() int {
	return d.Translation.Port
}

func (d *Configuration) TranslationLocale() string {
	return d.Translation.Locale
}

func (d *Configuration) TranslationProtocol() string {
	return d.Translation.Protocol
}

func (d *Configuration) TranslationPortString() string {
	return strconv.Itoa(d.Translation.Port)
}

func (d *Configuration) MaxEntries() int {
	return d.History.MaxEntries
}

func (d *Configuration) PushoverUserKey() string {
	return d.Pushover.UserKey
}

func (d *Configuration) PushoverAPIToken() string {
	return d.Pushover.APIToken
}

func (d *Configuration) StatusUNKNOWN() string {
	if d.Status.UNKNOWN == "" {
		return "UNKN"
	}
	return d.Status.UNKNOWN
}

func (d *Configuration) StatusONLINE() string {
	if d.Status.ONLINE == "" {
		return "ONLN"
	}
	return d.Status.ONLINE
}

func (d *Configuration) StatusOFFLINE() string {
	if d.Status.OFFLINE == "" {
		return "OFLN"
	}
	return d.Status.OFFLINE
}

func (d *Configuration) StatusERROR() string {
	if d.Status.ERROR == "" {
		return "ERRO"
	}
	return d.Status.ERROR
}

func (d *Configuration) StatusWARNING() string {
	if d.Status.WARNING == "" {
		return "WARN"
	}
	return d.Status.WARNING
}

func (d *Configuration) SecurityServiceUserName() string {
	return d.Security.ServiceUserName
}

func (d *Configuration) SecurityServiceUserCode() string {
	return d.Security.ServiceUserCode
}

func (d *Configuration) SecuritySessionExpiryPeriod() int {
	return d.Security.SessionExpiry
}

func (d *Configuration) SecuritySessionUserCodeKey() string {
	return d.Security.SessionUserCodeKey
}

func (d *Configuration) SecuritySessionUserIDKey() string {
	return d.Security.SessionUserIDKey
}

func (d *Configuration) SecuritySessionExpiryKey() string {
	return d.Security.SessionExpiryKey
}

func (d *Configuration) SecuritySessionTokenKey() string {
	return d.Security.SessionTokenKey
}

func (d *Configuration) SecuritySessionKey() string {
	return d.Security.SessionKeyName
}

func (d *Configuration) GetValidOrigins() []string {
	var origins []string
	for _, v := range d.AllowedOrigins {
		if v.Name != "" {
			origins = append(origins, v.Name)
		}
	}
	return origins
}

func (d *Configuration) GetValidHosts() []struct {
	Name string "toml:\"name\""
	FQDN string "toml:\"fqdn\""
	IP   string "toml:\"ip\""
	Zone string "toml:\"zone\""
} {
	return d.Hosts
}

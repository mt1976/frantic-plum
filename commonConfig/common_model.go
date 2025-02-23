package commonConfig

type Settings struct {
	Application struct {
		Name        string `toml:"name"`
		Prefix      string `toml:"prefix"`
		Home        string `toml:"home"`
		Description string `toml:"description"`
		Version     string `toml:"version"`
		Environment string `toml:"environment"`
		ReleaseDate string `toml:"releaseDate"`
		Copyright   string `toml:"copyright"`
		Author      string `toml:"author"`
		License     string `toml:"license"`
		Locale      string `toml:"locale"`
	} `toml:"Application"`
	Database struct {
		Version  int    `toml:"version"`
		Type     string `toml:"type"`
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Name     string `toml:"name"`
		User     string `toml:"user"`
		Pass     string `toml:"pass"`
		Path     string `toml:"path"`
		PoolSize int    `toml:"poolSize"`
		Timeout  int    `toml:"timeout"`
	} `toml:"Database"`
	Server struct {
		Host        string `toml:"host"`
		Port        int    `toml:"port"`
		Protocol    string `toml:"protocol"`
		Environment string `toml:"environment"`
	} `toml:"Server"`
	Translation struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Locale   string `toml:"locale"`
		Protocol string `toml:"protocol"`
	} `toml:"Translation"`
	Assets struct {
		Logo    string `toml:"logo"`
		Favicon string `toml:"favicon"`
	} `toml:"Assets"`
	Dates struct {
		DateTimeFormat string `toml:"dateTimeFormat"`
		DateFormat     string `toml:"dateFormat"`
		TimeFormat     string `toml:"timeFormat"`
		Backup         string `toml:"backup"`
		BackupFolder   string `toml:"backupFolder"`
		Human          string `toml:"human"`
		DMY2           string `toml:"dmy2"`
		YMD            string `toml:"ymd"`
		Internal       string `toml:"internal"`
	} `toml:"Dates"`
	History struct {
		MaxEntries int `toml:"maxEntries"`
	} `toml:"History"`
	Hosts []struct {
		Name string `toml:"name"`
		FQDN string `toml:"fqdn"`
		IP   string `toml:"ip"`
		Zone string `toml:"zone"`
	} `toml:"Hosts"`
	Security struct {
		SessionKeyName     string `toml:"sessionKey"`
		SessionUserIDKey   string `toml:"sessionUserIDKey"`
		SessionUserCodeKey string `toml:"sessionUserCodeKey"`
		SessionTokenKey    string `toml:"sessionTokenKey"`
		SessionExpiry      int    `toml:"sessionExpiry"`
		ServiceUserName    string `toml:"serviceUserName"`
		ServiceUserCode    string `toml:"serviceUserCode"`
		SessionExpiryKey   string `toml:"sessionExpiryKey"`
	} `toml:"Security"`
	Message struct {
		TypeKey    string `toml:"typeKey"`
		TitleKey   string `toml:"titleKey"`
		ContentKey string `toml:"contentKey"`
		ActionKey  string `toml:"actionKey"`
	} `toml:"Message"`
	Display struct {
		Delimiter string `toml:"delim"`
	} `toml:"Display"`
	Pushover struct {
		UserKey  string `toml:"userKey"`
		APIToken string `toml:"apiToken"`
	} `toml:"Pushover"`
	Status struct {
		UNKNOWN string `toml:"unknown"`
		ONLINE  string `toml:"online"`
		OFFLINE string `toml:"offline"`
		ERROR   string `toml:"error"`
		WARNING string `toml:"warning"`
	} `toml:"Status"`
	AllowedOrigins []struct {
		Name string `toml:"name"`
	} `toml:"Origins"`
	Locales []struct {
		Key  string `toml:"key"`
		Name string `toml:"name"`
	} `toml:"Locales"`
	Email struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		From     string `toml:"from"`
		Footer   string `toml:"footer"`
		Admin    string `toml:"admin"`
	} `toml:"Email"`
	Logging struct {
		Disable struct {
			General        string `toml:"general"`
			Timing         string `toml:"timing"`
			Service        string `toml:"service"`
			Audit          string `toml:"audit"`
			Translation    string `toml:"translation"`
			Trace          string `toml:"trace"`
			Warning        string `toml:"warning"`
			Event          string `toml:"event"`
			Security       string `toml:"security"`
			Database       string `toml:"database"`
			Api            string `toml:"api"`
			Import         string `toml:"import"`
			Export         string `toml:"export"`
			Communications string `toml:"comms"`
			All            string `toml:"all"`
		} `toml:"disable"`
		Defaults struct {
			MaxSize    string `toml:"maxSize"`
			MaxBackups string `toml:"maxBackups"`
			MaxAge     string `toml:"maxAge"`
			Compress   string `toml:"compress"`
		} `toml:"Defaults"`
	} `toml:"Logging"`
}

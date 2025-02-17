package maintenance

import (
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/logHandler"
	trnsl8r "github.com/mt1976/trnsl8r_connect"
)

var domain = "Maintenance"
var translation trnsl8r.Request
var cfg *commonConfig.Settings

func init() {

	logHandler.InfoLogger.Println("Text - Initialising")

	cfg = commonConfig.Get()
	err := error(nil)
	translation, err = trnsl8r.NewRequest().WithProtocol(cfg.GetTranslationServerProtocol()).WithHost(cfg.GetTranslationServerHost()).WithPort(cfg.GetTranslationServerPort()).WithLogger(logHandler.TranslationLogger).FromOrigin(cfg.GetApplicationName()).WithFilter(trnsl8r.LOCALE, cfg.GetApplicationLocale())
	if err != nil {
		logHandler.ErrorLogger.Println(err.Error())
		return
	}

	logHandler.InfoLogger.Println("Text - Initialised")

}

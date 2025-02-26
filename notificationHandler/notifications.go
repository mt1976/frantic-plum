package notificationHandler

import (
	"fmt"
	"strings"
	"time"

	"github.com/gregdel/pushover"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
)

var domain = "Notification"

func Send(inMessage, inTitle string, key int) error {
	clock := timing.Start(domain, actions.MESSAGE.GetCode(), "Pushover Notification")
	logHandler.CommunicationsLogger.Printf("[%v] Pushover - Sending...", domain)

	cfg := commonConfig.Get()

	poUserKey := cfg.GetCommunicationsPushover_UserKey()
	poAPIKey := cfg.GetCommunicationsPushover_APIToken()

	if cfg.IsApplicationMode(commonConfig.MODE_TEST) {
		if poUserKey == "" || poAPIKey == "" {
			poAPIKey = "autd5u19nczbs5v6zq2i7afpzjpe2v"
			poUserKey = "uyosdopsu9wxxo7b264bmnnhbfz8nj"
		}
	}

	logHandler.CommunicationsLogger.Printf("[%v] Api Token=[%v] User Key=[%v]", strings.ToUpper(domain), poAPIKey, poUserKey)

	if poUserKey == "" || poAPIKey == "" {
		logHandler.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(domain), "Pushover User Key or API Token not set, message not sent")
		return nil
	}

	app := pushover.New(poAPIKey)
	recipient := pushover.NewRecipient(poUserKey)
	var inCallbackUrl string

	if key != 0 {
		//inCallbackUrl = support.Application.BaseURL() + "view/" + key
		inCallbackUrl = fmt.Sprintf("http://%v:%v/view/%v", cfg.GetServer_Host(), cfg.GetServer_Port(), key)
	} else {
		inCallbackUrl = fmt.Sprintf("http://%v:%v/dashboard/", cfg.GetServer_Host(), cfg.GetServer_Port())
	}

	message := &pushover.Message{
		Message:    inMessage,
		Title:      inTitle,
		Priority:   pushover.PriorityNormal,
		URL:        inCallbackUrl,
		URLTitle:   cfg.GetApplication_Name(),
		Timestamp:  time.Now().Unix(),
		Retry:      60 * time.Second,
		Expire:     time.Hour,
		DeviceName: "",
		//CallbackURL: "http://yourapp.com/callback",
		CallbackURL: inCallbackUrl,
		Sound:       pushover.SoundCosmic,
	}
	//Spew(*message)

	logHandler.CommunicationsLogger.Printf("[%v] Pusover - Message [Title=%v] [Message=%v]", domain, message.Title, message.Message)

	_, err := app.SendMessage(message, recipient)
	if err != nil {
		logHandler.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(domain), err.Error())
		return commonErrors.WrapNotificationError(err)
	}

	clock.Stop(1)

	return nil
}

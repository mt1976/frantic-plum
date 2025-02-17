package notificationHandler

import (
	"fmt"
	"strings"
	"time"

	"github.com/gregdel/pushover"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/logHandler"
)

var name = "Notifications"

func Send(inMessage, inTitle string, key int) error {

	set := commonConfig.Get()

	poUserKey := set.GetPushoverUserKey()
	poAPIKey := set.GetPushoverToken()

	if set.IsApplicationMode(commonConfig.MODE_TEST) {
		if poUserKey == "" || poAPIKey == "" {
			poAPIKey = "autd5u19nczbs5v6zq2i7afpzjpe2v"
			poUserKey = "uyosdopsu9wxxo7b264bmnnhbfz8nj"
		}
	}

	logHandler.InfoLogger.Printf("[%v] Api Token=[%v] User Key=[%v]", strings.ToUpper(name), poAPIKey, poUserKey)

	if poUserKey == "" || poAPIKey == "" {
		logHandler.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), "Pushover User Key or API Token not set, message not sent")
		return nil
	}

	app := pushover.New(poAPIKey)
	recipient := pushover.NewRecipient(poUserKey)
	var inCallbackUrl string

	if key != 0 {
		//inCallbackUrl = support.Application.BaseURL() + "view/" + key
		inCallbackUrl = fmt.Sprintf("http://%v:%v/view/%v", set.GetServerHost(), set.GetServerPort(), key)
	} else {
		inCallbackUrl = fmt.Sprintf("http://%v:%v/dashboard/", set.GetServerHost(), set.GetServerPort())
	}

	message := &pushover.Message{
		Message:    inMessage,
		Title:      inTitle,
		Priority:   pushover.PriorityNormal,
		URL:        inCallbackUrl,
		URLTitle:   set.GetApplicationName(),
		Timestamp:  time.Now().Unix(),
		Retry:      60 * time.Second,
		Expire:     time.Hour,
		DeviceName: "",
		//CallbackURL: "http://yourapp.com/callback",
		CallbackURL: inCallbackUrl,
		Sound:       pushover.SoundCosmic,
	}
	//Spew(*message)

	logHandler.EventLogger.Printf("[%v] Message Title=[%v] Message=[%v]", strings.ToUpper(name), message.Title, message.Message)

	_, err := app.SendMessage(message, recipient)
	if err != nil {
		logHandler.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		return commonErrors.WrapNotificationError(err)
	}

	return nil
}

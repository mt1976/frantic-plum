package notifications

import (
	"fmt"
	"strings"
	"time"

	"github.com/gregdel/pushover"
	"github.com/mt1976/frantic-plum/common"
	"github.com/mt1976/frantic-plum/logger"
)

var name = "Notifications"

func Send(inMessage, inTitle string, key int) error {

	set := common.Get()

	poUserKey := set.PushoverUserKey()
	poAPIKey := set.PushoverAPIToken()

	if set.ApplicationModeIs(common.MODE_TEST) {
		if poUserKey == "" || poAPIKey == "" {
			poAPIKey = "autd5u19nczbs5v6zq2i7afpzjpe2v"
			poUserKey = "uyosdopsu9wxxo7b264bmnnhbfz8nj"
		}
	}

	logger.InfoLogger.Printf("[%v] Api Token=[%v]", strings.ToUpper(name), poAPIKey)
	logger.InfoLogger.Printf("[%v] User Key=[%v]", strings.ToUpper(name), poUserKey)

	if poUserKey == "" || poAPIKey == "" {
		logger.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), "Pushover User Key or API Token not set, message not sent")
		return nil
	}

	app := pushover.New(poAPIKey)
	recipient := pushover.NewRecipient(poUserKey)
	var inCallbackUrl string

	if key != 0 {
		//inCallbackUrl = support.Application.BaseURL() + "view/" + key
		inCallbackUrl = fmt.Sprintf("http://%v:%v/view/%v", set.ApplicationHost(), set.ApplicationPort(), key)
	} else {
		inCallbackUrl = fmt.Sprintf("http://%v:%v/dashboard/", set.ApplicationHost(), set.ApplicationPort())
	}

	message := &pushover.Message{
		Message:    inMessage,
		Title:      inTitle,
		Priority:   pushover.PriorityNormal,
		URL:        inCallbackUrl,
		URLTitle:   set.ApplicationName(),
		Timestamp:  time.Now().Unix(),
		Retry:      60 * time.Second,
		Expire:     time.Hour,
		DeviceName: "",
		//CallbackURL: "http://yourapp.com/callback",
		CallbackURL: inCallbackUrl,
		Sound:       pushover.SoundCosmic,
	}
	//Spew(*message)

	logger.InfoLogger.Printf("[%v] Message Title=[%v] Message=[%v]", strings.ToUpper(name), message.Title, message.Message)

	_, err := app.SendMessage(message, recipient)
	if err != nil {
		logger.WarningLogger.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		return err
	}

	return nil
}

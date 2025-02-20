package actions

import "strings"

type Action struct {
	Name        string `storm:"Name" csv:"Name"`
	userDefined bool   `storm:"userDefined" csv:"userDefined"`
	Code        string `storm:"Code" csv:"Code"`
}

var LIST Action = Action{Name: "List", userDefined: false, Code: "LIST"}
var VIEW Action = Action{Name: "View", userDefined: false, Code: "VIEW"}
var EDIT Action = Action{Name: "Edit", userDefined: false, Code: "EDIT"}
var UPDATE Action = Action{Name: "Update", userDefined: false, Code: "UPDATE"}
var DELETE Action = Action{Name: "Delete", userDefined: false, Code: "DELETE"}
var CREATE Action = Action{Name: "Create", userDefined: false, Code: "CREATE"}
var ENABLE Action = Action{Name: "Enable", userDefined: false, Code: "ENABLE"}
var DISABLE Action = Action{Name: "Disable", userDefined: false, Code: "DISABLE"}
var RESET Action = Action{Name: "Reset", userDefined: false, Code: "RESET"}
var ROUTE Action = Action{Name: "Route", userDefined: false, Code: "ROUTE"}
var MESSAGE Action = Action{Name: "Message", userDefined: false, Code: "MESSAGE"}
var LOGIN Action = Action{Name: "Login", userDefined: false, Code: "LOGIN"}
var LOGOUT Action = Action{Name: "Logout", userDefined: false, Code: "LOGOUT"}
var TIMEOUT Action = Action{Name: "Timeout", userDefined: false, Code: "TIMEOUT"}
var API Action = Action{Name: "API", userDefined: false, Code: "API"}
var GET Action = Action{Name: "GET", userDefined: false, Code: "GET"}
var GETALL Action = Action{Name: "GETALL", userDefined: false, Code: "GETALL"}
var EXPORT Action = Action{Name: "EXPORT", userDefined: false, Code: "EXPORT"}
var IMPORT Action = Action{Name: "Import", userDefined: false, Code: "IMPORT"}
var PROCESS Action = Action{Name: "Process", userDefined: false, Code: "PROCESS"}
var REPAIR Action = Action{Name: "Repair", userDefined: false, Code: "REPAIR"}
var AUDIT Action = Action{Name: "Audit", userDefined: false, Code: "AUDIT"}
var CONNECT Action = Action{Name: "Connect", userDefined: false, Code: "CONNECT"}
var DISCONNECT Action = Action{Name: "Disconnect", userDefined: false, Code: "DISCONNECT"}
var BACKUP Action = Action{Name: "Backup", userDefined: false, Code: "BACKUP"}
var VALIDATE Action = Action{Name: "Validate", userDefined: false, Code: "VALIDATE"}
var INITIALISE Action = Action{Name: "Initialise", userDefined: false, Code: "INITIALISE"}
var SHUTDOWN Action = Action{Name: "Shutdown", userDefined: false, Code: "SHUTDOWN"}
var RESTART Action = Action{Name: "Restart", userDefined: false, Code: "RESTART"}
var RELOAD Action = Action{Name: "Reload", userDefined: false, Code: "RELOAD"}

func New(name string) Action {
	return Action{Name: name, userDefined: true}
}

func (bt Action) GetName() string {
	return strings.ToUpper(bt.Name)
}
func (bt Action) GetCode() string {
	return strings.ToUpper(bt.Code[0:4])
}

func (bt Action) IsUserDefined() bool {
	return bt.userDefined
}

func (bt Action) Is(in Action) bool {
	return bt.Name == in.Name && bt.userDefined == in.userDefined
}

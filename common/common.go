package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/mt1976/frantic-plum/paths"
)

// var Data ConfigurationModel
var name = "common"
var filename = ""
var commonSettingsFile = "common"

var TRUE = "true"
var FALSE = "false"

func Get() *Settings {

	var thisConfig Settings
	filename = paths.Application().String() + paths.Config().String() + string(os.PathSeparator) + commonSettingsFile + ".toml"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("[%v] Error=[%v]", strings.ToUpper(name), err.Error())
		panic(err)
	}

	toml.Unmarshal(content, &thisConfig)

	return &thisConfig
}

func (s *Settings) Spew() {
	nm := strings.ToUpper(name)
	f2 := "[%v] %-15v : %v\n"
	fmt.Printf("[%v] Config loaded from file: %v\n", nm, filename)
	fmt.Printf("[%v][APPLICATION]\n", nm)
	fmt.Printf(f2, nm, "Name", s.Application.Name)
	fmt.Printf(f2, nm, "Version", s.Application.Version)
	fmt.Printf(f2, nm, "Description", s.Application.Description)
	fmt.Printf(f2, nm, "Prefix", s.Application.Prefix)
	fmt.Printf(f2, nm, "Environment", s.Application.Environment)
	fmt.Printf(f2, nm, "ReleaseDate", s.Application.ReleaseDate)
	fmt.Printf(f2, nm, "Copyright", s.Application.Copyright)
	fmt.Printf(f2, nm, "License", s.Application.License)
	fmt.Printf(f2, nm, "Author", s.Application.Author)

	fmt.Printf("[%v][SERVER]\n", nm)

	fmt.Printf(f2, nm, "Port", s.Server.Port)
	fmt.Printf(f2, nm, "Protocol", s.Server.Protocol)
	fmt.Printf(f2, nm, "Environment", s.Server.Environment)

	fmt.Printf("[%v][ASSETS]\n", nm)

	fmt.Printf(f2, nm, "Logo", s.Assets.Logo)
	fmt.Printf(f2, nm, "Favicon", s.Assets.Favicon)

	fmt.Printf("[%v][DATES]\n", nm)

	fmt.Printf(f2, nm, "DateTimeFormat", s.Dates.DateTimeFormat)
	fmt.Printf(f2, nm, "DateFormat", s.Dates.DateFormat)
	fmt.Printf(f2, nm, "TimeFormat", s.Dates.TimeFormat)
	fmt.Printf(f2, nm, "Backup", s.Dates.Backup)
	fmt.Printf(f2, nm, "BackupFolder", s.Dates.BackupFolder)
	fmt.Printf(f2, nm, "Human", s.Dates.Human)
	fmt.Printf(f2, nm, "DMY2", s.Dates.DMY2)
	fmt.Printf(f2, nm, "YMD", s.Dates.YMD)
	fmt.Printf(f2, nm, "Internal", s.Dates.Internal)

	fmt.Printf("[%v][HISTORY]\n", nm)
	fmt.Printf(f2, nm, "MaxEntries", s.History.MaxEntries)

	fmt.Printf("[%v][MESSAGE]\n", nm)
	fmt.Printf(f2, nm, "TypeKey", s.Message.TypeKey)
	fmt.Printf(f2, nm, "TitleKey", s.Message.TitleKey)
	fmt.Printf(f2, nm, "ContentKey", s.Message.ContentKey)
	fmt.Printf(f2, nm, "ActionKey", s.Message.ActionKey)

	fmt.Printf("[%v][DISPLAY]\n", nm)
	fmt.Printf(f2, nm, "Delimiter", s.Display.Delimiter)

}

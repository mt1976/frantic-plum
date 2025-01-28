package logger

import (
	"strings"
)

func Banner(class, name, action string) {
	hdr := "------------------------------------------------------------------------"
	InfoLogger.Println(hdr)
	InfoLogger.Printf("[%v] Activity=[%v] - %v", strings.ToUpper(class), name, action)
	InfoLogger.Println(hdr)
}

func ServiceBanner(class, name, action string) {
	hdr := "------------------------------------------------------------------------"
	ServiceLogger.Println(hdr)
	ServiceLogger.Printf("[%v] Activity=[%v] - %v", strings.ToUpper(class), name, action)
	ServiceLogger.Println(hdr)
}

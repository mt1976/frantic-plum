package mock

import (
	"github.com/mt1976/frantic-core/logHandler"
)

func report(in string) {
	logHandler.InfoLogger.Printf("Mocking - %s\n", in)
}

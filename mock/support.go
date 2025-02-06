package mock

import (
	"github.com/mt1976/frantic-plum/logger"
)

func report(in string) {
	logger.InfoLogger.Printf("Mocking - %s\n", in)
}

package jobs

import (
	"fmt"
	"time"

	"github.com/jsuar/go-cron-descriptor/pkg/crondescriptor"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/logHandler"
)

func StartOfDay(t time.Time) time.Time {
	// Purpose: To remove the time from a date
	return dateHelpers.StartOfDay(t)
}

func BeforeOrEqualTo(t1, t2 time.Time) bool {
	return dateHelpers.IsBeforeOrEqualTo(t1, t2)
}

func AfterOrEqualTo(t1, t2 time.Time) bool {
	return dateHelpers.IsAfterOrEqualTo(t1, t2)
}

func NextRun(name, schedule string) string {
	// Purpose: To determine the next run time of a job
	rtn := fmt.Sprintf("[%v] [%v] NextRun=[%v]", domain, name, GetHumanReadableCronFreq(schedule))
	logHandler.ServiceLogger.Println(rtn)
	return rtn
}

func Announce(name, inAction string) {
	logHandler.ServiceBanner(domain, name, inAction)
}

func GetHumanReadableCronFreq(freq string) string {
	bkHuman1, _ := crondescriptor.NewCronDescriptor(freq)
	bkHuman, _ := bkHuman1.GetDescription(crondescriptor.Full)
	return *bkHuman
}

package timing

import (
	"strconv"
	"strings"
	"time"

	"github.com/mt1976/frantic-core/logHandler"
	"golang.org/x/exp/rand"
)

var name = "Timing"

type Stopwatch struct {
	table   string
	action  string
	start   time.Time
	msg     string
	end     time.Time
	duraton time.Duration
}

func Start(table, action, msg string) Stopwatch {
	start := time.Now()

	//logger.InfoLogger.Printf("TIM: %v %v: [%v]", name, msg, start)
	return Stopwatch{table, action, start, msg, time.Now(), time.Duration(0)}
}

func (w *Stopwatch) Stop(count int) {
	w.end = time.Now()
	w.duraton = w.end.Sub(w.start)
	logHandler.TimingLogger.Printf("Object=[%v] Action=[%v] Msg=[%v] Count=[%v] Duration=[%v]", w.table, strings.ToUpper(w.action), w.msg, count, w.duraton)
}

// SnoozeFor snoozes the application for a given amount of time
// The function SnoozeFor takes in a polling interval and calls the snooze function with that interval.
func SnoozeFor(inPollingInterval string) {
	snooze(inPollingInterval)
}

// Snooze snoozes for a random period
// The Snooze function generates a random number between 0 and 10 and then calls the snooze function
// with that number as a string argument.
func Snooze() {
	rand.Seed(uint64(time.Now().UnixNano()))
	n := rand.Intn(10) // n will be between 0 and 10
	snooze(strconv.Itoa(n))
}

func snooze(inPollingInterval string) {
	pollingInterval, _ := strconv.Atoi(inPollingInterval)
	logHandler.InfoLogger.Printf("Snooze... Zzzzzz.... %d seconds...", pollingInterval)
	time.Sleep(time.Duration(pollingInterval) * time.Second)
}

package support

import "github.com/mt1976/frantic-plum/config"

var TRUE = "true"
var FALSE = "false"

var d *config.Configuration

func init() {
	d = config.Get()
}

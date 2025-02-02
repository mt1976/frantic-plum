package main

import "github.com/mt1976/frantic-plum/common"

var TRUE = "true"
var FALSE = "false"

var d *common.Settings

func init() {
	d = common.Get()
}

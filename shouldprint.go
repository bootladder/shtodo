package main

import (
	"time"
)

func shouldPrint(tnow, tbefore time.Time, duration int) bool {

	var delta = tnow.Sub(tbefore).Seconds()
	if int(delta) >= duration {
		return true
	}
	return false
}

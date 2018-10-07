package main

import "time"

func shouldPull(interval int) bool {

	return isTimeIntervalPassed(pathToLastPullTime, interval)
}

func shouldPrint(interval int) bool {

	return isTimeIntervalPassed(pathToLastPrintTime, interval)
}

func isTimeIntervalPassed(filename string, interval int) bool {

	touch(filename)
	var tbefore = readTimeFromFile(filename)
	var tnow = time.Now().UTC()

	return checkInterval(tnow, tbefore, interval)
}

func checkInterval(tnow, tbefore time.Time, interval int) bool {

	var delta = tnow.Sub(tbefore).Seconds()
	if int(delta) >= interval {
		return true
	}
	return false
}

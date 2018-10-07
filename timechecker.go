package main

import "time"

func shouldPull(interval int) bool {

	touch(pathToLastPullTime)
	var tbefore = readTimeFromFile(pathToLastPullTime)
	var tnow = time.Now().UTC()

	var delta = tnow.Sub(tbefore).Seconds()
	if int(delta) >= interval {
		return true
	}
	return false
}

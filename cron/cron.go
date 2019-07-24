/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the Main.
 */
package cron

import (
	"log"
	"github.com/robfig/cron"
)

// Created Global CRON
var cronJob = cron.New()
var cronStatus = false
var triggerCount = 0

func startCronJob() {
	registerCronJob(trigger)
}

func registerCronJob(fn func()) {
	if fn != nil {
		cronJob.AddFunc("5 * * * * *", fn)
		cronJob.Start()
		cronStatus = true
	}
}

func stopCronJob() {
	if cronStatus == true  {
		cronJob.Stop()
		cronStatus = false
	}
}

func resetCronJob() {
	cronJob = cron.New()
	cronStatus = false
	triggerCount = 0
}

func trigger() {
	triggerCount = triggerCount + 1
	log.Println("Trigger on every the 5 secs")
}
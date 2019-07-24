/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the cron.
 */
 package cron

 import "testing"
 
 func TestCronJobInitialStatus(t *testing.T) {
	t.Log("CRON: Test Cron Initial Stauts")
	if cronStatus != false {
		t.Errorf("CRON ERROR: Intial status should be %s", "false")
	}

	t.Log("CRON: Test Cron Initial Job trigger count")

	if triggerCount != 0 {
	   t.Errorf("CRON ERROR: Initial trigger should be %d", triggerCount)
	}
}

func TestRegisterCronJob(t *testing.T) {
	t.Log("CRON: Test Cron Initial Job start")
	registerCronJob(nil)

	if cronStatus == true {
		t.Errorf("CRON ERROR: After Job started. The status should be %s", "true")
	}
 }

 func TestCronJobStart(t *testing.T) {
	t.Log("CRON: Test Cron Initial Job start")
	startCronJob()

	if cronStatus != true {
		t.Errorf("CRON ERROR: After Job started. The status should be %s", "true")
	}
 }
 

 func TestCronJobTriggerCount(t *testing.T) {
	t.Log("CRON: Test Cron Job trigger count")

	trigger()
	if triggerCount != 0 {
		t.Logf("CRON ERROR: Number of times Job trigger called - %d", triggerCount)
	 }
}

func TestCronJobStop(t *testing.T) {
	t.Log("CRON: Test Cron Job Stop")

	stopCronJob()

	if cronStatus == true {
		t.Errorf("CRON ERROR: After Job stoped. The status should be %s", "false")
	}
 }

 func TestCronJobReset(t *testing.T) {
	t.Log("CRON: Test Cron Job Reset")

	resetCronJob()

	if cronStatus == true {
		t.Errorf("CRON ERROR: After Job Resetted. The status should be %s", "false")
	}

	if triggerCount != 0 {
		t.Errorf("CRON ERROR: Initial trigger should be %d", triggerCount)
	 }
 }
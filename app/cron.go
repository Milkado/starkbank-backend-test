package app

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/Milkado/stark-backend-test/helpers"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/robfig/cron/v3"
)

func StartCron(c *echo.Context) error {
	newCron := cron.New()
	cronId := uuid.New().String() //Id for logging
	var wg sync.WaitGroup
	var i = 1
	newCron.AddFunc("@every 3m", func() {
		wg.Add(1)
		defer wg.Done()
		//CreateInvoice()
		helpers.Log("Cont at: "+strconv.Itoa(i), "./logs/count_times")
		i++
		fmt.Println("Task ran at: ", time.Now().Format("2006-01-02 15:04:05"))
	})
	newCron.Start()

	message := "Cron: " + cronId + " started"
	helpers.Log(message, "./logs/cron_times.txt")

	stopAfter := 24 * time.Minute

	//Starts a Goroutine
	go func() {
		//Blocking call to sleep Goroutine for exactly 24h
		<-time.After(stopAfter)
		fmt.Println("Stopping cron at: ", time.Now().String())

		stopCtx := newCron.Stop()
		<-stopCtx.Done()

		// Ensure any running task finishes
		wg.Wait()
		fmt.Println("Cron finished gracefully")

		message = "Cron: " + cronId + " stopped"
		helpers.Log(message, "./logs/cron_times.txt")
	}()

	return c.String(http.StatusOK, "Cron started, running for the next 24h every 3h")
}

package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Milkado/stark-backend-test/helpers"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/robfig/cron/v3"
)

func StartCron(c *echo.Context) error {
	newCron := cron.New()
	cronId := uuid.New().String() //Id for logging

	newCron.AddFunc("@every 1h", func() {
		CreateInvoice()
		fmt.Println("Task ran at: ", time.Now().Format("2006-01-02 15:04:05"))
	})
	newCron.Start()

	message := "Cron: " + cronId + " started"
	helpers.Log(message, "./logs/cron_times.txt")

	stopAfter := 2 * time.Hour

	//Starts a Goroutine
	go func() {
		//Blocking call to sleep Goroutine for X time
		<-time.After(stopAfter)
		fmt.Println("Stopping cron at: ", time.Now().String())
		stopCtx := newCron.Stop()

		//Executes at the end of Goroutine (don't interrupt currently running cron tasks)
		defer func() {
			<-stopCtx.Done()
			fmt.Println("Cron finished gracefully")
		}()

		message = "Cron: " + cronId + " stopped"

		helpers.Log(message, "./logs/cron_times.txt")
	}()

	return c.String(http.StatusOK, "Cron started, running for the next 24h every 3h")
}

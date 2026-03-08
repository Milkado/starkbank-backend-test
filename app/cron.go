package app

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Milkado/stark-backend-test/helpers"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	cron "github.com/netresearch/go-cron"
)

var (
	cronMutex     sync.Mutex
	isCronRunning bool
)

func StartCron(c *echo.Context) error {
	cronMutex.Lock()
	if isCronRunning {
		cronMutex.Unlock()
		return c.String(http.StatusConflict, "A cron job is already running.")
	}
	isCronRunning = true
	cronMutex.Unlock()

	newCron := cron.New()
	cronId := uuid.New().String() //Id for logging
	var wg sync.WaitGroup

	newCron.AddFunc("@every 3h", func() {
		wg.Add(1)
		defer wg.Done() //Ensures that last task is completed before shutting down (some times last task is skipped on AWS)
		CreateInvoice()
		fmt.Println("Task ran at: ", time.Now().Format("2006-01-02 15:04:05"))
	})
	newCron.Start()

	message := "Cron: " + cronId + " started"
	helpers.Log(message, "./logs/cron_times.txt")

	stopAfter := 24 * time.Hour

	//Starts a Goroutine
	go func() {
		defer func() {
			cronMutex.Lock()
			isCronRunning = false
			cronMutex.Unlock()
		}()

		//Blocking call to sleep Goroutine for exactly 24h
		<-time.After(stopAfter)
		fmt.Println("Stopping cron at: ", time.Now().String())

		newCron.StopAndWait()

		// Ensure any running task finishes
		wg.Wait()
		fmt.Println("Cron finished gracefully")

		message = "Cron: " + cronId + " stopped"
		helpers.Log(message, "./logs/cron_times.txt")
	}()

	return c.String(http.StatusOK, "Cron started, running for the next 24h every 3h")
}

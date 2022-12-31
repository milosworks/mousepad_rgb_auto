package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/vyrekxd/exec_cmd_cronjob/pkg/config"
)

func cmdExec(action string) {
	Cmd := exec.Command("/bin/sh", "-c", "sudo uhubctl -a "+action+" -p "+config.Port+" -l "+config.Location)

	if err := Cmd.Run(); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Program started\nStarting test...")

	cmdExec(config.Actions[0])

	time.Sleep(5 * time.Second)

	cmdExec(config.Actions[1])

	fmt.Println("Test finished with no errors")

	scheduler := gocron.NewScheduler(time.Now().Location())

	scheduler.Day().At("8:00").Do(func() {
		t := time.Now()

		cmdExec(config.Actions[1])

		fmt.Println("Mousepad is now on: " + t.Format(time.Kitchen))
	})

	scheduler.Day().At("22:00").Do(func() {
		t := time.Now()

		cmdExec(config.Actions[0])

		fmt.Println("Mousepad is now off: " + t.Format(time.Kitchen))
	})

	scheduler.Every(1).Day().At("13:56").Do(func() {
		t := time.Now()

		cmdExec("toggle")

		fmt.Println("Mousepad is now toggled: " + t.Format(time.Kitchen))
	})

	scheduler.StartAsync()

	fmt.Println("Scheduler loaded with jobs and started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	scheduler.Stop()

	fmt.Println("\nJobs stopped and program stopped")
}

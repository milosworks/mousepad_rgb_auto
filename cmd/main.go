package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
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

	jobs := cron.New()

	jobs.AddFunc("0 0 8 * * *", func() {
		t := time.Now()

		cmdExec(config.Actions[1])

		fmt.Println("Mousepad is now on: " + t.Format(time.Kitchen))
	})

	jobs.AddFunc("0 0 22 * * *", func() {
		t := time.Now()

		cmdExec(config.Actions[0])

		fmt.Println("Mousepad is now off: " + t.Format(time.Kitchen))

	})

	fmt.Println("Jobs loaded")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	jobs.Stop()

	fmt.Println("Jobs stopped and program stopped")
}

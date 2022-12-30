package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/vyrekxd/exec_cmd_cronjob/pkg/config"
)

func cmdExec(action string) {
	CmdString := "sudo uhubctl -a " + action + " -p " + config.Port + " -l " + config.Location
	Cmd := exec.Command(CmdString)

	if err := Cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
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

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}

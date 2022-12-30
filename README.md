# Mousepad RGB Auto

A simple go program to automatically turn off and on a mousepad with a toggle switch for usb power in a raspberry pi 3B.

## Requirements

-   Raspberry Pi 3B
-   A [Node.JS](https://nodejs.org/en/) and [PM2](https://pm2.keymetrics.io/) installation in the raspberry pi.
-   A [uhubctl](https://github.com/mvp/uhubctl) installation also in the raspberry pi (Its optional to give it permissions as we will use sudo).
-   A [Golang](https://golang.org/doc/install) installation in the raspberry pi.

## Installation

-   Clone this repo and change variable sin `pkg/config/app.go` to your needs, being `port` and `location`, you can also change the **cron** format in `cmd/main.go` to your needs.
-   Use this command to build the binary: `go build -o mousepad_rbg_auto cmd/main.go`
-   Then use pm2 to run it in the background: `pm2 start mousepad_rbg_auto --name Mousepad RGB Auto`

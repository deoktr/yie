package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"
)

var (
	intervalFlag = flag.String("i", "", "interval between each notifications")
	summaryFlag  = flag.String("s", "", "notification message summary")
	bodyFlag     = flag.String("m", "", "notification message body")

	interval time.Duration
	summary  string
	body     string
)

func notify(summary string, body string) {
	cmd := exec.Command("notify-send", summary, body)
	err := cmd.Run()
	if err != nil {
		fmt.Println("error creating notification:", err)
		return
	}
}

func main() {
	flag.Parse()

	fmt.Println("starting yie")

	interval, err := time.ParseDuration(*intervalFlag)
	if err != nil {
		fmt.Println("failed to parse interval:", err)
		return
	}

	summary = *summaryFlag
	body = *bodyFlag

	for {
		notify(summary, body)
		time.Sleep(interval)
	}
}

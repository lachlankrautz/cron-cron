package main

import (
	"cron-cron/icon"
	"fmt"
	"github.com/getlantern/systray"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var running = true
var mStatus *systray.MenuItem
var monitorTick = 2 * time.Second

func main() {
	go systray.Run(onReady, onExit)

	log.Println("Start monitoring cron")

	for running {
		ensureCron()
		time.Sleep(monitorTick)
	}
}

func SetMonitorTick(duration time.Duration) {
	monitorTick = duration
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("WSL2 Cron Monitor")
	systray.SetTooltip("WSL2 Cron Monitor")
	mStatus = systray.AddMenuItem("-", "-")
	mQuit := systray.AddMenuItem("Exit", "Stop cron and exit")

	mQuit.SetIcon(icon.Data)

	go func() {
		for running {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {
	running = false

	log.Println("Stop monitoring cron")
	os.Exit(0)
}

func ensureCron() {
	if isCronRunning() {
		return
	}

	if err := startCron(); err != nil {
		log.Fatal(err)
	}
}

func isCronRunning() bool {
	cmd := exec.Command("wsl", "--", "pgrep", "-o", "cron")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()

	if err != nil {
		return false
	}

	message := fmt.Sprintf("Cron pid: %s", out)
	mStatus.SetTitle(message)
	mStatus.SetTooltip(message)
	log.Printf(message)

	return true
}

func startCron() error {
	log.Println("Starting cron")

	cmd := exec.Command("wsl", "--" , "sudo", "cron")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd.Start()
}

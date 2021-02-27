# cron-cron

cron-cron lives in the system tray and monitors and restarts a wsl cron. If the cron isn't running it will restart.

## Build

go build -ldflags -H=windowsgui -tags release

## Run

Run the executable and it will display in the system tray and show the pid of the wsl cron process.


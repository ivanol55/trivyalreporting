package manageWebserver

import (
	"fmt";
	"os/exec";
)

func StartStopReportsServer(option string) {
	switch option {
	case "start":
		fmt.Println("Starting up the local reports webserver...")
		var command string = "docker compose up -d"
		var commandObject = exec.Command("sh", "-c", command)
		fmt.Println("")
		commandObject.Run()
	case "stop":
		fmt.Println("Stopping the local reports webserver...")
		var command string = "docker compose down"
		var commandObject = exec.Command("sh", "-c", command)
		fmt.Println("")
		commandObject.Run()
	}
}